package kvk

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type VestigingsprofielConfig struct {
	Vestigingsnummer int64
	GeoData          *bool
}

func (service *Service) GetVestigingsprofiel(config *VestigingsprofielConfig) (*Vestiging, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config must not be nil")
	}

	values := url.Values{}
	if config.GeoData != nil {
		values.Set("geoData", fmt.Sprintf("%v", *config.GeoData))
	}

	vestiging := Vestiging{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("vestigingsprofielen/%v", config.Vestigingsnummer), &values),
		ResponseModel: &vestiging,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &vestiging, nil
}
