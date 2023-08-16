package api

import (
	"log"

	"github.com/apex/gateway"
)

func NewServer(cfg *Config) (*API, error) {
	a, err := NewSettings(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func (a *API) Run() {
	r := a.NewRouter()
	if err := gateway.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
