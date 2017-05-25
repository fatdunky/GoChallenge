package main

import (
    "net/http"
    "controller"
    "utilities/logging"
    "utilities/config"
)
/*
* Log the request and Send the request to the main controller 
*/
func main() {
	logging.SetLogLevel(logging.TRACE, true)
	configuration,_ := config.LoadConfig("config/config.json")
	logging.LogToGCould = configuration.LogToCloud
	logging.LogToLocal = configuration.LogToLocal
	logging.Startedf(nil, "StanChallenge", "init", "Started")
    http.HandleFunc("/", logging.LogReceivedThenHandle(controller.MainHandler, "StanChallenge", "init"))
    logging.Error(nil,http.ListenAndServe(":8080", nil))    
}




