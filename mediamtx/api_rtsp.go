package mediamtx

/*
GET http://localhost:9997/v2/rtspconns/list
returns all RTSP connections.
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/

func (o *mediamtxSdk) GetRtspConnections( requestQuery ListRequest ) error {
	o.debugPrint("GetRtspConnections ")
	// need to transform the struct
	resp, err := o.restyGet(GET_RTSP_LIST, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil

}

/*
GET http://localhost:9997/v2/rtspconns/get/{id}
returns a RTSP connection.
*/

func (o *mediamtxSdk) GetRtspConnection( name string ) error {
	o.debugPrint("GetRtspConnections ")
	// need to transform the struct
	resp, err := o.restyGet(GET_RTSP+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil

}

/*
GET http://localhost:9997/v2/rtspsessions/list
returns all RTSP sessions.

QUERY PARAMETERS
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/

func (o *mediamtxSdk) GetRtspSessions( requestQuery ListRequest ) error{
	o.debugPrint("GetRtspSessions ")
	// need to transform the struct
	resp, err := o.restyGet(GET_RTSP_SESSIONS, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}

/*
GET http://localhost:9997/v2/rtspsessions/get/{id}

*/

func (o *mediamtxSdk) GetRtspSession( name string ) error{
	o.debugPrint("GetRtspSession ")
	// need to transform the struct
	resp, err := o.restyGet(GET_RTSP_SESSION+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}


/*
POST http://localhost:9997/v2/rtspsessions/kick/{id}
kicks out a RTSP session from the server.
*/



func (o *mediamtxSdk) KickRtspSession( name string ) error {
	o.debugPrint("KickRtspSession ")
	// need to transform the struct
	resp, err := o.restyGet(KICK_RTSP_SESSION+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil

}

/*
GET http://localhost:9997/v2/rtspsconns/list
returns all RTSPS connections.
*/



/*
GET http://localhost:9997/v2/rtspsconns/get/{id}
returns a RTSPS connection.
*/



/*
POST http://localhost:9997/v2/rtspssessions/kick/{id}
kicks out a RTSPS session from the server.
*/

func (o *mediamtxSdk) KickRtspConnection( name string ) error {
	o.debugPrint("KickRtspConnection ")
	// need to transform the struct
	resp, err := o.restyGet(KICK_RTSP_SESSION+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil

}