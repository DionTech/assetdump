package assetdump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"

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

var fileSuffix = ".assetdump.json"

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

	fileName := "./" + dump.Domain + fileSuffix

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

func (dump *Dump) Load(path string, pretty bool) {
	fileName := dump.Domain + fileSuffix

	filepath.Walk(path, func(osPath string, f os.FileInfo, err error) error {
		f, err = os.Stat(osPath)

		// If no error
		if err != nil {
			return nil
		}

		// File & Folder Mode
		fMode := f.Mode()

		// Is folder
		if fMode.IsDir() {

		} else {
			if strings.Contains(osPath, fileName) {
				file, err := os.Open(filepath.FromSlash(osPath))

				if err != nil {
					stdoutformat.Printf("cannot open file %s", fileName)
					return nil
				}

				defer file.Close()

				byteValue, _ := ioutil.ReadAll(file)
				json.Unmarshal(byteValue, dump)

				if pretty {
					var prettyJSON bytes.Buffer
					json.Indent(&prettyJSON, byteValue, "", "\t")

					fmt.Println(string(prettyJSON.Bytes()))
				}

				return nil
			}

		}
		return nil
	})

	/*
		fileName := "./" + dump.Domain + fileSuffix

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
		}*/
}

func List(path string) {
	filepath.Walk(path, func(osPath string, f os.FileInfo, err error) error {
		f, err = os.Stat(osPath)

		// If no error
		if err != nil {
			return nil
		}

		// File & Folder Mode
		fMode := f.Mode()

		// Is folder
		if fMode.IsDir() {

		} else {
			if strings.Contains(osPath, fileSuffix) {
				fmt.Println(strings.Replace(osPath, fileSuffix, "", -1))
			}

		}
		return nil
	})
}
