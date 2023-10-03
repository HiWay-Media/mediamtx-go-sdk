package mediamtx

/*
{
"logLevel": "string",
"logDestinations": [
"string"
],
"logFile": "string",
"readTimeout": "string",
"writeTimeout": "string",
"writeQueueSize": 0,
"udpMaxPayloadSize": 0,
"externalAuthenticationURL": "string",
"api": true,
"apiAddress": "string",
"metrics": true,
"metricsAddress": "string",
"pprof": true,
"pprofAddress": "string",
"runOnConnect": "string",
"runOnConnectRestart": true,
"runOnDisconnect": "string",
"rtsp": true,
"protocols": [
"string"
],
"encryption": "string",
"rtspAddress": "string",
"rtspsAddress": "string",
"rtpAddress": "string",
"rtcpAddress": "string",
"multicastIPRange": "string",
"multicastRTPPort": 0,
"multicastRTCPPort": 0,
"serverKey": "string",
"serverCert": "string",
"authMethods": [
"string"
],
"rtmp": true,
"rtmpAddress": "string",
"rtmpEncryption": "string",
"rtmpsAddress": "string",
"rtmpServerKey": "string",
"rtmpServerCert": "string",
"hls": true,
"hlsAddress": "string",
"hlsEncryption": true,
"hlsServerKey": "string",
"hlsServerCert": "string",
"hlsAlwaysRemux": true,
"hlsVariant": "string",
"hlsSegmentCount": 0,
"hlsSegmentDuration": "string",
"hlsPartDuration": "string",
"hlsSegmentMaxSize": "string",
"hlsAllowOrigin": "string",
"hlsTrustedProxies": [
"string"
],
"hlsDirectory": "string",
"webrtc": true,
"webrtcAddress": "string",
"webrtcEncryption": true,
"webrtcServerKey": "string",
"webrtcServerCert": "string",
"webrtcAllowOrigin": "string",
"webrtcTrustedProxies": [
"string"
],
"webrtcICEServers2": [
{}
],
"webrtcICEHostNAT1To1IPs": [
"string"
],
"webrtcICEUDPMuxAddress": "string",
"webrtcICETCPMuxAddress": "string",
"srt": true,
"srtAddress": "string",
"record": true,
"recordPath": "string",
"recordFormat": "string",
"recordPartDuration": "string",
"recordSegmentDuration": "string",
"recordDeleteAfter": "string",
"paths": {
"property1": {},
"property2": {}
}
}
*/

type Config struct {
	LogLevel string `json:"logLevel"`
	
} 