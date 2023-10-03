package mediamtx

/*
GET http://localhost:9997/v2/config/get
returns the configuration.
*/

func(o *mediamtxSdk) GetConfig() (*Config, error)  {


	return nil, nil
}

/*
POST http://localhost:9997/v2/config/set
changes the configuration.
all fields are optional. paths can't be edited with this request, use /v2/config/paths/{operation}/{name} to edit them.
*/