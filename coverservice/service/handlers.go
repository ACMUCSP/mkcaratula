package service

import (
    "path"
    "os"
    "bytes"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/ACMUCSP/mkcaratula/common"
)

const AppName = "coverservice"

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


func filterSpecialCharacters(str string) string {
    var buffer bytes.Buffer
    openQuotes := false
    for _, c := range str {
        switch c {
        case '"':
            if openQuotes {
                buffer.WriteString("''")
            } else {
                buffer.WriteString("``")
            }
            openQuotes = !openQuotes
        case '&', '%', '$', '#', '_', '{', '}':
            buffer.WriteRune('\\')
            buffer.WriteRune(c)
        case '~':
            buffer.WriteString("\textasciitilde")
        case '^':
            buffer.WriteString("\textasciicircum")
        case '\\':
            buffer.WriteString("\textbackslash")
        default:
            buffer.WriteRune(c)
        }
    }
    return buffer.String()
}


func processContext(ctx *CoverContext) {
    ctx.Career = filterSpecialCharacters(ctx.Career)
    ctx.Title = filterSpecialCharacters(ctx.Title)
    ctx.Course = filterSpecialCharacters(ctx.Course)
    ctx.Semester = filterSpecialCharacters(ctx.Semester)
    for i, name := range ctx.Names {
        ctx.Names[i] = filterSpecialCharacters(name)
    }
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
        w.Header().Set("X-Accel-Redirect", "/" + path.Join(AppName, filename))
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(kFileNotFound))
    }
}
