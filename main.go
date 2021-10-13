package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/injection"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/handler"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/middleware"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
	"github.com/labstack/echo"
)

func Target(path string) (string, error) {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	return parts[0], nil
}

func getProxyUrl(proxyConditionRaw string) *url.URL {

	proxyCondition := strings.ToUpper(proxyConditionRaw)

	//a_condtion_url := ("SEARCHURL")
	//default_condtion_url := ("DEFAULT_CONDITION_URL")

	remote, err := url.Parse(proxyConditionRaw)
	if err != nil {
		panic(err)
	}

	if proxyCondition == "A" {
		return remote
	}

	return remote
}

func ReverseProxy(echoCtx echo.Context) {
	url := getProxyUrl("http://172.16.16.37:8084")
	//path, _ := Target(ginCtx.Request.URL.Path)

	proxy(url).ServeHTTP(echoCtx.Response().Writer, echoCtx.Request())

}

func proxy(remote *url.URL) *httputil.ReverseProxy {

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ModifyResponse = func(response *http.Response) error {
		//defer response.Body.Close()
		//u, _ := ioutil.ReadAll(response.Body)
		return nil
	}

	return proxy
}

var private_server = echo.New()
var public_server = echo.New()

func main() {
	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	handler.RouteServiceHandler(public_server, i.NewAppController())

	middleware.RegistryMiddleware(private_server, i.NewAppController())
	handler.RouteMicroserviceHandler(private_server, i.NewAppController())

	go func() {
		fmt.Println("Running Server A")
		if err := private_server.Start(":8080"); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fmt.Println("Running Server B")
	if err := public_server.Start(":8081"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
