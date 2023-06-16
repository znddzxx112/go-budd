package server

import (
	"github.com/dghubble/sling"
	"net/http"
)

var localBase *sling.Sling

func init() {
	localBase = sling.New().Base("http://127.0.0.1:7890").Client(http.DefaultClient)
}
