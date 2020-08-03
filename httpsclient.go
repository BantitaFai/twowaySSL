package main

import (
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// load client cert
	cert, err := tls.LoadX509KeyPair("kbank.pentest.2.crt", "kbank.pentest.2.key")
	if err != nil {
		log.Fatal(err)
	}

	// load CA cert
	caCert, err := ioutil.ReadFile("entrust_g2_ca.cer")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// https client tls config
	// InsecureSkipVerify true means not validate server certificate (so no need to set RootCAs)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		//RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	// https client request
	url := "https://uat.openapi-nonprod.kasikornbank.com/exercise/ssl"
	method := "POST"
  
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest(method, url, nil)
  
	if err != nil {
	  fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 5pCDoVL43fSkPbTWmHEg1RGBAUea")
	req.Header.Add("x-test-mode", "true")
  
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
  
	fmt.Println(string(body))
}
