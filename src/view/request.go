package view

import (
	"net/http"
    "utilities/logging"
    "encoding/json"
    "time"
)


/* 
*	This package contains the JSON config for incomming requests.
*/

type JSONTime struct {
    time.Time
}

const jsonTimeLayout = "1970/01/01|00:00:00"

//This should parse the time JSON objects. The examples were all null, so unsure of the format. 
func (t JSONTime)MarshalJSON(b []byte) (err error) {
    s := string(b)
    if s == "null" || s == "" {
       t.Time = time.Time{}
       return
    }
    t.Time, err = time.Parse(jsonTimeLayout, s)
    return
}

type Season struct {
	Slug			string			`json:"slug,omitempty"`
}

type Image struct {
	ShowImage		string			`json:"showImage,omitempty"`
}

type NextEpisode struct {
	Channel			int				`json:"channel,omitempty"`
	ChannelLogo		string			`json:"channelLogo,omitempty"`
	Date			JSONTime		`json:"date,omitempty"`
	Html			string			`json:"html,omitempty"`
	Url				string			`json:"url,omitempty"`
}
type Payload struct {
	Country			string			`json:"country,omitempty"`
	Description		string			`json:"desription,omitempty"`
	Drm				bool			`json:"drm,omitempty"`
	EpisodeCount	int				`json:"episodeCount,omitempty"`
	Genre			string			`json:"genre,omitempty"`
	ImageObj		Image			`json:"image"`
	Language		string			`json:"language,omitempty"`
	NextEpisodes	NextEpisode		`json:"nextEpisode,omitempty"`
	PrimaryColour	string			`json:"primaryColour,omitempty"`
	Seasons			[]Season		`json:"seasons,omitempty"`
	Slug			string			`json:"slug"`
	Title			string			`json:"title"`
	TvChannel		string			`json:"tvChannel,omitempty"`
}


type Request struct {
	Payloads 		[]Payload		`json:"payload,omitempty"`
	Skip  			int 			`json:"skip,omitempty"`
	Take       		int    			`json:"take,omitempty"`
	TotalRecords 	int 			`json:"totalRecords,omitempty"`
}

/*
* ParseRequest: This function will parse the JSON out of the incomming request and return the response struct
*/
func ParseRequest(r *http.Request) (Request,error){
	logging.Started("request", "ParseRequest")
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req); 
	if err != nil {
		logging.CompletedErrorf(err,"request", "ParseRequest","error encountered decoding request body")
		return req,err
	}
	logging.Completed("request", "ParseRequest")
	return req,nil
}

