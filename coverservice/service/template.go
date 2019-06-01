package service

import (
    "text/template"
    "github.com/sirupsen/logrus"
)

var coverTemplate *template.Template


func GetCoverTemplate() (tmpl *template.Template, err error) {
    if coverTemplate == nil {
        coverTemplate, err = template.ParseFiles("templates/cover.tex")
        if err != nil {
            logrus.Errorf("Failed to process template: %s", err)
            return
        }
    }

    tmpl = coverTemplate
    return
}
