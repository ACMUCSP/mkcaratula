package main

import (
    "math/rand"
    "time"
    "github.com/sirupsen/logrus"
    "github.com/ACMUCSP/mkcaratula/common"
    "github.com/ACMUCSP/mkcaratula/coverservice/service"
)

var appName = "coverservice"

func main() {
    logrus.Info("Starting %s", appName)
    rand.Seed(time.Now().UnixNano())
    common.StartWebServer("8000", service.Routes)
}
