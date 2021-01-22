package assetdump

import (
	"net"
	"sync"

	"github.com/DionTech/stdoutformat"
	"github.com/cheggaaa/pb/v3"
)

type HostLookup struct {
	IPs map[string][]string
}

type Dump struct {
	Bar               *pb.ProgressBar
	Domain            string
	HostLookup        HostLookup
	CrtShList         []string
	ThreatCrowdResult ThreatCrowdResult
	ForwardARecord    []net.IP
	MX                []*net.MX
	Nameserver        []*net.NS
	TXT               []string
	Subdomains        map[string]string
	Certificates      map[string]string
}

var ProcessWaitGroup sync.WaitGroup

func (dump *Dump) Init() {
	dump.HostLookup.IPs = make(map[string][]string, 0)
	dump.CrtShList = make([]string, 0)
	dump.Subdomains = make(map[string]string, 0)
	dump.Certificates = make(map[string]string, 0)
}

func (dump *Dump) ScanHosts() {
	ips, err := net.LookupHost(dump.Domain)

	if err != nil {
		stdoutformat.Error(err)
		ProcessWaitGroup.Done()
		dump.Bar.Increment()
		return
	}

	for _, ip := range ips {
		addresses, _ := net.LookupAddr(string(ip))
		dump.HostLookup.IPs[ip] = addresses
	}

	ProcessWaitGroup.Done()
	dump.Bar.Increment()
}

func (dump *Dump) ScanMXNames() {
	dump.MX, _ = net.LookupMX(dump.Domain)
	ProcessWaitGroup.Done()
	dump.Bar.Increment()
}

func (dump *Dump) ScanTXTRecords() {
	dump.TXT, _ = net.LookupTXT(dump.Domain)
	ProcessWaitGroup.Done()
	dump.Bar.Increment()
}

func (dump *Dump) ScanForwardARecord() {
	dump.ForwardARecord, _ = net.LookupIP(dump.Domain)
	ProcessWaitGroup.Done()
	dump.Bar.Increment()
}

func (dump *Dump) ScanNameserver() {
	dump.Nameserver, _ = net.LookupNS(dump.Domain)
	ProcessWaitGroup.Done()
	dump.Bar.Increment()
}
