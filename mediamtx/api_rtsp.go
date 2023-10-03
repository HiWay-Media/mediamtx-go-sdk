package mediamtx

/*
GET http://localhost:9997/v2/rtspconns/list
returns all RTSP connections.
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/


/*
GET http://localhost:9997/v2/rtspconns/get/{id}
returns a RTSP connection.
*/

/*
GET http://localhost:9997/v2/rtspsessions/list
returns all RTSP sessions.

QUERY PARAMETERS
page	integer
Default: 0
page number.

itemsPerPage	integer
Default: 100
items per page.
*/


/*
POST http://localhost:9997/v2/rtspsessions/kick/{id}
kicks out a RTSP session from the server.
*/


/*
GET http://localhost:9997/v2/rtspsconns/list
returns all RTSPS connections.
*/



/*
GET http://localhost:9997/v2/rtspsconns/get/{id}
returns a RTSPS connection.
*/



/*
POST http://localhost:9997/v2/rtspssessions/kick/{id}
kicks out a RTSPS session from the server.
*/