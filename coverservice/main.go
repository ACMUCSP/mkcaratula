package main

import (
    "math/rand"
    "time"
    "github.com/sirupsen/logrus"
    "github.com/ACMUCSP/mkcaratula/common"
    "github.com/ACMUCSP/mkcaratula/coverservice/service"
)

func main() {
    logrus.Info("Starting %s", service.AppName)
    rand.Seed(time.Now().UnixNano())
    common.StartWebServer("8000", service.Routes)
}
