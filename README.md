# Go Go Go - A quick intro to concurrency in Go

Go intro talk presented at Haifa:Dev meet up.

##### Requirements
- Install Go

## Structure
#### main.go
##### Main Func
A small Go http server running on port :8090.
##### APIs for getting car insurance qoutes
- GET /quote - , handles flow synchronosly
- GET /quote/v2 - handles flow asynchronosly, using gorountines
- GET /quote/v3 - handles flow asynchronosly, using gorountines + wait groups

##### Functions for mocking db/api calls
Each function has a synthetic sleep time of 300ms.
- getUser(“uid”) - fetch user by id
- getUserSub(“uid”) - fetch user subscription
- generateQuote(“uid”) - generate car insurance quoute

##### main_test.go
- Compare goroutine strategies
- Load test
