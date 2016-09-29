package ipwhois

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResponse(t *testing.T) {
	in := `
{
  "raw": null,
  "entities": [
    "GOGL"
  ],
  "asn_registry": "arin",
  "network": {
    "status": null,
    "handle": "NET-8-8-8-0-1",
    "name": "LVLT-GOGL-8-8-8",
    "links": [
      "https://rdap.arin.net/registry/ip/008.008.008.000",
      "https://whois.arin.net/rest/net/NET-8-8-8-0-1",
      "https://rdap.arin.net/registry/ip/008.000.000.000/8"
    ],
    "raw": null,
    "country": "US",
    "ip_version": "v4",
    "start_address": "8.8.8.0",
    "notices": [
      {
        "description": "By using the ARIN RDAP/Whois service, you are agreeing to the RDAP/Whois Terms of Use",
        "links": [
          "https://www.arin.net/whois_tou.html"
        ],
        "title": "Terms of Service"
      }
    ],
    "end_address": "8.8.8.255",
    "remarks": null,
    "parent_handle": "NET-8-0-0-0-1",
    "cidr": "8.8.8.0/24",
    "type": "allocated",
    "events": [
      {
        "action": "last changed",
        "timestamp": "2014-03-14T16:52:05-04:00",
        "actor": null
      }
    ]
  },
  "objects": {
    "GOGL": {
      "status": [
        "validated"
      ],
      "roles": [
        "registrant"
      ],
      "handle": "GOGL",
      "entities": [
        "ABUSE5250-ARIN",
        "ZG39-ARIN"
      ],
      "links": [
        "https://rdap.arin.net/registry/entity/GOGL",
        "https://whois.arin.net/rest/org/GOGL"
      ],
      "raw": null,
      "notices": [
         {
            "description":"By using the ARIN RDAP/Whois service, you are agreeing to the RDAP/Whois Terms of Use",
            "links":[
               "https://www.arin.net/whois_tou.html"
            ],
            "title":"Terms of Service"
         }
      ],
      "contact": {
        "kind": "org",
        "name": "Google Inc.",
        "title": null,
        "role": null,
        "phone": [
          {
            "type": null,
            "value":"55 66 55093511"
          }
        ],
        "email": [
          {
            "type": null,
            "value":"AAC51@NUMERACAO.REGISTRO.BR"
          }
        ],
        "address": [
          {
            "type": null,
            "value": "1600 Amphitheatre Parkway Mountain View CA 94043 UNITED STATES"
          }
        ]
      },
      "events_actor": null,
      "remarks": null,
      "events": [
        {
          "action": "last changed",
          "timestamp": "2015-11-06T15:45:54-05:00",
          "actor": null
        }
      ]
    }
  },
  "asn_country_code": "US",
  "asn_date": "",
  "asn_cidr": "8.8.8.0/24",
  "nir": {
		 "query":"1.0.16.0",
		 "nets":[
				{
					 "updated":"2012-03-23T03:29:04",
					 "handle":"",
					 "name":"i2ts,inc.",
					 "contacts":{
							"admin":{
								 "division":"Engineering Department",
								 "fax":"03-5287-6276",
								 "updated":"2012-03-19T02:20:04",
								 "phone":"03-5287-6250",
								 "organization":"i2ts, inc.",
								 "email":"tech-support@i2ts.ne.jp"
							},
							"tech":{
								 "division":"Engineering Department",
								 "fax":"03-5287-6276",
								 "updated":"2012-03-19T02:20:04",
								 "phone":"03-5287-6250",
								 "organization":"i2ts, inc.",
								 "email":"tech-support@i2ts.ne.jp"
							}
					 },
					 "nameservers":null,
					 "country":"JP",
					 "created":null,
					 "range":"1.0.16.1 - 1.0.31.255",
					 "postal_code":null,
					 "address":null,
					 "cidr":"1.0.16.0/20"
				}
		 ],
		 "raw":null
	},
  "query": "8.8.8.8",
  "asn": "15169"
}
  `

	var out Response
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	expected := Response{
		ASN:            "15169",
		ASNCider:       "8.8.8.0/24",
		ASNCountryCode: "US",
		ASNDate:        "",
		ASNRegistry:    "arin",
		Entities:       []string{"GOGL"},
		Network: Network{
			Handle: "NET-8-8-8-0-1",
			Name:   "LVLT-GOGL-8-8-8",
			Links: []string{
				"https://rdap.arin.net/registry/ip/008.008.008.000",
				"https://whois.arin.net/rest/net/NET-8-8-8-0-1",
				"https://rdap.arin.net/registry/ip/008.000.000.000/8",
			},
			Country:      "US",
			IPVersion:    "v4",
			StartAddress: "8.8.8.0",
			EndAddress:   "8.8.8.255",
			Notices: []Notice{
				Notice{
					Description: "By using the ARIN RDAP/Whois service, you are agreeing to the RDAP/Whois Terms of Use",
					Links:       []string{"https://www.arin.net/whois_tou.html"},
					Title:       "Terms of Service",
				},
			},
			ParentHandle: "NET-8-0-0-0-1",
			CIDR:         "8.8.8.0/24",
			Type:         "allocated",
			Events: []Event{
				Event{
					Action:    "last changed",
					Timestamp: "2014-03-14T16:52:05-04:00",
				},
			},
		},
		Objects: map[string]Object{
			"GOGL": Object{
				Status: []string{"validated"},
				Roles:  []string{"registrant"},
				Handle: "GOGL",
				Entities: []string{
					"ABUSE5250-ARIN",
					"ZG39-ARIN",
				},
				Links: []string{
					"https://rdap.arin.net/registry/entity/GOGL",
					"https://whois.arin.net/rest/org/GOGL",
				},
				Notices: []Notice{
					Notice{
						Description: "By using the ARIN RDAP/Whois service, you are agreeing to the RDAP/Whois Terms of Use",
						Links:       []string{"https://www.arin.net/whois_tou.html"},
						Title:       "Terms of Service",
					},
				},
				Contact: Contact{
					Kind: "org",
					Name: "Google Inc.",
					Phones: []Phone{
						Phone{
							Value: "55 66 55093511",
						},
					},
					Addresses: []Address{
						Address{
							Value: "1600 Amphitheatre Parkway Mountain View CA 94043 UNITED STATES",
						},
					},
					Emails: []Email{
						Email{
							Value: "AAC51@NUMERACAO.REGISTRO.BR",
						},
					},
				},
				Events: []Event{
					Event{
						Action:    "last changed",
						Timestamp: "2015-11-06T15:45:54-05:00",
					},
				},
			},
		},
		NIR: NIR{
			Query: "1.0.16.0",
			Networks: []NIRNetwork{
				NIRNetwork{
					Range:   "1.0.16.1 - 1.0.31.255",
					CIDR:    "1.0.16.0/20",
					Handle:  "",
					Name:    "i2ts,inc.",
					Country: "JP",
					Contacts: map[string]NIRContact{
						"admin": NIRContact{
							Phone:        "03-5287-6250",
							Fax:          "03-5287-6276",
							Email:        "tech-support@i2ts.ne.jp",
							Division:     "Engineering Department",
							Organization: "i2ts, inc.",
							Updated:      "2012-03-19T02:20:04",
						},
						"tech": NIRContact{
							Phone:        "03-5287-6250",
							Fax:          "03-5287-6276",
							Email:        "tech-support@i2ts.ne.jp",
							Division:     "Engineering Department",
							Organization: "i2ts, inc.",
							Updated:      "2012-03-19T02:20:04",
						},
					},
					Updated: "2012-03-23T03:29:04",
				},
			},
		},
		Query: "8.8.8.8",
	}

	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("expeceted\n%v\nbut got\n%v", expected, out)
	}
}
