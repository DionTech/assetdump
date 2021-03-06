package assetdump

import (
	"fmt"
	"sort"
	"strings"
)

func (dump *Dump) OutputHosts() {
	hosts := dump.GetHosts()

	//ordered map here
	keys := make([]string, 0, len(hosts))
	for k := range hosts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k)
	}
}

func (dump *Dump) GetHosts() map[string]bool {
	hosts := make(map[string]bool, 0)
	for _, hostList := range dump.HostLookup.IPs {
		for _, host := range hostList {
			clHosts := clearHost(host)
			for _, clHost := range clHosts {
				if _, exists := hosts[clHost]; exists == false {
					hosts[clHost] = true
				}
			}

		}
	}

	for subdomain := range dump.Subdomains {
		clHosts := clearHost(subdomain)
		for _, clHost := range clHosts {
			if _, exists := hosts[clHost]; exists == false {
				hosts[clHost] = true
			}
		}
	}

	for crt := range dump.Certificates {
		clHosts := clearHost(crt)
		for _, clHost := range clHosts {
			if _, exists := hosts[clHost]; exists == false {
				hosts[clHost] = true
			}
		}
	}

	return hosts
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
		fmt.Println(k)
	}

}

func clearHost(host string) []string {
	hosts := make([]string, 0)

	for _, el := range strings.Split(host, "\n") {
		prefixes := [4]string{
			"http", "https", "*"}

		for _, prefix := range prefixes {
			//do not know why, but HasPrefix not works correctly
			if strings.Contains(host, prefix) {
				el = strings.Replace(el, prefix, "", 1)
			}
			hosts = append(hosts, el)
		}
	}

	return hosts
}

/* not really used anymore
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
*/
