package controler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"main/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ChangeItem(c echo.Context) (err error) {
	item := new(model.Item)
	id := c.Param("id")
	err = c.Bind(item)
	if err != nil {
		return c.String(http.StatusBadRequest, "khong the bind")
	}
	err = c.Validate(item)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalidate")
	}

	client := new(http.Client)
	json_data, err := json.Marshal(item)
	req, err := http.NewRequest("PUT", "https://623aca0c3c66548cc7a225b6.mockapi.io/todoApp/"+id,
		bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if err == nil {
		respItem := model.Item{}
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal([]byte(responseBody), &respItem)
			return c.JSON(http.StatusBadRequest, respItem)
		}

	}
	// còn vấn đề parse to int Id (Id thuần là str)
	return c.String(http.StatusBadGateway, "co gi do khong on")
}
