package mediamtx


import (
	"github.com/go-resty/resty/v2"
)


type mediamtxSdk struct {
	Url        string
	restClient *resty.Client
	debug 		bool
}


type IMediamtxClient interface {
	//
	HealthCheck() error
	GetConfig() (*Config, error)
	SetConfig(name string) error
	AddPathConfig(name string) error
	EditPathConfig(name string) error
	GetHlsMuxers( requestQuery ListRequest ) error
	GetPathList( requestQuery ListRequest ) error
	//
}


func (o *mediamtxSdk) HealthCheck() error {
	_, err := o.restyGet(o.Url, nil)
	if err != nil {
		return nil
	}
	return nil
}


// Resty Methods

func (o *mediamtxSdk) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// get request
func (o *mediamtxSdk) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}