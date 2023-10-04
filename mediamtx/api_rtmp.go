package mediamtx

/*
GET http://localhost:9997/v2/rtmpconns/list
returns all RTMP connections.
QUERY PARAMETERS
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/

func (o *mediamtxSdk) GetRtmpConnections( requestQuery ListRequest ) error{
	o.debugPrint("GetRtmpConnections ")
	// need to transform the struct
	resp, err := o.restyGet(GET_RTMP_CONNECTIONS, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}


/*
GET http://localhost:9997/v2/rtmpconns/get/{name}
returns a rtmp connection
*/
func (o *mediamtxSdk) GetRtmpConnection( name string ) error{
	o.debugPrint("GetRtmpConnections ")
	// need to transform the struct
	resp, err := o.restyGet(GET_RTMP+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}

/*
POST http://localhost:9997/v2/rtmpconns/kick/{id}
kicks out a RTMP connection from the server.
*/

func (o *mediamtxSdk) KickRtmpConnection( name string ) error {
	o.debugPrint("KickRtmpConnection ")
	// need to transform the struct
	resp, err := o.restyGet(KICK_RTMP_CONNECTION+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil

}