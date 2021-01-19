package main

import (
	"flag"

	"github.com/DionTech/assetdump/pckg/assetdump"
	"github.com/DionTech/stdoutformat"
)

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, "show help")

	var load bool
	flag.BoolVar(&load, "load", false, "load the stored scan")

	var pretty bool
	flag.BoolVar(&pretty, "pretty", false, "show json pretty print of the scan")

	var ips bool
	flag.BoolVar(&ips, "ips", false, "show all fetched ips")

	var hosts bool
	flag.BoolVar(&hosts, "hosts", false, "show all fetched hosts")

	flag.Parse()
	domain := flag.Arg(0)

	if help {
		stdoutformat.Logo()
		flag.PrintDefaults()
		return
	}

	dump := assetdump.Dump{
		Domain: domain}

	dump.Init()

	if load {
		dump.Load(pretty)

		if ips {
			dump.OutputIPs()
		}

		if hosts {
			dump.OutputHosts()
		}

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

	//@TODO: make a decision at cmd between output or save as json
	//dump.Output()

	dump.Save()

}
