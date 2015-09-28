package poego

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Poego struct {
	client *http.Client
}

const (
	BaseUrl = "api.pathofexile.com"
	Scheme  = "http"
)

///Initialize the API with a http.client and base values
func NewApi() *Poego {
	return &Poego{
		client: &http.Client{},
	}
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
