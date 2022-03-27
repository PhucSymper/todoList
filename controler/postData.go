package controler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"main/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostNewItem(c echo.Context) (err error) {
	item := new(model.Item)
	err = c.Bind(item)
	if err != nil {
		return c.String(http.StatusBadRequest, "khong the bind")
	}
	err = c.Validate(item)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalidate")
	}
	json_data, err := json.Marshal(item)
	resp, err := http.Post("https://623aca0c3c66548cc7a225b6.mockapi.io/todoApp", "application/json",
		bytes.NewBuffer(json_data))
	if err == nil {
		respITem := new(model.Item)
		responseBody, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(responseBody), &respITem)
		return c.JSON(http.StatusOK, respITem)
	}
	// bind to object for response json
	return c.String(http.StatusBadGateway, "co cai gi do sai sai")
}
