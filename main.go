package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func getPublicIP() (string, error) {
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

func main() {
	domain := flag.String("domain", "", "domain to update")
	token := flag.String("token", "", "token for authentication")
	ip := flag.String("ip", "", "IP optional")

	flag.Parse()

	if *domain == "" {
		fmt.Println("Invalid domain: field cannot be empty")
		return
	}
	fmt.Println("Domain: " + *domain)

	if *token == "" {
		fmt.Println("Invalid token: field cannot be empty")
		return
	}

	var publicIP string
	if *ip != "" {
		publicIP = *ip
	} else {
		var err error
		publicIP, err = getPublicIP()
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	}
	fmt.Println("IP: " + publicIP)

	// Check if the IP is already associated with the domain
	associatedIP, _ := net.LookupIP(*domain + ".duckdns.org")
	if associatedIP[0].String() == publicIP {
		fmt.Println("IP is already associated with the domain")
	} else {

		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		response, err := httpClient.Get("https://www.duckdns.org/update?domains=" + *domain + ".duckdns.org&token=" + *token + "&ip=" + publicIP)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
		defer response.Body.Close()

		fmt.Println("Response: " + response.Status)
	}
}
