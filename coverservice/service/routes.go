package service

import (
    "github.com/ACMUCSP/mkcaratula/common"
)


var Routes = common.Routes{
    common.Route{
        "GenerateCover",
        "POST",
        "/generate",
        GenerateCoverHandler,
    },
    common.Route{
        "GetCover",
        "GET",
        "/{key:[a-z]+}.pdf",
        GetCoverHandler,
    },
}
