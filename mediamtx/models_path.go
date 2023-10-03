package mediamtx

/*{
"name": "string",
"confName": "string",
"conf": {
"source": "string",
"sourceFingerprint": "string",
"sourceOnDemand": true,
"sourceOnDemandStartTimeout": "string",
"sourceOnDemandCloseAfter": "string",
"maxReaders": 0,
"record": true,
"publishUser": "string",
"publishPass": "string",
"publishIPs": [
"string"
],
"readUser": "string",
"readPass": "string",
"readIPs": [
"string"
],
"overridePublisher": true,
"fallback": "string",
"srtPublishPassphrase": "string",
"srtReadPassphrase": "string",
"sourceProtocol": "string",
"sourceAnyPortEnable": true,
"rtspRangeType": "string",
"rtspRangeStart": "string",
"sourceRedirect": "string",
"rpiCameraCamID": 0,
"rpiCameraWidth": 0,
"rpiCameraHeight": 0,
"rpiCameraHFlip": true,
"rpiCameraVFlip": true,
"rpiCameraBrightness": 0,
"rpiCameraContrast": 0,
"rpiCameraSaturation": 0,
"rpiCameraSharpness": 0,
"rpiCameraExposure": "string",
"rpiCameraAWB": "string",
"rpiCameraDenoise": "string",
"rpiCameraShutter": 0,
"rpiCameraMetering": "string",
"rpiCameraGain": 0,
"rpiCameraEV": 0,
"rpiCameraROI": "string",
"rpiCameraHDR": true,
"rpiCameraTuningFile": "string",
"rpiCameraMode": "string",
"rpiCameraFPS": 0,
"rpiCameraIDRPeriod": 0,
"rpiCameraBitrate": 0,
"rpiCameraProfile": "string",
"rpiCameraLevel": "string",
"rpiCameraAfMode": "string",
"rpiCameraAfRange": "string",
"rpiCameraAfSpeed": "string",
"rpiCameraLensPosition": 0,
"rpiCameraAfWindow": "string",
"rpiCameraTextOverlayEnable": true,
"rpiCameraTextOverlay": "string",
"runOnInit": "string",
"runOnInitRestart": true,
"runOnDemand": "string",
"runOnDemandRestart": true,
"runOnDemandStartTimeout": "string",
"runOnDemandCloseAfter": "string",
"runOnReady": "string",
"runOnReadyRestart": true,
"runOnNotReady": "string",
"runOnRead": "string",
"runOnReadRestart": true,
"runOnUnread": "string"
},
"source": {
"type": "hlsMuxer",
"id": "string"
},
"ready": true,
"readyTime": "string",
"tracks": [
"string"
],
"bytesReceived": 0,
"readers": [
{
"type": "hlsMuxer",
"id": "string"
}
]
}*/

