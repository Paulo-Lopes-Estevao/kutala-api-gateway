package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxy(remote *url.URL) *httputil.ReverseProxy {

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ModifyResponse = func(response *http.Response) error {

		return nil
	}

	return proxy
}

func getProxyUrl(api string) (*url.URL, error) {

	remote, err := url.Parse(api)

	if err != nil {
		panic(err)
	}

	return remote, nil
}

func ReverseProxy() *httputil.ReverseProxy {

	url1, _ := getProxyUrl("http://127.0.0.1:2000")

	return proxy(url1)

}

/* func NewPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//fmt.Fprintln(w, "This is the people handler.", r.RequestURI)
		url, _ := getProxyUrl("http://127.0.0.1:2000")
		proxy(url)
	})
} */
