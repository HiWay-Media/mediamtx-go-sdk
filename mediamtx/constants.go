package mediamtx

const (
	// documentation https://bluenviron.github.io/mediamtx/
	// Configs
	GET_CONFIG 					= "/v2/config/get"
	SET_CONFIG 					= "/v2/config/set" // post
	ADD_PATH_CONFIG 			= "/v2/config/paths/add/" // post
	EDIT_PATH_CONFIG			= "/v2/config/paths/edit/" // post
	REMOVE_PATH_CONFIG			= "/v2/config/paths/remove/" // post
	// HLS
	GET_HLS						= "/v2/hlsmuxers/list"
	GET_HLS_MUXER				= "/v2/hlsmuxers/get/"
	GET_PATHS					= "/v2/paths/list"
	GET_PATH 					= "/v2/paths/get/"
	// RTSP
	GET_RTSP_LIST				= "/v2/rtspconns/list"
	GET_RTSP 					= "/v2/rtspconns/get/"
	// RTSP sessions
	GET_RTSP_SESSIONS			= "/v2/rtspsessions/list"
	GET_RTSP_SESSION 			= "/v2/rtspsessions/get/"
	KICK_RTSP_SESSION			= "/v2/rtspsessions/kick/" //post
	// RTMP 
	GET_RTMP_CONNECTIONS		= "/v2/rtmpconns/list"
	GET_RTMP 					= "/v2/rtmpconns/get/"
	KICK_RTMP_CONNECTION		= "/v2/rtmpconns/kick" // post
	// SRT 
	GET_SRT_CONNECTIONS 		= "/v2/srtconns/list"
	GET_SRT 					= "/v2/srtconns/get/"
	KICK_SRT_CONNECTION			= "/v2/srtconns/kick/" // post
	//
)