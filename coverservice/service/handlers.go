package service

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/ACMUCSP/mkcaratula/common"
    "path"
    "os"
)

type CoverResponse struct {
    Url string `json:"url"`
}

type TexResponse struct {
    Tex string `json:"tex"`
}

const kFileNotFound = "File not found"
const kFailedToProcessIncomingData = "Failed to process incoming data"
const kTmpDir = "tmp"
const kGenDir = "gen"


func processContext(ctx *CoverContext) {
    // TODO: filter quotes
}


func GenerateCoverHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var ctx CoverContext
    err := decoder.Decode(&ctx)
    if err != nil {
        data, _ := json.Marshal(common.ErrorResponse{kFailedToProcessIncomingData})
        common.WriteJsonResponse(w, http.StatusBadRequest, data)
        return
    }
    processContext(&ctx)

    if _, ok := r.URL.Query()["tex"]; ok {

        tex, err := GetTexOutput(ctx)

        if err != nil {
            common.WriteInternalErrorResponse(w)
        }

        data, _ := json.Marshal(TexResponse{tex})
        common.WriteJsonResponse(w, http.StatusOK, data)
    } else {

        filename, err := GenerateCoverPage(kTmpDir, kGenDir, ctx)

        if err != nil {
            common.WriteInternalErrorResponse(w)
            return
        }
        pdfUrl := *r.URL
        pdfUrl.Path = path.Join(path.Dir(r.URL.Path), filename)

        data, _ := json.Marshal(CoverResponse{pdfUrl.String()})
        common.WriteJsonResponse(w, http.StatusOK, data)
    }
}


func GetCoverHandler(w http.ResponseWriter, r *http.Request) {
    key := mux.Vars(r)["key"]
    filename := path.Join(kGenDir, key + ".pdf")
    if _, err := os.Stat(filename); err == nil {
        w.Header().Set("Content-Type", "application/pdf")
        w.Header().Set("X-Accel-Redirect", filename)
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(kFileNotFound))
    }
}
