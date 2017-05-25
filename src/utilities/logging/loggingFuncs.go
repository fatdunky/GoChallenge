package logging

/*
*	This purpose of this package is to add extra functionaily to logging by wrapping the go "log" functions. At the moment it
*	only logs to stdout. However it will be gradually improved in the future.
*	
*	It is a modified version of tracelog from https://github.com/goinggo/tracelog
*/

import (
	"fmt"
    "golang.org/x/net/context"
)

type logLocalMessage func(format string, a ...interface{})
type logGCloudMessage func(ctx context.Context, format string, a ...interface{})

func logMessage(ctx context.Context,localMessage logLocalMessage, gCloudMessage logGCloudMessage, format string, a ...interface{}) {
	message :=  fmt.Sprintf(format,a...)
	if LogToLocal == true {
		localMessage(message)
	}
	if LogToGCould == true && ctx != nil {
		gCloudMessage(ctx, message)
	}
}

// Log received messages
func AccessReceived(ctx context.Context, format string, a ...interface{}) {
	if LogToLocal == true {
		logger.Access.Output(2, fmt.Sprintf("Received : %s\n", fmt.Sprintf(format, a...)))
	} 
	//No need to log access in gcloud as apengine does it allready
}

//** STARTED AND COMPLETED

// Started uses the Serialize destination and adds a Started tag to the log line
/*func Started(ctx context.Context, title string, functionName string, functionName string) {
	message :=  fmt.Sprintf("%s : %s : Started\n", title, functionName)
	logMessage(ctx, LlogTracef,GlogDebugf ,message)
}*/

// Startedf uses the Serialize destination and writes a Started tag to the log line
func Startedf(ctx context.Context, title string, functionName string, format string, a ...interface{}) {
	message :=  fmt.Sprintf("%s : %s : Started : %s\n", title, functionName, fmt.Sprintf(format, a...))
	logMessage(ctx, LlogTracef,GlogDebugf ,message)
}

// Completed uses the Serialize destination and writes a Completed tag to the log line
func Completed(ctx context.Context, title string, functionName string) {
	message :=  fmt.Sprintf("%s : %s : Completed\n", title, functionName)
	logMessage(ctx, LlogTracef,GlogDebugf ,message)
}

// Completedf uses the Serialize destination and writes a Completed tag to the log line
func Completedf(ctx context.Context, title string, functionName string, format string, a ...interface{}) {
	message :=  fmt.Sprintf("%s : %s : Completed : %s\n", title, functionName, fmt.Sprintf(format, a...))
	logMessage(ctx, LlogTracef,GlogDebugf ,message)
}

// CompletedError uses the Error destination and writes a Completed tag to the log line
func CompletedError(ctx context.Context, err error, title string, functionName string) {
	message :=  fmt.Sprintf("%s : %s : Completed : ERROR : %s\n", title, functionName, err)
	logMessage(ctx, LlogErrorf,GlogErrorf ,message)
	
}

// CompletedErrorf uses the Error destination and writes a Completed tag to the log line
func CompletedErrorf(ctx context.Context, err error, title string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : Completed : ERROR : %s : %s\n", title, functionName, fmt.Sprintf(format, a...), err)
	logMessage(ctx, LlogErrorf,GlogErrorf ,message)
}

//** TRACE

// Trace writes to the Trace destination
func Trace(ctx context.Context, format string, a ...interface{}) {
	message := fmt.Sprintf("Trace : %s\n", fmt.Sprintf(format, a...))
	logMessage(ctx, LlogTracef,GlogDebugf ,message)
}

//** INFO

// Info writes to the Info destination
func Info(ctx context.Context, format string, a ...interface{}) {
	message := fmt.Sprintf("Info : %s\n", fmt.Sprintf(format, a...))
	logMessage(ctx, LlogInfof,GlogInfof ,message)
}

//** WARNING

// Warning writes to the Warning destination
func Warning(ctx context.Context, format string, a ...interface{}) {
	message := fmt.Sprintf("Info : %s\n", fmt.Sprintf(format, a...))
	logMessage(ctx, LlogWarningf,GlogWarningf ,message)
}

//** ERROR

// Error writes to the Error destination and accepts an err
func Error(ctx context.Context, err error) {
	message := fmt.Sprintf("ERROR : %s\n", err)
	logMessage(ctx, LlogErrorf,GlogErrorf ,message)
}

// Errorf writes to the Error destination and accepts an err
func Errorf(ctx context.Context, err error, format string, a ...interface{}) {
	message := fmt.Sprintf("ERROR : %s : %s\n", fmt.Sprintf(format, a...), err)
	logMessage(ctx, LlogErrorf,GlogErrorf ,message)
}

// Criticalf writes to the Error destination and accepts an err
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	message := fmt.Sprintf("CRITICAL : %s : %s\n", fmt.Sprintf(format, a...))
	logMessage(ctx, LlogCriticalf,GlogCriticalf ,message)
}




