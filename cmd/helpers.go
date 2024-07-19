package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetPublicIP() (string, error) {
	response, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		fmt.Println("Unable to retrive PublicIP\nError: " + err.Error())
		return "", err
	}
	defer response.Body.Close()

	ipBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to retrive PublicIP\nError: " + err.Error())
		return "", err
	}
	return strings.TrimSpace(string(ipBytes)), nil
}

func MakeAPICall(domain, token, ip string) (*http.Response, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	response, err := httpClient.Get("https://www.duckdns.org/update?domains=" + domain + ".duckdns.org&token=" + token + "&ip=" + ip)
	if err != nil {
		return response, err
	}
	
	defer response.Body.Close()

	return response, nil
}
