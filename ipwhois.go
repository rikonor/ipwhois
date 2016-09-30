package ipwhois

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var pyCheckImport = `\
try:
  import %s
except ImportError, e:
  print "not installed"
`

var pyWhoisQuery = `\
import sys
import json
from ipwhois import IPWhois

obj = IPWhois("%s")
results = obj.lookup_rdap(depth=1)
print json.dumps(results)
`

// execPythonScript executes a given python script
func execPythonScript(script string) (string, error) {
	cmd := exec.Command("python", "-c", script)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func checkPythonInstalled() bool {
	res, err := execPythonScript(`print "Hello"`)
	if err != nil || res != "Hello\n" {
		return false
	}
	return true
}

func checkPythonPackageInstalled(name string) bool {
	s := fmt.Sprintf(pyCheckImport, name)
	res, err := execPythonScript(s)
	if err != nil || res == "not installed\n" {
		return false
	}
	return true
}

var (
	// PreqreqsMet means that the prerequisites for running this package have been met
	// meaning python and py-ipwhois are installed
	PreqreqsMet = false
)

func init() {
	// Check if python is present
	// Check if py-ipwhois is present
	if checkPythonInstalled() && checkPythonPackageInstalled("ipwhois") {
		PreqreqsMet = true
	}
}

// LookupIP performs an ip whois lookup on the given ip and returns a similar result
// for all RIRs as specified by the py-ipwhois package
func LookupIP(ip string) (*Response, error) {
	if !PreqreqsMet {
		panic("calling `LookupIP` requires you have `python` and the `ipwhois` python package installed")
	}

	// call python's ipwhois
	s := fmt.Sprintf(pyWhoisQuery, ip)
	strRes, err := execPythonScript(s)
	if err != nil {
		return nil, fmt.Errorf("call to py-whois failed: %s", err)
	}

	// convert string response to struct
	var res Response
	if err := json.Unmarshal([]byte(strRes), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
