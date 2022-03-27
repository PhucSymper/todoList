package controler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"main/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteItem(c echo.Context) error {
	id := c.Param("id")
	Item := model.Item{}
	jsonReq, _ := json.Marshal(Item)
	req, err := http.NewRequest(http.MethodDelete, "https://623aca0c3c66548cc7a225b6.mockapi.io/todoApp/"+id, bytes.NewBuffer(jsonReq))
	if err != nil {
		return c.String(http.StatusBadGateway, "Item khong ton tai")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	respItem := model.Item{}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		json.Unmarshal([]byte(responseBody), &respItem)
		return c.JSON(http.StatusBadRequest, respItem)
	}
	// còn vấn đề parse to int Id (Id thuần là str)
	return c.String(http.StatusBadGateway, "co gi do khong on")
	// còn vấn đề parse to int Id (Id thuần là str)
}
