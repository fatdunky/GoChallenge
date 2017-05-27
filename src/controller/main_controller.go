package controller 

import (
    "net/http"
    "fmt"
    "utilities/logging"
    "encoding/json"
    "view"
    "model"
    "golang.org/x/net/context"
    "google.golang.org/appengine"
)
/* Main Handler: Handles all requests comming into app. It is the main controller. (This app only handles one route).
*  				 It parses the expected JSON request, send the parsed response to the logic class, then encodes 
				 and sends the response. 
				 It will send a 400 status code with error message when an error is encountered.
*/
func MainHandler(w http.ResponseWriter, r *http.Request) {
	 
	 var ctx context.Context
	 if (logging.LogToGCould == true) {
		 ctx = appengine.NewContext(r) 
	 } else {
	 	ctx = nil
	 }
    
    logging.Startedf(ctx,"main_controller", "MainHandler","")
	logging.Info(ctx,"Parsing incomming request")
	req,err := view.ParseRequest(ctx,r)
    if err != nil {
    	returnCode400(ctx,w,r,err)    	
    	return
    } 
   	logging.Trace(ctx,"%s",req)
   	
   	logging.Info(ctx,"Sending parsed request for proccessing")
   	resp,err := model.ProcessShows(ctx,req)
   	if err != nil {
    	returnCode400(ctx,w,r,err)    	
    	return
    } 
   	logging.Info(ctx,"Sending response")
   	 //On inital testing the other end didnt pick up the replys as JSON
   	//message, _ := json.Marshal(resp)
   	w.Header().Set("Content-Type", "application/json")
   	//w.Write(message)
    enc := json.NewEncoder(w)
    enc.Encode(resp)
    
    logging.Completed(ctx,"main_controller", "MainHandler")
}

/* Function to send 400 bad status and error message */
func returnCode400(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	 logging.Startedf(ctx,"main_controller", "returnCode400","")
	 
	 errResp := view.ErrorResponse{ ErrorMessage : fmt.Sprintf("Could not decode request: %s", err) }
	 //On inital testing the other end didnt pick up the replys as JSON
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusBadRequest)
	 //errorMessage, _ := json.Marshal(errResp)
	 //http.Error(w,string(errorMessage), http.StatusBadRequest)
	 json.NewEncoder(w).Encode(errResp)
	 //w.Header().Set("Content-Type", "application/json; charset=utf-8")
	 //w.Write(errorMessage)

	
	 logging.Warning(ctx,"Sent 400 response to client, response = %s",errResp)
		 
	 logging.Completed(ctx,"main_controller", "returnCode400")
 }
