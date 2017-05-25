package tests

import (
	"testing"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"controller"
	"strings"
	"encoding/json"
	"view"
)



func TestEmptyRequest(t *testing.T) {
	
	reqBody := strings.NewReader("{}")
	req := httptest.NewRequest("GET", "/", reqBody)
	w := httptest.NewRecorder()
	controller.MainHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//Empty request (in JSON format) should be okay. 
	if resp.StatusCode != 200 {
		 t.Error("Empty Request: Expected status code of 200 got:",resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		  t.Error("Empty Request: Expected Content-Type of 'application/json' got:",resp.Header.Get("Content-Type"))
	}
	var resStruct view.ErrorResponse
	
	err := json.NewDecoder(resp.Body).Decode(&resStruct); 

	if err == nil {
		 t.Error("Empty Request: returned error", err)
	}
	if resStruct.ErrorMessage != "" {
		 t.Error("Empty Request: errorResponse is not null")
	}
}

func TestMinimumRequest(t *testing.T) {
	
	reqBody := strings.NewReader("{ \"payload\": [{ \"title\": \"Test Show\", \"drm\": true, \"episodeCount\": 3, \"slug\": \"show/test\", \"image\" : {\"showImage\":\"http://testImage\"}}]}")
	req := httptest.NewRequest("GET", "/", reqBody)
	w := httptest.NewRecorder()
	controller.MainHandler(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//Empty request (in JSON format) should be okay. 
	if resp.StatusCode != 200 {
		 t.Error("Miniumum Request: Expected status code of 200 got:",resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		  t.Error("Miniumum Request: Expected Content-Type of 'application/json' got:",resp.Header.Get("Content-Type"))
	}
	var resStruct view.Response
	
	err := json.NewDecoder(resp.Body).Decode(&resStruct); 
	if err != nil {
		 t.Error("Miniumum Request: returned error", err)
	}
	
	if len(resStruct.Shows) != 1 {
		t.Error("Miniumum Request:, test returned result. Expect 1 shows returned got", len(resStruct.Shows))
	}
	
	for _, show := range resStruct.Shows {
		if show.Title == "" {
			 t.Error("Miniumum Request: title is not null")
		}
		if show.Slug == "" {
			 t.Error("Miniumum Request: slug is not null")
		}
		if show.Image == "" {
			 t.Error("Miniumum Request: Image is not null")
		}
	} 
	
}

func TestBadStringRequest(t *testing.T) {
	
	reqBody := strings.NewReader("{ \"payload\": [{ \"title\": \"Test Show\", \"drm\": true, \"episodeCount\": 3, \"genre\": genre ,\"slug\": \"show/test\", \"image\" : {\"showImage\":\"http://testImage\"}}]}")
	req := httptest.NewRequest("GET", "/", reqBody)
	w := httptest.NewRecorder()
	controller.MainHandler(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//Empty request (in JSON format) should be okay. 
	if resp.StatusCode != 400 {
		 t.Error("Bad String Request: Expected status code of 400 got:",resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		  t.Error("Bad String Request: Expected Content-Type of 'application/json' got:",resp.Header.Get("Content-Type"))
	}
	var resStruct view.ErrorResponse
	
	err := json.NewDecoder(resp.Body).Decode(&resStruct); 
	if err != nil {
		 t.Error("Bad String Request: returned error", err)
	}
	
	if resStruct.ErrorMessage == "" {
		t.Error("Bad String Request: ErrorMessage == nil. ", resStruct)
	}
	
}

func TestBadBooleanRequest(t *testing.T) {
	
	reqBody := strings.NewReader("{ \"payload\": [{ \"title\": \"Test Show\", \"drm\": incorrect, \"episodeCount\": 3,\"slug\": \"show/test\", \"image\" : {\"showImage\":\"http://testImage\"}}]}")
	req := httptest.NewRequest("GET", "/", reqBody)
	w := httptest.NewRecorder()
	controller.MainHandler(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//Empty request (in JSON format) should be okay. 
	if resp.StatusCode != 400 {
		 t.Error("Bad Boolean Request: Expected status code of 400 got:",resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		  t.Error("Bad Boolean Request: Expected Content-Type of 'application/json' got:",resp.Header.Get("Content-Type"))
	}
	var resStruct view.ErrorResponse
	
	err := json.NewDecoder(resp.Body).Decode(&resStruct); 
	if err != nil {
		 t.Error("Bad Boolean Request: returned error", err)
	}
	
	if resStruct.ErrorMessage == "" {
		t.Error("Bad Boolean Request: ErrorMessage == nil. ", resStruct)
	}
	
}
func TestBadIntegerRequest(t *testing.T) {
	
	reqBody := strings.NewReader("{ \"payload\": [{ \"title\": \"Test Show\", \"drm\": true, \"episodeCount\": \"incorrect\",\"slug\": \"show/test\", \"image\" : {\"showImage\":\"http://testImage\"}}]}")
	req := httptest.NewRequest("GET", "/", reqBody)
	w := httptest.NewRecorder()
	controller.MainHandler(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//Empty request (in JSON format) should be okay. 
	if resp.StatusCode != 400 {
		 t.Error("Bad Integer Request: Expected status code of 400 got:",resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		  t.Error("Bad Integer Request: Expected Content-Type of 'application/json' got:",resp.Header.Get("Content-Type"))
	}
	var resStruct view.ErrorResponse
	
	err := json.NewDecoder(resp.Body).Decode(&resStruct); 
	if err != nil {
		 t.Error("Bad Integer Request: returned error", err)
	}
	
	if resStruct.ErrorMessage == "" {
		t.Error("Bad Integer Request: ErrorMessage == nil. ", resStruct)
	}
	
}

func TestFullRequest(t *testing.T) {
	
	reqBody := strings.NewReader("{ \"payload\": [{ \"country\":\"TEST\", \"drm\": true, \"episodeCount\": 3, \"genre\": \"Comedy\",\"image\": { \"showImage\" : \"testShowImage\"},\"language\":\"English\", \"nextEpisode\" : {\"channel\" : 1, \"channelLogo\" : \"testChannelLogo\", \"date\" : \"2017-05-25T18:25:43.511Z\", \"html\":\"Test Test\", \"url\" : \"http://test\"}, \"primaryColour\":\"#test\", \"seasons\" : [{\"slug\":\"test/test\"}],\"title\": \"Test Show\", \"slug\": \"show/test\", \"tvChannel\": \"GO!\"}]}")
	req := httptest.NewRequest("GET", "/", reqBody)
	w := httptest.NewRecorder()
	controller.MainHandler(w, req)

	resp := w.Result()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//Empty request (in JSON format) should be okay. 
	if resp.StatusCode != 200 {
		 t.Error("Full Request: Expected status code of 200 got:",resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		  t.Error("Full Request: Expected Content-Type of 'application/json' got:",resp.Header.Get("Content-Type"))
	}
	var resStruct view.Response
	
	err := json.NewDecoder(resp.Body).Decode(&resStruct); 
	if err != nil {
		 t.Error("Full Request: returned error", err)
	}
	
	if len(resStruct.Shows) != 1 {
		t.Error("Full Request:, test returned result. Expect 1 shows returned got", len(resStruct.Shows))
	}
	
	for _, show := range resStruct.Shows {
		if show.Title == "" {
			 t.Error("Full Request: title is not null")
		}
		if show.Slug == "" {
			 t.Error("Full Request: slug is not null")
		}
		if show.Image == "" {
			 t.Error("Full Request: Image is not null")
		}
	} 
	
}
