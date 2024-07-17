package kvk

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type KvkType string

const (
	KvkTypeHoofdvestiging KvkType = "hoofdvestiging"
	KvkTypeNevenvestiging KvkType = "nevenvestiging"
	KvkTypeRechtspersoon  KvkType = "rechtspersoon"
)

type ZoekenConfig struct {
	KvkNummer          *string
	Rsin               *string
	Vestigingsnummer   *string
	Naam               *string
	Straatnaam         *string
	Plaats             *string
	PostcodeHuisnummer *struct {
		Postcode      string
		Huisnummer    *int64
		Huisletter    *string
		Postbusnummer *int64
	}
	Type                           *KvkType
	InclusiefInactieveRegistraties *bool
	Pagina                         *int
	ResultatenPerPagina            *int
}

// GetAccounts returns all accounts
func (service *Service) Zoeken(config *ZoekenConfig) (*[]ResultaatItem, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("config must not be nil")
	}

	values := url.Values{}
	if config.KvkNummer != nil {
		values.Set("kvkNummer", *config.KvkNummer)
	}
	if config.Rsin != nil {
		values.Set("rsin", *config.Rsin)
	}
	if config.Vestigingsnummer != nil {
		values.Set("vestigingsnummer", *config.Vestigingsnummer)
	}
	if config.Naam != nil {
		values.Set("naam", *config.Naam)
	}
	if config.Straatnaam != nil {
		values.Set("straatnaam", *config.Straatnaam)
	}
	if config.Plaats != nil {
		values.Set("plaats", *config.Plaats)
	}
	if config.PostcodeHuisnummer != nil {
		values.Set("postcode", config.PostcodeHuisnummer.Postcode)
		if config.PostcodeHuisnummer.Huisnummer != nil {
			values.Set("huisnummer", fmt.Sprintf("%v", *config.PostcodeHuisnummer.Huisnummer))
		}
		if config.PostcodeHuisnummer.Huisletter != nil {
			values.Set("huisletter", *config.PostcodeHuisnummer.Huisletter)
		}
		if config.PostcodeHuisnummer.Postbusnummer != nil {
			values.Set("postbusnummer", fmt.Sprintf("%v", *config.PostcodeHuisnummer.Postbusnummer))
		}
	}
	if config.Type != nil {
		values.Set("type", string(*config.Type))
	}
	if config.InclusiefInactieveRegistraties != nil {
		values.Set("inclusiefInactieveRegistraties", fmt.Sprintf("%v", *config.InclusiefInactieveRegistraties))
	}
	if config.ResultatenPerPagina != nil {
		values.Set("resultatenPerPagina", fmt.Sprintf("%v", *config.ResultatenPerPagina))
	}

	page := 1
	if config.Pagina != nil {
		page = *config.Pagina
	}

	results := []ResultaatItem{}

	for {
		values.Set("pagina", fmt.Sprintf("%v", page))

		resultaat := Resultaat{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.urlV2("zoeken", &values),
			ResponseModel: &resultaat,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		results = append(results, resultaat.Resultaten...)

		if resultaat.Volgende == "" {
			break
		}

		if config.Pagina != nil {
			// explicit page was requested
			break
		}

		page += 1
	}

	return &results, nil
}
