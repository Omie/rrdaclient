rrdaclient
==========

Go bindings for RRDA https://github.com/fcambus/rrda

Go version support yet to be tested, coded with 1.3.1 linux/amd64

## Description

[https://github.com/fcambus/rrda](https://github.com/fcambus/rrda)

RRDA is a REST API written in Go allowing to perform DNS queries over HTTP, and to get reverse PTR records for both IPv4 and IPv6 addresses. It outputs JSON-encoded DNS responses.

RRDA is a recursive acronym for "RRDA REST DNS API".

This library (rrdaclient) provides Go bindings to consume this API

## How to use

    //get library
    $ go get github.com/omie/rrdaclient

    //import library
    import github.com/omie/rrdaclient

    //call functions
    response = rrdaclient.GetHostAddress("<input>")

    //Response is of type *rrdaclient.APIResponse
    //and the structure resembles the RRDA JSON schema


## Further Reading
[http://www.statdns.com/api/](http://www.statdns.com/api/)


## License
MIT License
