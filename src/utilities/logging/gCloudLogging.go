package logging

import (
	"fmt"
    "google.golang.org/appengine/log"
    "golang.org/x/net/context"
)

func GlogDebugf(ctx context.Context, format string, a ...interface{}) {
	if GetLogLevel() <= TRACE && ctx != nil {
		log.Debugf(ctx, fmt.Sprintf("%s", fmt.Sprintf(format, a...)))
	}
}

func GlogInfof(ctx context.Context, format string, a ...interface{}) {
	if GetLogLevel() <= INFO && ctx != nil {
		log.Infof(ctx, fmt.Sprintf("%s", fmt.Sprintf(format, a...)))
	}
}
func GlogWarningf(ctx context.Context, format string, a ...interface{}) {
	if GetLogLevel() <= WARN && ctx != nil {
		log.Warningf(ctx, fmt.Sprintf("%s", fmt.Sprintf(format, a...)))
	}
}
func GlogErrorf(ctx context.Context, format string, a ...interface{}) {
	if GetLogLevel() <= ERROR && ctx != nil {
		log.Errorf(ctx, fmt.Sprintf("%s", fmt.Sprintf(format, a...)))
	}
}
func GlogCriticalf(ctx context.Context, format string, a ...interface{}) {
	if GetLogLevel() <= CRIT && ctx != nil {
		log.Criticalf(ctx, fmt.Sprintf("%s", fmt.Sprintf(format, a...)))
	}
}


