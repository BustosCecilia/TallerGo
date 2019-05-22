package myml

import (
	"github.com/gin-gonic/gin"
	myml2 "github.com/mercadolibre/taller-go/src/api/domain/myml"
	"github.com/mercadolibre/taller-go/src/api/services/myml"
	"github.com/mercadolibre/taller-go/src/api/utils/apierrors"
	"net/http"
	"strconv"
	"sync"
)

const (
	paramUserID = "userID"
)

func GetUser(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := myml.GetUserFromAPI(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	canal := make(chan myml2.Struct, 3)

	canal <- myml2.Struct{
		User: *user,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go myml.GetCurrencyFromAPI(user.CountryID, canal, &wg)
	go myml.GetSiteFromAPI(user.SiteID, canal, &wg)
	wg.Wait()
	var datosStruct myml2.Struct

	aux := <-canal
	datosStruct.User = aux.User
	aux = <-canal
	datosStruct.Site = aux.Site
	datosStruct.Country = aux.Country
	aux = <-canal
	if datosStruct.Country.ID == "" {
		datosStruct.Country = aux.Country
	}
	if datosStruct.Site.ID == "" {
		datosStruct.Site= aux.Site
	}
	c.JSON(http.StatusOK, datosStruct)
	return
}
