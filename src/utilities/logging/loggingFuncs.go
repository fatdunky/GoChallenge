package logging

/*
*	This purpose of this package is to add extra functionaily to logging by wrapping the go "log" functions. At the moment it
*	only logs to stdout. However it will be gradually improved in the future.
*	
*	It is a modified version of tracelog from https://github.com/goinggo/tracelog
*/

import (
	"fmt"
)

// Log received messages
func AccessReceived(format string, a ...interface{}) {
	logger.Access.Output(2, fmt.Sprintf("Received : %s\n", fmt.Sprintf(format, a...)))
}


//** STARTED AND COMPLETED

// Started uses the Serialize destination and adds a Started tag to the log line
func Started(title string, functionName string) {
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Started\n", title, functionName))
}

// Startedf uses the Serialize destination and writes a Started tag to the log line
func Startedf(title string, functionName string, format string, a ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Started : %s\n", title, functionName, fmt.Sprintf(format, a...)))
}

// Completed uses the Serialize destination and writes a Completed tag to the log line
func Completed(title string, functionName string) {
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Completed\n", title, functionName))
}

// Completedf uses the Serialize destination and writes a Completed tag to the log line
func Completedf(title string, functionName string, format string, a ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Completed : %s\n", title, functionName, fmt.Sprintf(format, a...)))
}

// CompletedError uses the Error destination and writes a Completed tag to the log line
func CompletedError(err error, title string, functionName string) {
	logger.Error.Output(2, fmt.Sprintf("%s : %s : Completed : ERROR : %s\n", title, functionName, err))
}

// CompletedErrorf uses the Error destination and writes a Completed tag to the log line
func CompletedErrorf(err error, title string, functionName string, format string, a ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("%s : %s : Completed : ERROR : %s : %s\n", title, functionName, fmt.Sprintf(format, a...), err))
}

//** TRACE

// Trace writes to the Trace destination
func Trace(format string, a ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf("Trace : %s\n", fmt.Sprintf(format, a...)))
}

//** INFO

// Info writes to the Info destination
func Info(format string, a ...interface{}) {
	logger.Info.Output(2, fmt.Sprintf("Info : %s\n", fmt.Sprintf(format, a...)))
}

//** WARNING

// Warning writes to the Warning destination
func Warning(format string, a ...interface{}) {
	logger.Warning.Output(2, fmt.Sprintf("Info : %s\n", fmt.Sprintf(format, a...)))
}

//** ERROR

// Error writes to the Error destination and accepts an err
func Error(err error) {
	logger.Error.Output(2, fmt.Sprintf("ERROR : %s\n", err))
}

// Errorf writes to the Error destination and accepts an err
func Errorf(err error, format string, a ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("ERROR : %s : %s\n", fmt.Sprintf(format, a...), err))
}




