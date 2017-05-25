package tests

import (
    "testing"
    "view"
    "model"
    "fmt"
)

func TestProcessShowsHappyPath(t *testing.T) {
	//Test 1: Happy Path
	img1 := view.Image{ShowImage:"imageTest1"}
	payload1 := view.Payload{Title:"Test1", Slug:"SlugTest1", ImageObj:img1, Drm:true, EpisodeCount:1}
	payloads := []view.Payload{payload1}
	test1 := view.Request{Payloads:payloads}

	retVal,err := model.ProcessShows(nil, test1)
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("Happy Path Test 1, returned error", err)
	}
	for index, show := range retVal.Shows {
		img := show.Image
		slug := show.Slug
		title := show.Title
		
		if (img != "imageTest1") {
			 t.Error("Happy Path Test 1: Expected imageTest1 got", img)
		}
		if (slug != "SlugTest1") {
			 t.Error("Happy Path Test 1: Expected SlugTest1 got", slug)
		}
		if (title != "Test1") {
			 t.Error("Happy Path Test 1: Expected Test1 got", title)
		}
		if (index > 1) {
			 t.Error("Happy Path Test 1: More then one show returned:", index)
		}
	}
}

func TestProcessShowsEpCountZero(t *testing.T) {
	img1 := view.Image{ShowImage:"imageTest1"}
	payload1 := view.Payload{Title:"Test1", Slug:"SlugTest1", ImageObj:img1, Drm:true, EpisodeCount:0}
	payloads := []view.Payload{payload1}
	test1 := view.Request{Payloads:payloads}

	retVal,err := model.ProcessShows(nil, test1)
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("Episode Count Zero: Test 1, returned error", err)
	}
	
	if len(retVal.Shows) != 0 {
		t.Error("Episode Count Zero: Test 1, test returned result. Expect 0 shows returned got %s", len(retVal.Shows))
	}
}

func TestProcessShowsEpCountOne(t *testing.T) {
	img1 := view.Image{ShowImage:"imageTest1"}
	payload1 := view.Payload{Title:"Test1", Slug:"SlugTest1", ImageObj:img1, Drm:true, EpisodeCount:1}
	payloads := []view.Payload{payload1}
	test1 := view.Request{Payloads:payloads}

	retVal,err := model.ProcessShows(nil, test1)
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("Episode Count One: Test 1, returned error", err)
	}
	
	if len(retVal.Shows) != 1 {
		t.Error("Episode Count One: Test 1, test returned result. Expect 1 shows returned got", len(retVal.Shows))
	}
	
	for _, show := range retVal.Shows {
		img := show.Image
		slug := show.Slug
		title := show.Title
		
		if (img != "imageTest1") {
			 t.Error("Happy Path Test 1: Expected imageTest1 got", img)
		}
		if (slug != "SlugTest1") {
			 t.Error("Happy Path Test 1: Expected SlugTest1 got", slug)
		}
		if (title != "Test1") {
			 t.Error("Happy Path Test 1: Expected Test1 got", title)
		}
	}
}

func TestProcessShowsDRMFalse(t *testing.T) {
	img1 := view.Image{ShowImage:"imageTest1"}
	payload1 := view.Payload{Title:"Test1", Slug:"SlugTest1", ImageObj:img1, Drm:false, EpisodeCount:1}
	payloads := []view.Payload{payload1}
	test1 := view.Request{Payloads:payloads}

	retVal,err := model.ProcessShows(nil, test1)
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("DRM False: Test 1, returned error", err)
	}
	
	if len(retVal.Shows) != 0 {
		t.Error("DRM False: Test 1, test returned result. Expect 0 shows returned got", len(retVal.Shows))
	}
	
}