type Path struct {
	Name     string `json:"name"`
	ConfName string `json:"confName"`
	Conf     struct {
		Source                     string   `json:"source"`
		SourceFingerprint          string   `json:"sourceFingerprint"`
		SourceOnDemand             bool     `json:"sourceOnDemand"`
		SourceOnDemandStartTimeout string   `json:"sourceOnDemandStartTimeout"`
		SourceOnDemandCloseAfter   string   `json:"sourceOnDemandCloseAfter"`
		MaxReaders                 int      `json:"maxReaders"`
		Record                     bool     `json:"record"`
		PublishUser                string   `json:"publishUser"`
		PublishPass                string   `json:"publishPass"`
		PublishIPs                 []string `json:"publishIPs"`
		ReadUser                   string   `json:"readUser"`
		ReadPass                   string   `json:"readPass"`
		ReadIPs                    []string `json:"readIPs"`
		OverridePublisher          bool     `json:"overridePublisher"`
		Fallback                   string   `json:"fallback"`
		SrtPublishPassphrase       string   `json:"srtPublishPassphrase"`
		SrtReadPassphrase          string   `json:"srtReadPassphrase"`
		SourceProtocol             string   `json:"sourceProtocol"`
		SourceAnyPortEnable        bool     `json:"sourceAnyPortEnable"`
		RtspRangeType              string   `json:"rtspRangeType"`
		RtspRangeStart             string   `json:"rtspRangeStart"`
		SourceRedirect             string   `json:"sourceRedirect"`
		RpiCameraCamID             int      `json:"rpiCameraCamID"`
		RpiCameraWidth             int      `json:"rpiCameraWidth"`
		RpiCameraHeight            int      `json:"rpiCameraHeight"`
		RpiCameraHFlip             bool     `json:"rpiCameraHFlip"`
		RpiCameraVFlip             bool     `json:"rpiCameraVFlip"`
		RpiCameraBrightness        int      `json:"rpiCameraBrightness"`
		RpiCameraContrast          int      `json:"rpiCameraContrast"`
		RpiCameraSaturation        int      `json:"rpiCameraSaturation"`
		RpiCameraSharpness         int      `json:"rpiCameraSharpness"`
		RpiCameraExposure          string   `json:"rpiCameraExposure"`
		RpiCameraAWB               string   `json:"rpiCameraAWB"`
		RpiCameraDenoise           string   `json:"rpiCameraDenoise"`
		RpiCameraShutter           int      `json:"rpiCameraShutter"`
		RpiCameraMetering          string   `json:"rpiCameraMetering"`
		RpiCameraGain              int      `json:"rpiCameraGain"`
		RpiCameraEV                int      `json:"rpiCameraEV"`
		RpiCameraROI               string   `json:"rpiCameraROI"`
		RpiCameraHDR               bool     `json:"rpiCameraHDR"`
		RpiCameraTuningFile        string   `json:"rpiCameraTuningFile"`
		RpiCameraMode              string   `json:"rpiCameraMode"`
		RpiCameraFPS               int      `json:"rpiCameraFPS"`
		RpiCameraIDRPeriod         int      `json:"rpiCameraIDRPeriod"`
		RpiCameraBitrate           int      `json:"rpiCameraBitrate"`
		RpiCameraProfile           string   `json:"rpiCameraProfile"`
		RpiCameraLevel             string   `json:"rpiCameraLevel"`
		RpiCameraAfMode            string   `json:"rpiCameraAfMode"`
		RpiCameraAfRange           string   `json:"rpiCameraAfRange"`
		RpiCameraAfSpeed           string   `json:"rpiCameraAfSpeed"`
		RpiCameraLensPosition      int      `json:"rpiCameraLensPosition"`
		RpiCameraAfWindow          string   `json:"rpiCameraAfWindow"`
		RpiCameraTextOverlayEnable bool     `json:"rpiCameraTextOverlayEnable"`
		RpiCameraTextOverlay       string   `json:"rpiCameraTextOverlay"`
		RunOnInit                  string   `json:"runOnInit"`
		RunOnInitRestart           bool     `json:"runOnInitRestart"`
		RunOnDemand                string   `json:"runOnDemand"`
		RunOnDemandRestart         bool     `json:"runOnDemandRestart"`
		RunOnDemandStartTimeout    string   `json:"runOnDemandStartTimeout"`
		RunOnDemandCloseAfter      string   `json:"runOnDemandCloseAfter"`
		RunOnReady                 string   `json:"runOnReady"`
		RunOnReadyRestart          bool     `json:"runOnReadyRestart"`
		RunOnNotReady              string   `json:"runOnNotReady"`
		RunOnRead                  string   `json:"runOnRead"`
		RunOnReadRestart           bool     `json:"runOnReadRestart"`
		RunOnUnread                string   `json:"runOnUnread"`
	} `json:"conf"`
	Source struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	} `json:"source"`
	Ready         bool     `json:"ready"`
	ReadyTime     string   `json:"readyTime"`
	Tracks        []string `json:"tracks"`
	BytesReceived int      `json:"bytesReceived"`
	Readers       []struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	} `json:"readers"`
}

type PathListResponse struct {
	PageCount int    `json:"pageCount"`
	Items     []Path `json:"items"`
}
