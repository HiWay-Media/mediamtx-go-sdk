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
	LogLevel                  string   `json:"logLevel"`
	LogDestinations           []string `json:"logDestinations"`
	LogFile                   string   `json:"logFile"`
	ReadTimeout               string   `json:"readTimeout"`
	WriteTimeout              string   `json:"writeTimeout"`
	WriteQueueSize            int      `json:"writeQueueSize"`
	UdpMaxPayloadSize         int      `json:"udpMaxPayloadSize"`
	ExternalAuthenticationURL string   `json:"externalAuthenticationURL"`
	Api                       bool     `json:"api"`
	ApiAddress                string   `json:"apiAddress"`
	Metrics                   bool     `json:"metrics"`
	MetricsAddress            string   `json:"metricsAddress"`
	Pprof                     bool     `json:"pprof"`
	PprofAddress              string   `json:"pprofAddress"`
	RunOnConnect              string   `json:"runOnConnect"`
	RunOnConnectRestart       bool     `json:"runOnConnectRestart"`
	RunOnDisconnect           string   `json:"runOnDisconnect"`
	Rtsp                      bool     `json:"rtsp"`
	Protocols                 []string `json:"protocols"`
	Encryption                string   `json:"encryption"`
	RtspAddress               string   `json:"rtspAddress"`
	RtspsAddress              string   `json:"rtspsAddress"`
	RtpAddress                string   `json:"rtpAddress"`
	RtcpAddress               string   `json:"rtcpAddress"`
	MulticastIPRange          string   `json:"multicastIPRange"`
	MulticastRTPPort          int      `json:"multicastRTPPort"`
	MulticastRTCPPort         int      `json:"multicastRTCPPort"`
	ServerKey                 string   `json:"serverKey"`
	ServerCert                string   `json:"serverCert"`
	AuthMethods               []string `json:"authMethods"`
	Rtmp                      bool     `json:"rtmp"`
	RtmpAddress               string   `json:"rtmpAddress"`
	RtmpEncryption            string   `json:"rtmpEncryption"`
	RtmpsAddress              string   `json:"rtmpsAddress"`
}
