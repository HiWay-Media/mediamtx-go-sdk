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
	//
	return mediamtxClient, nil
}


func (o *haivisionSdk) debugPrint(data interface{}) {
	if o.debug {
		log.Println(data)
	}
}