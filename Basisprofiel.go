package kvk

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type BasisprofielConfig struct {
	KvkNummer string
	GeoData   *bool
}

func (service *Service) GetBasisprofiel(config *BasisprofielConfig) (*Basisprofiel, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config must not be nil")
	}

	values := url.Values{}
	if config.GeoData != nil {
		values.Set("geoData", fmt.Sprintf("%v", *config.GeoData))
	}

	basisprofiel := Basisprofiel{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("basisprofielen/%s", config.KvkNummer), &values),
		ResponseModel: &basisprofiel,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &basisprofiel, nil
}

func (service *Service) GetBasisprofielEigenaar(config *BasisprofielConfig) (*Eigenaar, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config must not be nil")
	}

	values := url.Values{}
	if config.GeoData != nil {
		values.Set("geoData", fmt.Sprintf("%v", *config.GeoData))
	}

	eigenaar := Eigenaar{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("basisprofielen/%s/eigenaar", config.KvkNummer), &values),
		ResponseModel: &eigenaar,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &eigenaar, nil
}

func (service *Service) GetBasisprofielHoofdvestiging(config *BasisprofielConfig) (*Vestiging, *errortools.Error) {
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
		Url:           service.url(fmt.Sprintf("basisprofielen/%s/hoofdvestiging", config.KvkNummer), &values),
		ResponseModel: &vestiging,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &vestiging, nil
}

func (service *Service) GetBasisprofielVestigingen(config *BasisprofielConfig) (*VestigingList, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config must not be nil")
	}

	values := url.Values{}
	if config.GeoData != nil {
		values.Set("geoData", fmt.Sprintf("%v", *config.GeoData))
	}

	vestigingList := VestigingList{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("basisprofielen/%s/vestigingen", config.KvkNummer), &values),
		ResponseModel: &vestigingList,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &vestigingList, nil
}
