package crtfinder

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/norvd/cpm/internal/crtchecker/data"
	"github.com/norvd/cpm/internal/requester"
)

type CrtFinder struct {
	cr *requester.Requester
}

func New(url string) *CrtFinder {
	return &CrtFinder{
		cr: requester.New(url + "/json"),
	}
}

func (cf *CrtFinder) Find(identity string) ([]data.Entry, error) {
	if identity == "" {
		return nil, errors.New("invalid input value")
	}
	params := map[string]string{"q": identity}

	result := cf.cr.GetWithParams(params)

	var certs []data.Entry
	err := json.Unmarshal(result, &certs)
	if err != nil {
		log.Fatalf("failed to unmarshal response json: %v", err)
	}

	return certs, nil

}
