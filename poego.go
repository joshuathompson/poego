/*
Package poego provides a wrapper of the Path of Exile API in Go
*/
package poego

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Poego struct {
	client *http.Client
}

//This is the error response structure from the Path of Exile API
//Potential values are defined as follows:
//Code	Message
//1		Resource not found
//2		Invalid query
//3		Rate limit exceeded
//4		Internal error
//5		Unexpected content type
//6		Forbidden
type ApiError struct {
	Error struct {
		Code    int8   `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

const (
	BaseUrl                    = "api.pathofexile.com"
	Scheme                     = "http"
	TotalPossibleLadderEntries = 15000
	MaxLadderSegmentSize       = 200
)

//Initialize the API by creating a http.client
func NewPoeApi() *Poego {
	return &Poego{
		client: &http.Client{},
	}
}

func (p *Poego) buildRequestsForEntireLadder(method, endpoint string) []*http.Request {

	numUrls := TotalPossibleLadderEntries / MaxLadderSegmentSize
	var requests []*http.Request

	for i := 0; i < numUrls; i++ {

		v := url.Values{}
		v.Add("limit", "200")
		v.Add("offset", strconv.Itoa(i*200))

		requests = append(requests, p.buildRequest(method, endpoint, v))
	}

	return requests
}

func (p *Poego) buildRequest(method, endpoint string, v url.Values) *http.Request {

	if v == nil {
		v = url.Values{}
	}

	u := &url.URL{
		Scheme:   Scheme,
		Path:     BaseUrl + endpoint,
		RawQuery: v.Encode(),
	}

	r, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return r
}

func (p *Poego) makeRequest(r *http.Request, d interface{}) error {

	res, err := p.client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return handleApiError(res)
	}

	err = json.NewDecoder(res.Body).Decode(d)

	return err
}

func handleApiError(res *http.Response) error {

	var apiError ApiError
	err := json.NewDecoder(res.Body).Decode(&apiError)

	if err != nil {
		return fmt.Errorf("Status Code: %d\nMessage: Couldn't decode API error", res.StatusCode)
	}

	return fmt.Errorf("Status Code: %d\nPoE Error Code: %d\nMessage: %s", res.StatusCode, apiError.Error.Code, apiError.Error.Message)
}
