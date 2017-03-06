package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

var (
	cert      = flag.String("certificate", os.Getenv("SSL_CERT"), "Path to the certificate file")
	key       = flag.String("key", os.Getenv("SSL_KEY"), "Path to the key file")
	proxyPort = flag.String("proxy-port", os.Getenv("PROXY_PORT"), "The port on localhost to proxy requests to.")
)

func main() {
	flag.Parse()

	if len(*cert) == 0 {
		log.Fatal("--certificate or $SSL_CERT is required")
	}
	if len(*key) == 0 {
		log.Fatal("--key or $SSL_KEY is required")
	}
	if len(*proxyPort) == 0 {
		log.Fatal("--proxy-port or $PROXY_PORT is required")
	}
	port, err := strconv.Atoi(*proxyPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	u, err := url.Parse(fmt.Sprintf("http://localhost:%d/", port))
	if err != nil {
		log.Fatal(err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServeTLS(":443", *cert, *key, nil))
}
