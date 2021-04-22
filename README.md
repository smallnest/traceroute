# Traceroute in Go

A traceroute library written in Go.

[![Build Status](https://api.travis-ci.org/smallnest/traceroute.svg)](https://travis-ci.org/smallnest/traceroute)
[![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/traceroute)](https://goreportcard.com/report/github.com/smallnest/traceroute)
[![GoDoc](https://godoc.org/github.com/smallnest/traceroute?status.svg)](https://godoc.org/github.com/smallnest/traceroute)

Forked from [aeden/traceroute](https://github.com/aeden/traceroute).

## CLI App

```sh
go install github.com/smallnest/traceroute/cmd/tracert
sudo tracert example.com
```

## Library

See the code in cmd/gotraceroute.go for an example of how to use the library from within your application.

The traceroute.Traceroute() function accepts a domain name and an options struct and returns a TracerouteResult struct that holds an array of TracerouteHop structs.

## Resources

Useful resources:

* http://en.wikipedia.org/wiki/Traceroute
* http://tools.ietf.org/html/rfc792
* http://en.wikipedia.org/wiki/Internet_Control_Message_Protocol

## Notes

* https://code.google.com/p/go/source/browse/src/pkg/net/ipraw_test.go
* http://godoc.org/code.google.com/p/go.net/ipv4
* http://golang.org/pkg/syscall/
