package mediamtx


import (
	"github.com/go-resty/resty/v2"
)


type mediamtxSdk struct {
	Url        string
	restClient *resty.Client
}