package proxy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var values map[string]string

func ReverseProxy(ctx echo.Context, c controller.AppController) error {
	result, err := c.Microservice.GetMicroservice(ctx)
	if err != nil {
		log.Error(err.Error())
	}

	if err := ctx.Bind(&values); !errors.Is(err, nil) {
		log.Error(err.Error())
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

	jsonData, _ := json.Marshal(values)

	return jsonData

}

func sendRequest(client *http.Client, ctx echo.Context, microservice *entities.Microservice, jsonData []byte) error {

	endpoint := microservice.Api + "" + microservice.Path

	method := microservice.Method

	fmt.Println()

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
