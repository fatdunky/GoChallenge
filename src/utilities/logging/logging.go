package logging

/*
*	This purpose of this package is to add extra functionaily to logging by wrapping the go "log" functions. At the moment it
*	only logs to stdout. However it will be gradually improved in the future.
*	
*	It is a modified version of tracelog from https://github.com/goinggo/tracelog
*/
import (
    "log"
    "os"
    "sync/atomic"
    "io/ioutil"
)

/*
* Set up logging level variable
*/
const (
	TRACE int32 	= 1
	INFO int32 		= 2
	WARN int32 		= 3
	ERROR int32 	= 4
)

/*
* Struct for outside packages to log through
* The access value is seperate from the other levels. 
* This is so you could have the logging level at error level but still have access logs
* Used an Int rather then bool, for thread saftey
*/
type loggingInfo struct {
	logLevel           int32
	logAccess		   int32	
	Access             *log.Logger
	Trace              *log.Logger
	Info               *log.Logger
	Warning            *log.Logger
	Error              *log.Logger

}

var logger loggingInfo


// Called to init the logging system.
func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// Start initializes tracelog and only displays the specified logging level.
func SetLogLevel(logLevel int32, logAccess bool) {
	turnOnLogging(logLevel, logAccess)
}

// turnOnLogging configures the logging writers.
func turnOnLogging(logLevel int32, logToAccess bool) {
	accessHandle := ioutil.Discard
	traceHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warnHandle := ioutil.Discard
	errorHandle := ioutil.Discard

	if logLevel&TRACE != 0 {
		traceHandle = os.Stdout
		infoHandle = os.Stdout
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&INFO != 0 {
		infoHandle = os.Stdout
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&WARN != 0 {
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&ERROR != 0 {
		errorHandle = os.Stderr
	}

	if logToAccess == true {
		accessHandle = os.Stdout
	}
	
	logger.Trace = 	log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = 	log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(warnHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = 	log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Access = log.New(accessHandle, "ACCESS: ", log.Ldate|log.Ltime|log.Lshortfile)

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

