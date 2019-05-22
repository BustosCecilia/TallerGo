package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
}

const urlCounty = "https://api.mercadolibre.com/classified_locations/countries/"

func (country *Country) Get() *apierrors.ApiError {
	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "countryID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s", urlCounty, country.ID)
	response, err := http.Get(final)
	defer response.Body.Close()
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &country); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
