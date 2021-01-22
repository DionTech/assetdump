package assetdump

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sort"
	"strings"

	"github.com/DionTech/stdoutformat"
)

/**
a result would be something like this:
"issuer_ca_id": 16418,
"issuer_name": "C=US, O=Let's Encrypt, CN=Let's Encrypt Authority X3",
"common_name": "daylite.publicare.de",
"name_value": "daylite.publicare.de",
"id": 3627724442,
"entry_timestamp": "2020-11-10T05:32:34.189",
"not_before": "2020-11-10T04:32:33",
"not_after": "2021-02-08T04:32:33",
"serial_number": "035844984b81b368ea1072c8f35e5a6aa0b5"
**/
type CrtShResult struct {
	Name string `json:"name_value"`
}

func (dump *Dump) ScanByCrtSh() {
	var results []CrtShResult
	var uniqueResults = make(map[string]CrtShResult, 0)

	resp, err := http.Get(
		fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", dump.Domain),
	)
	if err != nil {
		stdoutformat.Error(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		stdoutformat.Error(err)
		return
	}

	//need to be make unique results here!!!
	for _, res := range results {
		_, present := uniqueResults[res.Name]

		if !present {
			uniqueResults[res.Name] = res
			dump.CrtShList = append(dump.CrtShList, res.Name)
		}
	}

	sort.Strings(dump.CrtShList)

	for _, cert := range dump.CrtShList {
		scanCert(dump, cert)
	}

	ProcessWaitGroup.Done()
	dump.Bar.Increment()
}

func scanCert(dump *Dump, cert string) {
	ips, err := net.LookupHost(cert)

	if err != nil {
		dump.Certificates[cert] = " n/a "
		return
	}

	dump.Certificates[cert] = strings.Join(ips, " | ")
}
