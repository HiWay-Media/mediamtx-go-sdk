package mediamtx

import (
	"encoding/json"
)

/*
GET http://localhost:9997/v2/config/get
returns the configuration.
*/

func(o *mediamtxSdk) GetConfig() (*Config, error)  {
	//log.Println("GetConfig ")
	o.debugPrint("GetConfig")
	resp, err := o.restyGet(GET_CONFIG, nil)
	if err != nil {
		return nil, err
	}
	//
	o.debugPrint(resp)
	var obj Config
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

/*
POST http://localhost:9997/v2/config/set
changes the configuration.
all fields are optional. paths can't be edited with this request, use /v2/config/paths/{operation}/{name} to edit them.
*/
func (o *mediamtxSdk) SetConfig(name string) error {
	o.debugPrint("SetConfig "+ name)
	resp, err := o.restyPost(SET_CONFIG+name, nil)
	if err != nil {
		return err
	}
	//
	o.debugPrint(resp)
	return nil
}


/*
POST http://localhost:9997/v2/config/paths/add/{name}

adds the configuration of a path.
all fields are optional.
*/

func (o *mediamtxSdk) AddPathConfig(name string) error {
	o.debugPrint("AddPathConfig "+ name)
	resp, err := o.restyPost(ADD_PATH_CONFIG+name, nil)
	if err != nil {
		return err
	}
	// it miss the request
	//
	o.debugPrint(resp)
	return nil
}

/*
POST http://localhost:9997/v2/config/paths/edit/{name}

changes the configuration of a path.
all fields are optional.
*/

func (o *mediamtxSdk) EditPathConfig(name string) error {
	o.debugPrint("EditPathConfig "+ name)
	resp, err := o.restyPost(EDIT_PATH_CONFIG+name, nil)
	if err != nil {
		return err
	}
	// it miss the request
	//
	o.debugPrint(resp)
	return nil
}



/*
POST http://localhost:9997/v2/config/paths/remove/{name}
removes the configuration of a path.
*/


func (o *mediamtxSdk) RemovePathConfig(name string) error{
	o.debugPrint("EditPathConfig "+ name)
	resp, err := o.restyPost(EDIT_PATH_CONFIG+name, nil)
	if err != nil {
		return err
	}
	// it miss the request
	//
	o.debugPrint(resp)
	return nil
}