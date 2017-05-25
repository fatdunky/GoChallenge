package logging

import (
	"sync/atomic"
)

/*
* Set up logging level variable
*/
const (
	TRACE int32 	= 1
	INFO int32 		= 2
	WARN int32 		= 3
	ERROR int32 	= 4
	CRIT int32		= 5
)

var logger loggingInfo
//Used as a flag to specify to use Google API for logging
var LogToGCould bool = false
//Used as a flag to specify to log locally or not. This was included as a work around for an
//issue encountered with Uint Tests
var LogToLocal bool = true

//Used to flag if we want to log HTTP request
var logToAccess bool = true

// Start initializes tracelog and only displays the specified logging level.
func SetLogLevel(logLevel int32, logAccess bool) {
	var logToAccessInt int32 = 0 //false
	if (logToAccess == true) {
		logToAccessInt = 1 //true
	}
	atomic.StoreInt32(&logger.logAccess, logToAccessInt )
	atomic.StoreInt32(&logger.logLevel, logLevel)
}

func GetLogLevel() int32 {
	return atomic.LoadInt32(&logger.logLevel)
}

func GetAccessLogging() bool{
	if atomic.LoadInt32(&logger.logAccess) != 0 {
		return true
	}
	return false
}