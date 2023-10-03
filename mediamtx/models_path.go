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
	Name string `json:"name"`
	ConfName string `json:"confName"`
}