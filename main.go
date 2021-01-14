package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/DionTech/assetdump/pckg/assetdump"
	"github.com/DionTech/stdoutformat"
)

func main() {

	flag.Parse()
	domain := flag.Arg(0)

	dump := assetdump.Dump{
		Domain: domain}

	dump.Init()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanHosts()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanMXNames()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanTXTRecords()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanForwardARecord()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanByCrtSh()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanByThreatCrowd()

	assetdump.ProcessWaitGroup.Wait()

	stdoutformat.Logo()

	//@TODO: make a decision at cmd between output or save as json
	//dump.Output()

	type wrapper = struct {
		Domain         string
		HostLookup     assetdump.HostLookup
		ForwardARecord []net.IP
		MX             []*net.MX
		TXT            []string
		Subdomains     map[string]string
		Certificates   map[string]string
	}

	data := wrapper{
		Domain:         dump.Domain,
		HostLookup:     dump.HostLookup,
		ForwardARecord: dump.ForwardARecord,
		MX:             dump.MX,
		TXT:            dump.TXT,
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
