# Poego
Wrapper for Path of Exile Web API written in Go

[![Build Status](https://travis-ci.org/JoshuaThompson/poego.svg?branch=master)](https://travis-ci.org/JoshuaThompson/poego)

###Supported API Calls
1. GetLeagues
2. GetLeague
3. GetLeagueRules
4. GetLeagueRule
5. GetLadder
6. GetPvpMatches

###Additional Functionality
1. GetEntireLadder
2. GetEntireLeagueLadder

These two calls make 75 (unfotunately, due to GGG's API implementation) requests over the course of 25~ seconds to the api in order to get the 15000 maximum ladder entries, with the league function returning the league information as well.  Be careful with how often you call either of these functions due to rate limiting.  


More information about these calls can be found on the [Path of Exile API Website](https://www.pathofexile.com/developer/docs/api) and in the [godoc](http://godoc.org/github.com/JoshuaThompson/poego).

###Installation
```
go get github.com/joshuathompson/poego
```

###Tests
```
go test
```

###Usage 
```go
package main

import (
	"fmt"
	"log"
	"github.com/joshuathompson/poego"
)

func main() {

	p := poego.NewPoeApi()

	//optional arguemnts can be passed via url.Values or map[string][]string
	//this is clearly demonstrated in a few of the tests
	leagues, err := p.GetLeagues(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", leagues)
}

```

###License
MIT