This project will read HTTP JSON requests and apply the following business logic.

From the list of shows in the request payload, return the ones with DRM enabled (drm: true)
and at least one episode (episodeCount > 0)

The returned JSON should have a response key with an array of shows.
Each element should have the following fields from the request:

image - corresponding to image/showImage from the request payload
slug
title

Main URL:https://1-dot-stan-challenge.appspot.com/

Repo: https://github.com/fatdunky/GoChallenge

Project Structure:

This app has been set up using the standard MVC structure.

.project 					-- The .project file is an eclipse config file. It can be ignored
app.yaml 					-- I have hosted it using google cloud. Which uses the deployment descriptor app.yaml. It will kick off the app using
							 			 the init() func in StanChallenge.go
README.md
LICENSE						-- Included license file. Logging functionality is a modified version of  https://github.com/goinggo/tracelog
src/							-- The code is under here
StanChallenge.go	-- Entry point for app. It will map requests to the main controller. (It will go through a middleware that logs requests)
test/							-- Contains the unit testing. Unfortunately my logging seemed to cause a invalid memory issue with the testing.

/src/..
controller/				-- The controller package that handles requests
model/						-- The model packge, this contains the business logic
utilities/				-- Contains the my logging package
view/							-- Contains the JSON parse and config

Running:

As this was built for the google cloud. You will need to change the init() func in StanChallenge.go to main() to run this from the command line.
After that it should run fine with go run StanChallenge.go

Licensing:

The logging functions are a modified version of the package found at https://github.com/goinggo/tracelog, which contains "Simplified BSD License". This has been included.

Testing:
 Unfortunely i ran out of time and was not able to complete the unit testing models.
 The logging function caused a panic for the unit test (did not do this in real curl tests)

 If you know the cause i would be happy for help: fatdunky@gmail.com

 --- FAIL: TestProcessShows (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
        panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0x5c9fcb]

goroutine 18 [running]:
testing.tRunner.func1(0xc042042dd0)
        C:/Go/src/testing/testing.go:622 +0x2a4
panic(0x667b20, 0x7e5600)
        C:/Go/src/runtime/panic.go:489 +0x2dd
log.(*Logger).Output(0x0, 0x2, 0xc0420d7220, 0x1f, 0x0, 0x0)
        C:/Go/src/log/log.go:149 +0x5b
utilities/logging.Started(0x6b2ca3, 0x5, 0x6b4b93, 0xc)
        c:/Users/mcrick/workspace/StanChallenge/src/utilities/logging/loggingFuncs.go:24 +0x141
model.ProcessShows(0xc04201fe90, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
        c:/Users/mcrick/workspace/StanChallenge/src/model/logic.go:28 +0x117
_/c_/Users/mcrick/workspace/StanChallenge/test.TestProcessShows(0xc042042dd0)
        c:/Users/mcrick/workspace/StanChallenge/test/logic_test.go:16 +0x14f
testing.tRunner(0xc042042dd0, 0x6c3638)
        C:/Go/src/testing/testing.go:657 +0x9d
created by testing.(*T).Run
        C:/Go/src/testing/testing.go:697 +0x2d1
FAIL    _/c_/Users/mcrick/workspace/StanChallenge/test  0.382s



