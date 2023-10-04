package mediamtx

import (

	"gopkg.in/validator.v2"
)

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
	if errs := validator.Validate(requestQuery); errs != nil {
		// values not valid, deal with errors here
		return errs
	}
	// need to transform the struct
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

func (o *mediamtxSdk) GetPath(name string) error {
	o.debugPrint("GetPath ")
	// need to transform the struct
	resp, err := o.restyGet(GET_PATH+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}
