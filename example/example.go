package main

import (
    "fmt"
    "github.com/omie/rrdaclient"
)

func printResp(addr *rrdaclient.APIResponse, err error) {
    if err != nil {
        fmt.Println("error: ", err)
    } else {
        if len(addr.Question)>0 {
            fmt.Printf("%s\t%s\n%s\n\n", addr.Question[0].Name, 
                                         addr.Question[0].Type, 
                                         addr)
        } else {
            fmt.Println(addr)
        }
    }
}

func main() {
    fmt.Println("starting ... ")
    var d string = "statdns.net"
    printResp(rrdaclient.GetHostAddress(d))
    printResp(rrdaclient.GetHostAddressV6(d))
    printResp(rrdaclient.GetCertificate(d))
    printResp(rrdaclient.GetCanonicalName("cname."+d))
    printResp(rrdaclient.GetDHCIDRecords(d))
    printResp(rrdaclient.GetDLVRecords(d))
    printResp(rrdaclient.GetDNAMERecords("dname."+d))
    printResp(rrdaclient.GetDNSKeyRecords(d))
    printResp(rrdaclient.GetDSRecords(d))
    printResp(rrdaclient.GetHINFORecords(d))
    printResp(rrdaclient.GetHIPRecords(d))
    printResp(rrdaclient.GetIPSECKeyRecords(d))
    printResp(rrdaclient.GetKXRecords(d))
    printResp(rrdaclient.GetLOCRecords(d))
    printResp(rrdaclient.GetMXRecords(d))
    printResp(rrdaclient.GetNAPTRRecords(d))
    printResp(rrdaclient.GetNSRecords(d))
    printResp(rrdaclient.GetNSECRecords(d))
    printResp(rrdaclient.GetNSEC3Records("hashvalue."+d))
    printResp(rrdaclient.GetNSEC3ParamRecords(d))
    printResp(rrdaclient.GetOPTRecords(d))
    printResp(rrdaclient.GetPTRRecords("8.8.8.8.in-addr.arpa"))
    printResp(rrdaclient.GetRRSIGRecords(d))
    printResp(rrdaclient.GetSOARecords(d))
    printResp(rrdaclient.GetSPFRecords(d))
    printResp(rrdaclient.GetSRVRecords("_sip._tcp."+d))
    printResp(rrdaclient.GetSSHFPRecords(d))
    printResp(rrdaclient.GetTARecords(d))
    printResp(rrdaclient.GetTALINKRecords(d))
    printResp(rrdaclient.GetTLSARecords("_443._tcp.www."+d))
    printResp(rrdaclient.GetTXTRecords(d))
    printResp(rrdaclient.GetReversePTRRecords("193.0.6.139"))
    printResp(rrdaclient.GetReversePTRRecordsV6("2001:67c:2e8:22::c100:68b"))
    fmt.Println("Done")
}


