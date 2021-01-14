package assetdump

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type ThreatCrowdResult struct {
	Resolutions []struct {
		LastResolved string `json:"last_resolved"`
		IP           string `json:"ip_address"`
	} `json:"resolutions"`
	Subdomains []string `json:"subdomains"`
	Emails     []string `json:"emails"`
	Hashes     []string `json:"hashes"`
}

func (dump *Dump) ScanByThreatCrowd() {
	resp, err := http.Get(
		fmt.Sprintf("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s", dump.Domain),
	)
	checkError(err)
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&dump.ThreatCrowdResult)

	for _, subdomain := range dump.ThreatCrowdResult.Subdomains {
		scanSubdomain(dump, subdomain)
	}

	checkError(err)

	ProcessWaitGroup.Done()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func scanSubdomain(dump *Dump, subdomain string) {
	ips, err := net.LookupHost(subdomain)

	if err != nil {
		dump.Subdomains[subdomain] = " n/a "
		return
	}

	dump.Subdomains[subdomain] = strings.Join(ips, " | ")
}
