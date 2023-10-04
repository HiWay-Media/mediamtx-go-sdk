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
	RemovePathConfig(name string) error
	GetHlsMuxers( requestQuery ListRequest ) error
	GetHlsMuxer( name string ) error
	GetPathList( requestQuery ListRequest ) error
	GetPath( name string ) error
	GetRtspConnections( requestQuery ListRequest ) error
	GetRtspConnection( name string ) error
	GetRtspSessions( requestQuery ListRequest ) error
	GetRtspSession( name string ) error
	KickRtspSession( name string ) error
	KickRtspConnection( name string ) error
	GetRtmpConnections( requestQuery ListRequest ) error
	GetRtmpConnection( name string ) error
	KickRtmpConnection( name string ) error
	GetSrtConnections( requestQuery ListRequest ) error
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