package mediamtx

/*
{
	"path": "string",
	"created": "string",
	"lastRequest": "string",
	"bytesSent": 0
}
*/

type HLSMuxer struct {
	Path string `json:"path"`
	Created string `json:"created"`
	LastRequest string `json:"lastRequest"`
	BytesSent string `json:"bytesSent"`
}