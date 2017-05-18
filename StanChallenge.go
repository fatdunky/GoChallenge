package main

import (
    "net/http"
    "controller"
    "utilities/logging"
)
/*
* Log the request and Send the request to the main controller 
*/
func main() {
	logging.SetLogLevel(logging.TRACE, true)
	logging.Started("StanChallenge", "init")
    http.HandleFunc("/", logging.LogReceivedThenHandle(controller.MainHandler, "StanChallenge", "init"))
    logging.Error(http.ListenAndServe(":8080", nil))    
}




