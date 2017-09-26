# hostutils

[![Build Status](https://travis-ci.org/Wing924/hostutils.svg?branch=addStrfunc)](https://travis-ci.org/Wing924/hostutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/Wing924/hostutils)](https://goreportcard.com/report/github.com/Wing924/hostutils)

A golang library for packing and unpacking hosts list

## Functions

```
func Normalize(hosts []string) (packedHosts []string)
    Normalize Unpack and pack hosts

func NormalizeString(hosts string) (packedHosts []string)
    NormalizeString Unpack and pack hosts

func Pack(hosts []string) (packedHosts []string)
    Pack Pack full hosts list into short abbreviated hosts.

func PackString(hosts string) (packedHosts []string)
    PackString Pack space septated full hosts list into short abbreviated
    hosts.

func Unpack(packedHosts []string) (hosts []string)
    Unpack Unpack short abbreviated hosts into full hosts list.

func UnpackString(packedHosts string) (hosts []string)
    Unpack Unpack space septated short abbreviated hosts into full hosts
    list.
```
