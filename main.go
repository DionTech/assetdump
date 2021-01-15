package main

import (
	"flag"

	"github.com/DionTech/assetdump/pckg/assetdump"
	"github.com/DionTech/stdoutformat"
)

func main() {
	var load bool
	flag.BoolVar(&load, "load", false, "load the stored scan")

	flag.Parse()
	domain := flag.Arg(0)

	dump := assetdump.Dump{
		Domain: domain}

	dump.Init()

	if load {
		stdoutformat.Logo()
		dump.Load()
		return
	}

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanHosts()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanMXNames()

	assetdump.ProcessWaitGroup.Add(1)
	go dump.ScanNameserver()

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

	dump.Save()

}
