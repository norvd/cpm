package crtchecker

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/norvd/cpm/internal/crtchecker/data"
	"github.com/norvd/cpm/internal/requester"
)

type CrtChecker struct {
	r *requester.Requester
}

func New(url string) *CrtChecker {
	return &CrtChecker{
		r: requester.New(url + "/json"),
	}
}

func (cc *CrtChecker) Find(identity string) ([]data.Entry, error) {
	if identity == "" {
		return nil, errors.New("invalid input value")
	}
	params := map[string]string{"q": identity}

	result := cc.r.GetWithParams(params)

	var certs []data.Entry
	err := json.Unmarshal(result, &certs)
	if err != nil {
		log.Fatalf("failed to unmarshal response json: %v", err)
	}

	return certs, nil

}
