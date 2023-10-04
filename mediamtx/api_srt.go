package mediamtx



/*
GET http://localhost:9997/v2/srtconns/list
returns all RTSP connections.
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/
func (o *mediamtxSdk) GetSrtConnections( requestQuery ListRequest ) error {
	o.debugPrint("GetSrtConnections ")
	// need to transform the struct
	resp, err := o.restyGet(GET_SRT_CONNECTIONS, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}