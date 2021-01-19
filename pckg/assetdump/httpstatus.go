package assetdump

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/DionTech/stdoutformat"
)

var UserAgent = "Mozilla/5.0 (compatible; scrape/1.0; +github.com/DionTech/scrape)"

var transport = &http.Transport{
	TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	DisableKeepAlives: true,
	DialContext: (&net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: time.Second,
		DualStack: true,
	}).DialContext,
}

var httpClient = &http.Client{
	Transport: transport,
}

func HTTPStatus(domain string) {
	if strings.HasPrefix(domain, ".") {
		domain = strings.Replace(domain, ".", "", 1)
	}

	if !strings.HasPrefix(domain, "www.") {
		domain = "www." + domain
	}

	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		makeRequest("http://" + domain)
		makeRequest("https://" + domain)

		return
	}

	makeRequest(domain)
}

func makeRequest(domain string) {

	var req *http.Request
	var err error
	req, err = http.NewRequest("GET", domain, nil)

	if err != nil {
		stdoutformat.Error(err)
		fmt.Println("\n\n")
		return
	}
	req.Close = true
	req.Header.Set("UserAgent", UserAgent)

	resp, err := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		stdoutformat.Error(err)
		fmt.Println("\n\n")
		return
	}

	//body, _ := ioutil.ReadAll(resp.Body)

	// extract the response headers
	//hs := make([]string, 0)
	fmt.Println(domain)
	for k, vs := range resp.Header {
		for _, v := range vs {
			//hs = append(hs, fmt.Sprintf("%s: %s", k, v))
			fmt.Printf("%s: %s \n", k, v)
		}
	}

	fmt.Println("\n\n")
}
