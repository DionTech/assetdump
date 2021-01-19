package assetdump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/DionTech/stdoutformat"
)

type wrapper = struct {
	Domain         string
	HostLookup     HostLookup
	ForwardARecord []net.IP
	MX             []*net.MX
	TXT            []string
	Nameserver     []*net.NS
	Subdomains     map[string]string
	Certificates   map[string]string
}

func (dump *Dump) Save() {
	data := wrapper{
		Domain:         dump.Domain,
		HostLookup:     dump.HostLookup,
		ForwardARecord: dump.ForwardARecord,
		MX:             dump.MX,
		TXT:            dump.TXT,
		Nameserver:     dump.Nameserver,
		Subdomains:     dump.Subdomains,
		Certificates:   dump.Certificates,
	}

	fileName := "./" + dump.Domain + ".json"

	file, err := os.Open(filepath.FromSlash(fileName))

	if err != nil {
		file, err = os.Create(fileName)

		if err != nil {
			stdoutformat.Fatalf("cannot create file %s", fileName)
		}
	}

	defer file.Close()

	jsonVar, err := json.Marshal(data)
	err = ioutil.WriteFile(fileName, jsonVar, os.ModePerm)

	if err != nil {
		stdoutformat.Fatalf("cannot write file: %s", err)
	}
}

func (dump *Dump) Load(pretty bool) {
	fileName := "./" + dump.Domain + ".json"

	file, err := os.Open(filepath.FromSlash(fileName))

	if err != nil {
		stdoutformat.Printf("cannot open file %s", fileName)
		return
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, dump)

	if pretty {
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, byteValue, "", "\t")

		fmt.Println(string(prettyJSON.Bytes()))
	}

}
