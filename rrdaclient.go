package rrdaclient

import (
        "encoding/json"
        "io/ioutil"
        "net/http"
)

// base URL that points to RRDA REST API
var BaseURL string = "http://api.statdns.com/"

// represents the Question section in API response
type Question struct {
    Name    string  `json: "name"`
    Type    string  `json: "type"`
    Class   string  `json: "class"`
}

// represents Answer, Autority and Additional sections in API response
type DnsData struct {
    Name        string  `json: "name"`
    Type        string  `json: "type"`
    Class       string  `json: "class"`
    Ttl         int     `json: "ttl"`
    Rdlength    int     `json: "rdlength"`
    Rdata       string  `json: "rdata"`
}

// RRDA API Response
type APIResponse struct {
    Question    []Question `json: "question"`
    Answer      []DnsData  `json: "answer"`
    Authority   []DnsData  `json: "answer"`
    Additional  []DnsData  `json: "answer"`
}

// get data from remote url and return unparsed output
func getData(url string) ([]byte, error) {
        client := &http.Client{}
        req, err := http.NewRequest("GET", url, nil)
        resp, err := client.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

        return body, nil
}

// get data from API and return APIResponse object
func getAPIResponse(url string) (*APIResponse, error) {
    data, err := getData(url)
    if err != nil {
        return nil, err
    }

    var resp *APIResponse
    err = json.Unmarshal(data, &resp)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

/*  =====   RRDA API METHODS ===== */

// Get A Record
//
// http://api.statdns.com/<domain>/a
func GetHostAddress(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/a"
    return getAPIResponse(url)
}

//Get IPv6 Host Address (AAAA records)
//
//http://api.statdns.com/statdns.net/aaaa
func GetHostAddressV6(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/aaaa"
    return getAPIResponse(url)
}

//Get Certificate (CERT records)
//
//http://api.statdns.com/statdns.net/cert
func GetCertificate(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/cert"
    return getAPIResponse(url)
}

//Get Canonical Name (CNAME records)
//
//http://api.statdns.com/statdns.net/cname
func GetCanonicalName(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/cname"
    return getAPIResponse(url)
}

//Get DHCP Identifier (DHCID records)
//
//http://api.statdns.com/statdns.net/dhcid
func GetDHCIDRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/dhcid"
    return getAPIResponse(url)
}

//Get DNSSEC Lookaside Validation record (DLV records)
//
//http://api.statdns.com/statdns.net/dlv
func GetDLVRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/dlv"
    return getAPIResponse(url)
}

//Get Delegation name (DNAME records)
//
//http://api.statdns.com/statdns.net/dname
func GetDNAMERecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/dname"
    return getAPIResponse(url)
}

//Get DNS Key record (DNSKEY records)
//
//http://api.statdns.com/statdns.net/dnskey
func GetDNSKeyRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/dnskey"
    return getAPIResponse(url)
}

//Get Delegation Signer (DS records)
//
//http://api.statdns.com/statdns.net/ds
func GetDSRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/ds"
    return getAPIResponse(url)
}

//Get Host Information (HINFO records)
//
//http://api.statdns.com/statdns.net/hinfo
func GetHINFORecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/hinfo"
    return getAPIResponse(url)
}

//Get Host Identity Protocol (HIP records)
//
//http://api.statdns.com/statdns.net/hip
func GetHIPRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/hip"
    return getAPIResponse(url)
}

//Get IPSec Key (IPSECKEY records)
//
//http://api.statdns.com/statdns.net/ipseckey
func GetIPSECKeyRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/ipseckey"
    return getAPIResponse(url)
}

//Get Key eXchanger record (KX records)
//
//http://api.statdns.com/statdns.net/kx
func GetKXRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/kx"
    return getAPIResponse(url)
}

//Get Location record (LOC records)
//
//http://api.statdns.com/statdns.net/loc
func GetLOCRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/loc"
    return getAPIResponse(url)
}

//Get Mail Exchange record (MX records)
//
//http://api.statdns.com/statdns.net/mx
func GetMXRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/mx"
    return getAPIResponse(url)
}

//Get Name Authority Pointer (NAPTR records)
//
//http://api.statdns.com/statdns.net/naptr
func GetNAPTRRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/naptr"
    return getAPIResponse(url)
}

//Get Name Servers (NS records)
//
//http://api.statdns.com/statdns.net/ns
func GetNSRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/ns"
    return getAPIResponse(url)
}

//Get Next-Secure record (NSEC records)
//
//http://api.statdns.com/statdns.net/nsec
func GetNSECRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/nsec"
    return getAPIResponse(url)
}

//Get NSEC record version 3 (NSEC3 records)
//
//http://api.statdns.com/statdns.net/nsec3
func GetNSEC3Records(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/nsec3"
    return getAPIResponse(url)
}

//Get NSEC3 parameters (NSEC3PARAM records)
//
//http://api.statdns.com/statdns.net/nsec3param
func GetNSEC3ParamRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/nsec3param"
    return getAPIResponse(url)
}

//Get Option record (OPT records)
//
//http://api.statdns.com/statdns.net/opt
func GetOPTRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/opt"
    return getAPIResponse(url)
}

//Get Pointer record (PTR records)
//
//http://api.statdns.com/statdns.net/ptr
func GetPTRRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/ptr"
    return getAPIResponse(url)
}

//Get Resource Records Signature (RRSIG records)
//
//http://api.statdns.com/statdns.net/rrsig
func GetRRSIGRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/rrsig"
    return getAPIResponse(url)
}

//Get Start of Authority (SOA record)
//
//http://api.statdns.com/statdns.net/soa
func GetSOARecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/soa"
    return getAPIResponse(url)
}

//Get Sender Policy Framework (SPF records)
//
//http://api.statdns.com/statdns.net/spf
func GetSPFRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/spf"
    return getAPIResponse(url)
}

//Get Service Locator (SRV records)
//
//http://api.statdns.com/statdns.net/srv
func GetSRVRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/srv"
    return getAPIResponse(url)
}

//Get SSH Public Key Fingerprint (SSHFP records)
//
//http://api.statdns.com/statdns.net/sshfp
func GetSSHFPRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/sshfp"
    return getAPIResponse(url)
}

//Get DNSSEC Trust Authorities (TA records)
//
//http://api.statdns.com/statdns.net/ta
func GetTARecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/ta"
    return getAPIResponse(url)
}

//Get Trust Anchor LINK (TALINK records)
//
//http://api.statdns.com/statdns.net/talink
func GetTALINKRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/talink"
    return getAPIResponse(url)
}

//Get TLSA records
//
//http://api.statdns.com/_443._tcp.www.statdns.net/tlsa
func GetTLSARecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/tlsa"
    return getAPIResponse(url)
}

//Get Text record (TXT records)
//
//http://api.statdns.com/statdns.net/txt
func GetTXTRecords(domain string) (*APIResponse, error) {
    url := BaseURL + domain + "/txt"
    return getAPIResponse(url)
}

//Get reverse (PTR) record from IPv4 addresses
//
//http://api.statdns.com/x/193.0.6.139 
func GetReversePTRRecords(ip string) (*APIResponse, error) {
    url := BaseURL + "/x/" + ip
    return getAPIResponse(url)
}

//Get reverse (PTR) record from IPv6 addresses
//
//http://api.statdns.com/x/2001:67c:2e8:22::c100:68b
func GetReversePTRRecordsV6(ip string) (*APIResponse, error) {
    url := BaseURL + "/x/" + ip
    return getAPIResponse(url)
}


