package controller 

import (
    "net/http"
    "fmt"
    "utilities/logging"
    "encoding/json"
    "view"
    "model"
)
/* Main Handler: Handles all requests comming into app. It is the main controller. (This app only handles one route).
*  				 It parses the expected JSON request, send the parsed response to the logic class, then encodes 
				 and sends the response. 
				 It will send a 400 status code with error message when an error is encountered.
*/
func MainHandler(w http.ResponseWriter, r *http.Request) {
    logging.Started("main_controller", "MainHandler")
	logging.Info("Parsing incomming request")
	req,err := view.ParseRequest(r)
    if err != nil {
    	returnCode400(w,r,err)    	
    	return
    } 
   	logging.Trace("%s",req)
   	
   	logging.Info("Sending parsed request for proccessing")
   	resp,err := model.ProcessShows(req)
   	if err != nil {
    	returnCode400(w,r,err)    	
    	return
    } 
   	logging.Info("Sending response")
    enc := json.NewEncoder(w)
    enc.Encode(resp)
    
    logging.Completed("main_controller", "MainHandler")
}

/* Function to send 400 bad status and error message */
func returnCode400(w http.ResponseWriter, r *http.Request, err error) {
	 logging.Started("main_controller", "returnCode400")
	 
	 errResp := view.ErrorResponse{ ErrorMessage : fmt.Sprintf("Could not decode request: %s", err) }
	
	 w.WriteHeader(http.StatusBadRequest)
	 json.NewEncoder(w).Encode(errResp)
	 logging.Warning("Sent 400 response to client, response = %s",errResp)
		 
	 logging.Completed("main_controller", "returnCode400")
 }
