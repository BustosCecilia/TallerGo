package myml

import (
	"github.com/mercadolibre/taller-go/src/api/domain/myml"
	"github.com/mercadolibre/taller-go/src/api/utils/apierrors"
	"sync"
)

func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {
	user := &myml.User{
		ID: userID,
	}
	if apiErr := user.Get(); apiErr != nil {
		return nil, apiErr
	}
	return user, nil
}

//GetSiteFromAPI
func GetSiteFromAPI(siteID string, ch chan myml.Struct, wb *sync.WaitGroup) {
	defer wb.Done()
	site := &myml.Site{
		ID: siteID,
	}
	if apiErr := site.Get(); apiErr != nil {
	}
	ch <- myml.Struct{
		Site: *site,
	}

}

func GetCurrencyFromAPI(currencyID string, ch chan myml.Struct, wb *sync.WaitGroup) {
	defer wb.Done()
	currency := &myml.Country{
		ID: currencyID,
	}
	if apiErr := currency.Get(); apiErr != nil {
	}

	ch <- myml.Struct{
		Country: *currency,
	}

}
