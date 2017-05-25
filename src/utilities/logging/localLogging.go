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
    "fmt"
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
	Critical           *log.Logger

}

var lLogger loggingInfo


func LlogTracef(format string, a ...interface{}) {
	if GetLogLevel() <= TRACE {
		logger.Trace.Output(2, fmt.Sprintf("Trace : %s", fmt.Sprintf(format, a...)))
	}
}

func LlogInfof(format string, a ...interface{}) {
	if GetLogLevel() <= INFO {
		logger.Info.Output(2, fmt.Sprintf("Info : %s", fmt.Sprintf(format, a...)))
	}
}
func LlogWarningf(format string, a ...interface{}) {
	if GetLogLevel() <= WARN {
		logger.Warning.Output(2, fmt.Sprintf("Info : %s", fmt.Sprintf(format, a...)))
	}
}
func LlogErrorf(format string, a ...interface{}) {
	if GetLogLevel() <= ERROR {
		logger.Error.Output(2, fmt.Sprintf("ERROR : %s", fmt.Sprintf(format, a...)))
	}
}
func LlogCriticalf(format string, a ...interface{}) {
	if GetLogLevel() <= CRIT {
		logger.Critical.Output(2,fmt.Sprintf("%s", fmt.Sprintf(format, a...)))
	}
}

// Called to init the logging system.
func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	
	logger.Trace = 	log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = 	log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = 	log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Critical = 	log.New(os.Stderr, "CRITICAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Access = log.New(os.Stdout, "ACCESS: ", log.Ldate|log.Ltime|log.Lshortfile)
}