func TestProcessShowsDRMTrue(t *testing.T) {
	img1 := view.Image{ShowImage:"imageTest1"}
	payload1 := view.Payload{Title:"Test1", Slug:"SlugTest1", ImageObj:img1, Drm:true, EpisodeCount:1}
	payloads := []view.Payload{payload1}
	test1 := view.Request{Payloads:payloads}

	retVal,err := model.ProcessShows(nil, test1)
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("DRM true: Test 1, returned error", err)
	}
	
	if len(retVal.Shows) != 1 {
		t.Error("DRM true: Test 1, test returned result. Expect 1 shows returned got", len(retVal.Shows))
	}
	
	for _, show := range retVal.Shows {
		img := show.Image
		slug := show.Slug
		title := show.Title
		
		if (img != "imageTest1") {
			 t.Error("Happy Path Test 1: Expected imageTest1 got", img)
		}
		if (slug != "SlugTest1") {
			 t.Error("Happy Path Test 1: Expected SlugTest1 got", slug)
		}
		if (title != "Test1") {
			 t.Error("Happy Path Test 1: Expected Test1 got", title)
		}
	} 
}
func TestProcessShowsMultiHappyPath(t *testing.T) {
	var payloads []view.Payload
	for i := 1; i < 21; i ++ {
		img1 := view.Image{ShowImage:(fmt.Sprintf("imageTest%d", i))}
		var payload1 view.Payload
		if (i < 11) {			
			if ( i  % 2 == 0 ) {
				payload1 = view.Payload{Title:(fmt.Sprintf("Test%d", i)), Slug:(fmt.Sprintf("SlugTest%d", i)), ImageObj:img1, Drm:true, EpisodeCount:1}
			} else {
				payload1 = view.Payload{Title:(fmt.Sprintf("Test%d", i)), Slug:(fmt.Sprintf("SlugTest%d", i)), ImageObj:img1, Drm:false, EpisodeCount:1}
			}
		} else {
			if ( i  % 2 == 0 ) {
				payload1 = view.Payload{Title:(fmt.Sprintf("Test%d", i)), Slug:(fmt.Sprintf("SlugTest%d", i)), ImageObj:img1, Drm:true, EpisodeCount:0}
			} else {
				payload1 = view.Payload{Title:(fmt.Sprintf("Test%d", i)), Slug:(fmt.Sprintf("SlugTest%d", i)), ImageObj:img1, Drm:false, EpisodeCount:0}
			}
		}
		payloads = append(payloads, payload1)
	}
	test1 := view.Request{Payloads:payloads}
	retVal,err := model.ProcessShows(nil, test1)
	
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("Happy Path: Multi test, returned error", err)
	}
	if len(retVal.Shows) != 5 {
		t.Error("Happy Path: Multi test, test returned result. Expect 5 shows returned got", len(retVal.Shows))
	}
	for index, show := range retVal.Shows {
		img := show.Image
		slug := show.Slug
		title := show.Title
		correctShow := (index + 1) * 2
		
		if (img != fmt.Sprintf("imageTest%d", correctShow)) {
			 message := fmt.Sprintf("Happy Path: Multi test %d: Expected: %s got %s",index,fmt.Sprintf("imageTest%d", correctShow),img)
			 t.Error(message)
		}
		if (slug != fmt.Sprintf("SlugTest%d", correctShow)) {
			 message := fmt.Sprintf("Happy Path: Multi test %d: Expected: %s got %s",index,fmt.Sprintf("SlugTest%d", correctShow),img)
			 t.Error(message)
		}
		if (title != fmt.Sprintf("Test%d", correctShow)) {
			 message := fmt.Sprintf("Happy Path: Multi test %d: Expected: %s got %s",index,fmt.Sprintf("Test%d", correctShow),img)
			 t.Error(message)
		}
		
	}
}


func TestProcessShowsEmpty(t *testing.T) {
	var payloads []view.Payload
	test1 := view.Request{Payloads:payloads}
	retVal,err := model.ProcessShows(nil, test1)
	
	
	//We expect 1 show returned
	if err != nil {
		 t.Error("Empty Test, returned error", err)
	}

	if len(retVal.Shows) != 0 {
		t.Error("Empty Test: Test 1, test returned result. Expect 0 shows returned got", len(retVal.Shows))
	}
}

	
