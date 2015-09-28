# Poego
Wrapper for Path of Exile Web API written in Go

###Supported API Calls
1. GetLeagues
2. GetLeague
3. GetLeagueRules
4. GetLeagueRule
5. GetLadder
6. GetPvpMatches


More information about these calls can be found on the [Path of Exile API Website](https://www.pathofexile.com/developer/docs/api) and in the [godoc](http://godoc.org/github.com/JoshuaThompson/poego).

###Installation
```
go get github.com/joshuathompson/poego
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

	p := poego.NewApi()

	//optional arguemnts can be passed via url.Values or map[string][]string
	leagues, err := p.GetLeagues()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", leagues)
}

```

###License
MIT