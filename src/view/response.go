package view

import (
	"utilities/logging"
)

/* 
*	This package contains the JSON config for out going responses.
*/

type ErrorResponse struct {
	 ErrorMessage string `json:"error"`
}
 
type ResponseShow struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type Response struct {
	Shows []ResponseShow `json:"response"`
}
//This function adds a matched show into the reponse. This could have been done outside this package, but it felt cleaner here.
func AddShowToReponse(resp Response, image, slug, title string) Response {
	logging.Started("response", "AddShowToResponse")
	temp := ResponseShow {Image:image,Slug:slug,Title:title}
	resp.Shows = append(resp.Shows, temp)
	logging.Completed("repsonse", "AddShowToResponse")
	return resp
}


