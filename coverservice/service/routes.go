package service

import (
    "github.com/ACMUCSP/mkcaratula/common"
)


var Routes = common.Routes{
    common.Route{
        "GenerateCover",
        "POST",
        "/caratula/generate",
        GenerateCoverHandler,
    },
    common.Route{
        "GetCover",
        "GET",
        "/caratula/{key:[a-z]+}.pdf",
        GetCoverHandler,
    },
}
