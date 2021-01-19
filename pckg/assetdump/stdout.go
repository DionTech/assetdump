package assetdump

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tatsushid/go-prettytable"
)

func (dump *Dump) OutputHosts() {
	hosts := make(map[string]bool, 0)

	for _, hostList := range dump.HostLookup.IPs {
		for _, host := range hostList {
			if _, exists := hosts[host]; exists == false {
				hosts[host] = true
			}
		}
	}

	for subdomain := range dump.Subdomains {
		if _, exists := hosts[subdomain]; exists == false {
			hosts[subdomain] = true
		}
	}

	for crt := range dump.Certificates {
		if _, exists := hosts[crt]; exists == false {
			hosts[crt] = true
		}
	}

	//ordered map here
	keys := make([]string, 0, len(hosts))
	for k := range hosts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, hosts[k])
	}
}

func (dump *Dump) OutputIPs() {
	ips := make(map[string]bool, 0)

	for ip, _ := range dump.HostLookup.IPs {
		if _, exists := ips[ip]; exists == false {
			ips[ip] = true
		}
	}

	for _, ip := range dump.Certificates {
		for _, part := range strings.Split(ip, "|") {
			part = strings.Trim(part, " ")
			if part != "n/a" {
				if _, exists := ips[part]; exists == false {
					ips[part] = true
				}
			}
		}
	}

	for _, ip := range dump.Subdomains {
		for _, part := range strings.Split(ip, "|") {
			part = strings.Trim(part, " ")
			if part != "n/a" {
				if _, exists := ips[part]; exists == false {
					ips[part] = true
				}
			}
		}
	}

	//ordered map here
	keys := make([]string, 0, len(ips))
	for k := range ips {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, ips[k])
	}

}

func (dump *Dump) Output() {
	fmt.Printf("\n\n_____________HOST-LookUp:\n\n")

	PrintHostLookup(dump)

	fmt.Printf("\n\n_____________DNS TXT Records:\n")

	PrintTXTRecords(dump)

	fmt.Printf("\n\n_____________ForwardARecord: \n")

	PrintForwardARecord(dump)

	fmt.Printf("\n\n_____________MX:\n\n")

	PrintMXRecords(dump)

	fmt.Printf("\n\n_____________Certificate Search:\n")

	PrintCrtShList(dump)

	fmt.Printf("\n\n_____________ThreatCrowd:\n\n")

	PrintThreatcrowd(dump)

}

func PrintHostLookup(dump *Dump) {
	tbl, _ := prettytable.NewTable([]prettytable.Column{
		{Header: "IP"},
		{Header: "Hosts"},
	}...)

	for ip, hosts := range dump.HostLookup.IPs {
		hostsOutput := strings.Join(hosts, "|")

		tbl.AddRow(ip, hostsOutput)
	}

	tbl.Print()
}

func PrintCrtShList(dump *Dump) {
	tbl, _ := prettytable.NewTable([]prettytable.Column{
		{Header: "Certificate"}, {Header: "IP"}}...)

	for crt, ip := range dump.Certificates {
		tbl.AddRow(crt, ip)
	}

	tbl.Print()
}

func PrintTXTRecords(dump *Dump) {
	tbl, _ := prettytable.NewTable([]prettytable.Column{
		{}}...)

	for _, txt := range dump.TXT {
		tbl.AddRow(txt)
	}

	tbl.Print()
}

func PrintForwardARecord(dump *Dump) {
	tbl, _ := prettytable.NewTable([]prettytable.Column{
		{}}...)

	for _, ip := range dump.ForwardARecord {
		tbl.AddRow(ip)
	}

	tbl.Print()
}

func PrintMXRecords(dump *Dump) {
	tbl, _ := prettytable.NewTable([]prettytable.Column{
		{Header: "Host"},
		{Header: "Pref"}}...)

	for _, mx := range dump.MX {
		tbl.AddRow(mx.Host, fmt.Sprintf("%d", mx.Pref))
	}

	tbl.Print()
}

func PrintThreatcrowd(dump *Dump) {
	tbl, _ := prettytable.NewTable([]prettytable.Column{
		{Header: "Subdomain"}, {Header: "IP"}}...)

	for subdomain, ip := range dump.Subdomains {
		tbl.AddRow(subdomain, ip)
	}

	tbl.Print()

	fmt.Printf("\n\n")

	tbl, _ = prettytable.NewTable([]prettytable.Column{
		{Header: "Emails"}}...)

	for _, email := range dump.ThreatCrowdResult.Emails {
		tbl.AddRow(email)
	}

	tbl.Print()

	fmt.Printf("\n\n")

	tbl, _ = prettytable.NewTable([]prettytable.Column{
		{Header: "Emails"}}...)

	for _, email := range dump.ThreatCrowdResult.Emails {
		tbl.AddRow(email)
	}

	tbl.Print()

	fmt.Printf("\n\n")

	tbl, _ = prettytable.NewTable([]prettytable.Column{
		{Header: "Voted Malware"}, {Header: "More Infos"}}...)

	for _, hash := range dump.ThreatCrowdResult.Hashes {
		tbl.AddRow(hash, fmt.Sprintf("https://www.threatcrowd.org/malware.php?md5=%s", hash))
	}

	tbl.Print()
}
