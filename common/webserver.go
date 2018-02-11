package common


import (
    "net/http"
    "github.com/sirupsen/logrus"
)


func StartWebServer(port string, routes Routes) {
    r := NewRouter(routes)
    http.Handle("/", r)
    logrus.Infof("Starting HTTP service at port %s", port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        logrus.Errorf("An error ocurred starting HTTP listener at port %s", port)
        logrus.Error("Error: " + err.Error())
    }
}
