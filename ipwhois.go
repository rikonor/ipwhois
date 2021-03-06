package ipwhois

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

var ErrRateLimit = errors.New("call was rate-limited")

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
from ipwhois.exceptions import HTTPRateLimitError

# Change the user agent
from urllib2 import build_opener
opener = build_opener()
opener.addheaders = [('User-Agent', '%s')]

obj = IPWhois("%s", proxy_opener=opener)
try:
	results = obj.lookup_rdap(depth=1, retry_count=0)
	print json.dumps(results)
except HTTPRateLimitError:
	print "rate-limit"
except:
	raise
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

	// defaultUserAgent is the default user-agent the request to rdap whois will be made with
	defaultUserAgent = "ipwhois"
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
	return LookupIPWithUA(ip, defaultUserAgent)
}

// LookupIPWithUA is similar to LookupIP but allows you to specify the User-Agent
// the request to the RDAP whois server will be made with
func LookupIPWithUA(ip string, ua string) (*Response, error) {
	if !PreqreqsMet {
		panic("calling `LookupIP` requires you have `python` and the `ipwhois` python package installed")
	}

	// call python's ipwhois
	s := fmt.Sprintf(pyWhoisQuery, ua, ip)
	strRes, err := execPythonScript(s)
	if err != nil {
		return nil, fmt.Errorf("call to py-whois failed: %s", err)
	}

	// Check if the python error is due to rate-limiting
	if strRes == "rate-limit\n" {
		return nil, ErrRateLimit
	}

	// convert string response to struct
	var res Response
	if err := json.Unmarshal([]byte(strRes), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
