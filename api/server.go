package api

import "log"

func NewServer(cfg *Config) (*API, error) {
	a, err := NewSettings(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func (a *API) Run() {
	r := a.NewRouter()
	r.Run(":8080")
}
