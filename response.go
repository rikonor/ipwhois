package ipwhois

type Response struct {
	ASN            string            `json:"asn"`
	ASNCider       string            `json:"asn_cidr"`
	ASNCountryCode string            `json:"asn_country_code"`
	ASNDate        string            `json:"asn_date"`
	ASNRegistry    string            `json:"asn_registry"`
	Entities       []string          `json:"entities"`
	Network        Network           `json:"network"`
	Objects        map[string]Object `json:"objects"`
	NIR            NIR               `json:"nir"`
	Query          string            `json:"query"`
	Raw            string            `json:"raw"`
}

type Network struct {
	Status       string   `json:"status"`
	Handle       string   `json:"handle"`
	Name         string   `json:"name"`
	Links        []string `json:"links"`
	Country      string   `json:"country"`
	IPVersion    string   `json:"ip_version"`
	StartAddress string   `json:"start_address"`
	EndAddress   string   `json:"end_address"`
	Notices      []Notice `json:"notices"`
	Remarks      []string `json:"remarks"`
	ParentHandle string   `json:"parent_handle"`
	CIDR         string   `json:"cidr"`
	Type         string   `json:"type"`
	Events       []Event  `json:"events"`
	Raw          string   `json:"raw"`
}

type Notice struct {
	Description string   `json:"description"`
	Links       []string `json:"links"`
	Title       string   `json:"title"`
}

type Event struct {
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
	Actor     string `json:"actor"`
}

type Object struct {
	Status      []string `json:"status"`
	Roles       []string `json:"roles"`
	Handle      string   `json:"handle"`
	Entities    []string `json:"entities"`
	Links       []string `json:"links"`
	Notices     []Notice `json:"notices"`
	Contact     Contact  `json:"contact"`
	EventsActor string   `json:"events_actor"`
	Remarks     []string `json:"remarks"`
	Events      []Event  `json:"events"`
	Raw         string   `json:"raw"`
}

type Contact struct {
	Kind      string    `json:"kind"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Role      string    `json:"role"`
	Phones    []Phone   `json:"phone"`
	Addresses []Address `json:"address"`
	Emails    []Email   `json:"email"`
}

type Address struct {
	Value string `json:"value"`
}

type Phone struct {
	Value string `json:"value"`
}

type Email struct {
	Value string `json:"value"`
}

type NIR struct {
	Query    string       `json:"query"`
	Networks []NIRNetwork `json:"nets"`
	Raw      string       `json:"raw"`
}

type NIRNetwork struct {
	Range       string                `json:"range"`
	CIDR        string                `json:"cidr"`
	Handle      string                `json:"handle"`
	Name        string                `json:"name"`
	NameServers []string              `json:"nameservers"`
	Address     string                `json:"address"`
	PostalCode  string                `json:"postal_code"`
	Country     string                `json:"country"`
	Contacts    map[string]NIRContact `json:"contacts"`
	Created     string                `json:"created"`
	Updated     string                `json:"updated"`
}

type NIRContact struct {
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
	Division     string `json:"division"`
	Organization string `json:"organization"`
	Updated      string `json:"updated"`
}
