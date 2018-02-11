package service

import (
    "io/ioutil"
    "os/exec"
    "io"
    "text/template"
    "math/rand"
    "os"
    "path"
    "bytes"
    "github.com/sirupsen/logrus"
)

type CoverContext struct {
    Career string `json:"career"`
    Title string `json:"title"`
    Course string `json:"course"`
    Gender string `json:"gender"`
    Names []string `json:"names"`
    Semester string `json:"semester"`
}

func (ctx CoverContext) StudentsDeclaration() string {
    var students string
    if len(ctx.Names) > 1 {
        if ctx.Gender == "f" {
            students = "Las alumnas"
        } else {
            students = "Los alumnos"
        }
        return students + " declaran"
    } else {
        if ctx.Gender == "f" {
            students = "La alumna"
        } else {
            students = "El alumno"
        }
        return students + " declara"
    }
}

const kKeyLength = 4

var kAlphabet = [][]byte{
    []byte("bcdfghjklmnpqrstvwxyz"),
    []byte("aeiou"),
}

func easyRandom() string {
    buffer := make([]byte, kKeyLength)
    for i := 0; i < kKeyLength; i++ {
        buffer[i] = kAlphabet[i & 1][rand.Intn(len(kAlphabet[i & 1]))]
    }
    return string(buffer)
}


func storeCoverPage(tmpDir, genDir string) (filename string, err error) {
    filename = easyRandom() + ".pdf"
    err = os.Rename(path.Join(tmpDir, "texput.pdf"), path.Join(genDir, filename))
    if err != nil {
        filename = ""
    }
    return
}


func GenerateCoverPage(tmpDir, genDir string, ctx CoverContext) (filename string, err error) {
    var tmpl *template.Template
    tmpl, err = GetCoverTemplate()
    if err != nil {
        return
    }

    cmd := exec.Command("pdflatex", "-halt-on-error")
    cmd.Dir, err = ioutil.TempDir(tmpDir, "")
    if err != nil {
        logrus.Errorf("Failed to create temporary directory: %s", err)
        return
    }
    var texStdin io.WriteCloser
    texStdin, err = cmd.StdinPipe()
    if err != nil {
        logrus.Errorf("Failed to open pipe for pdflatex: %s", err)
        return
    }
    go func() {
        defer texStdin.Close()
        tmpl.Execute(texStdin, ctx)
    }()

    var out []byte
    out, err = cmd.CombinedOutput()
    if err != nil {
        logrus.Errorf("Failed to run pdflatex: %s", err)
        logrus.Debug(string(out))
        return
    }

    filename, err = storeCoverPage(cmd.Dir, genDir)

    if err != nil {
        logrus.Errorf("Failed to store file: %s", err)
        return
    }

    err = os.RemoveAll(cmd.Dir)

    if err != nil {
        filename = ""
        logrus.Errorf("Failed to remove temporary directory: %s", err)
    }

    return
}


func GetTexOutput(ctx CoverContext) (tex string, err error) {
    var tmpl *template.Template
    tmpl, err = GetCoverTemplate()
    if err != nil {
        return
    }
    var buffer bytes.Buffer
    tmpl.Execute(&buffer, ctx)
    tex = buffer.String()
    return
}
