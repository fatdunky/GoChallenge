package tests

import (
    "testing"
)

/*
* Unfortunely i ran out of time and was not able to complete the unit testing models.
* The logging function caused a panic for the unit test (did not do this in real curl tests)
*
* If you know the cause i would be happy for help: fatdunky@gmail.com
*
* --- FAIL: TestProcessShows (0.00s)
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
*
*
*/

func TestXYZ(t *testing.T) {

}

