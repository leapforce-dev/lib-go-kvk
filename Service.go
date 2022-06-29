package kvk

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName     string = "KvK"
	apiPath     string = "https://api.kvk.nl/api/v1"
	apiPathTest string = "https://developers.kvk.nl/test/api/v1"
)

type Service struct {
	apiKey        string
	isTest        bool
	httpService   *go_http.Service
	errorResponse *ErrorResponse
}

type ServiceConfig struct {
	ApiKey string
	IsTest bool
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ApiKey == "" {
		return nil, errortools.ErrorMessage("Service ApiKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiKey:      serviceConfig.ApiKey,
		isTest:      serviceConfig.IsTest,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add error model
	service.errorResponse = &ErrorResponse{}
	(*requestConfig).ErrorModel = &service.errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if len(service.errorResponse.Fout) > 0 {
		e.SetMessage(service.errorResponse.Fout[0].Omschrijving)
	}

	return request, response, e
}

func (service *Service) url(path string, values *url.Values) string {
	values_ := url.Values{}
	if values != nil {
		values_ = *values
	}
	values_.Set("user_key", service.apiKey)

	apiPath_ := apiPath
	if service.isTest {
		apiPath_ = apiPathTest
	}

	return fmt.Sprintf("%s/%s?%s", apiPath_, path, values_.Encode())
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}

func (service *Service) ErrorResponse() *ErrorResponse {
	return service.errorResponse
}
