package mediamtx

import (
	"log"
	"github.com/go-resty/resty/v2"
)

// Builder is used to build a new mediamtx client
func BuildMediaMtx(url string, debug bool) (IMediamtxClient, error){
	// init mediamtxClient
	mediamtxClient := &mediamtxSdk{
		Url:        url,
		restClient: resty.New(),
	}
	//
	if debug {
		mediamtxClient.restClient.SetDebug(true)
		mediamtxClient.debug = true
		log.Println("Debug mode is enabled for the haivision client ")
	}
	// You can override all below settings and options at request level if you want to
	//--------------------------------------------------------------------------------
	// Host URL for all request. So you can use relative URL in the request
	mediamtxClient.restClient.SetBaseURL(url)
	//
	return mediamtxClient, nil
}


func (o *mediamtxSdk) debugPrint(data interface{}) {
	if o.debug {
		log.Println(data)
	}
}