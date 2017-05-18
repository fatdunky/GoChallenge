package model

import (
    "utilities/logging"
    "view"
    "errors"
)

/*
* ProcessShows: Handles the main business logic for the application. It will throw an error if not all the variables can be
*				returned on a matched show, not sure if this is the correct behaviour. 
*/

func ProcessShows(req view.Request)(view.Response,error) {
	/*
	From the list of shows in the request payload, return the ones with DRM enabled (drm: true) 
	and at least one episode (episodeCount > 0)
	
	The returned JSON should have a response key with an array of shows. 
	Each element should have the following fields from the request:

	image - corresponding to image/showImage from the request payload
	slug
	title
	*/
	retVal := view.Response{}
	for _, playload := range req.Payloads {
		logging.Started("logic", "processShows")
		title := playload.Title
		slug := playload.Slug
		drm := playload.Drm
		epCount := playload.EpisodeCount
		image := playload.ImageObj.ShowImage
		
		logging.Trace("Title:%s, drm:%s. episodeCount:%s",title,drm,epCount)
		if (drm == true && epCount > 0 ) {
			if title == "" ||  slug == "" || image == "" {
				err := errors.New("Show does not contain enough information to return")
				logging.CompletedErrorf(err,"logic", "processShows","title(%s), slug(%s), image(%s)",title,slug,image)
				return retVal, err
			}
			logging.Info("Adding playload to response, title(%s), slug(%s), image(%s)",title,slug,image)
			retVal = view.AddShowToReponse(retVal,title,slug,image)
		}
	}
	logging.Completed("logic", "processShows")
	return retVal,nil
}
