package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/labstack/echo/v4"
)

func ReverseProxy(ctx echo.Context, c controller.AppController) error {
	var values = map[string]string{}
	result, err := c.Microservice.GetMicroservice(ctx)
	if err != nil {
		log.Fatal("Error get microservice", err.Error())
	}

	if err := ctx.Bind(&values); err != nil {
		log.Fatal("Error Bind", err.Error())
	}

	clientRequest := httpClient()

	body := Body(values)

	reponse := sendRequest(clientRequest, ctx, result, body)

	return reponse
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func Body(values map[string]string) []byte {

	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal("Error body", err.Error())
	}

	return jsonData

}

func sendRequest(client *http.Client, ctx echo.Context, microservice *entities.Microservice, jsonData []byte) error {

	endpoint := microservice.Api + "" + microservice.Endpoint

	method := microservice.Method

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return ctx.JSON(400, fmt.Sprintf("Error Occurred. %v", err))
	}

	response, err := client.Do(req)
	if err != nil {
		return ctx.JSON(400, fmt.Sprintf("Error sending request to API endpoint. %v", err))
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ctx.JSON(400, fmt.Sprintf("Couldn't parse response body. %+v", err))
	}

	return ctx.JSONBlob(200, body)
}
