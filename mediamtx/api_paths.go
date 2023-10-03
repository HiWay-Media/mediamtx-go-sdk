package mediamtx

/*
GET http://localhost:9997/v2/paths/list
returns all paths.
QUERY PARAMETERS
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/

func (o *mediamtxSdk) GetPathList(requestQuery ListRequest) error {
	o.debugPrint("GetPathList ")
	resp, err := o.restyGet(GET_PATHS, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}


/*
GET http://localhost:9997/v2/paths/get/{name}
returns a path.
*/