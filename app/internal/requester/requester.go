package requester

import (
	"io"
	"log"
	"net/http"
)

type Requester struct {
	Url    string
	Client *http.Client
}

func New(url string) *Requester {
	return &Requester{
		Url:    url,
		Client: &http.Client{},
	}

}

func (r *Requester) GetWithParams(params map[string]string) []byte {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		log.Fatalf("error preparing request: %v", err)
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := r.Client.Do(req)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK HTTP status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return body
}
