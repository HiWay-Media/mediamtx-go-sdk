package mediamtx

import (

	"gopkg.in/validator.v2"
)

/*
GET http://localhost:9997/v2/hlsmuxers/list
returns all HLS muxers.
QUERY PARAMETERS
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/

func (o *mediamtxSdk) GetHlsMuxers( requestQuery ListRequest ) error {
	o.debugPrint("GetPathList ")
	if errs := validator.Validate(requestQuery); errs != nil {
		// values not valid, deal with errors here | return nil, errs
		return errs
	}
	// need to transform the struct
	resp, err := o.restyGet(GET_HLS, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}


/*
GET http://localhost:9997/v2/hlsmuxers/get/{name}
returns a HLS muxer.
*/