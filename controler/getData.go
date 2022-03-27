package controler

import (
	"encoding/json"
	"io/ioutil"
	"main/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllItem(c echo.Context) error {

	response, err := http.Get("https://623aca0c3c66548cc7a225b6.mockapi.io/todoApp")
	if err != nil {
		return c.String(http.StatusBadRequest, "khong the lay du lieu")
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "oops, something wrong")
	}
	data1 := []model.Item{}
	json.Unmarshal([]byte(responseBody), &data1)
	return c.JSON(http.StatusOK, data1)
}

func GetDetailItem(c echo.Context) error {
	id := c.Param("id")
	response, err := http.Get("https://623aca0c3c66548cc7a225b6.mockapi.io/todoApp/" + id)
	if err != nil {
		panic(err)
	}
	if err == nil {
		respItem := model.Item{}
		responseBody, err := ioutil.ReadAll(response.Body)
		if err == nil {
			json.Unmarshal([]byte(responseBody), &respItem)
			return c.JSON(http.StatusBadRequest, respItem)
		}

	}
	// còn vấn đề parse to int Id (Id thuần là str)
	return c.String(http.StatusBadGateway, "co gi do khong on")
}

