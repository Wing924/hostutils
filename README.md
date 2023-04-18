# hostutils

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/Wing924/hostutils/v2)](https://goreportcard.com/report/github.com/Wing924/hostutils/v2)
[![codecov](https://codecov.io/gh/Wing924/hostutils/branch/master/graph/badge.svg)](https://codecov.io/gh/Wing924/hostutils)
[![GoDoc](https://godoc.org/github.com/Wing924/hostutils/v2?status.svg)](https://pkg.go.dev/github.com/Wing924/hostutils/v2)

A golang library for packing and unpacking hosts list

## Install

```bash
go get github.com/Wing924/hostutils/v2
```

## Examples

```go
package main

import (
    "fmt"

    "github.com/Wing924/hostutils/v2"
)

func main() {
  // Pack
  pack1 := hostutils.Pack([]string{"example101z.com", "example102z.com", "example103z.com"})
  fmt.Println(pack1) // [example[101-103]z.com]

  pack2 := hostutils.Pack([]string{"example101z.com", "example102z.com", "example201z.com"})
  fmt.Println(pack2) // [example[101-102,201]z.com]

  pack3 := hostutils.Pack([]string{"example01z.com example02z.com"})
  fmt.Println(pack3) // [example[01-02]z.com]

  // Unpack
  unpack1 := hostutils.Unpack([]string{"example[101-103]z.com"})
  fmt.Println(unpack1) // [example101z.com example102z.com example103z.com]

  unpack2 := hostutils.Unpack([]string{"example[1-2][101-102]z.com"})
  fmt.Println(unpack2) // [example1101z.com example1102z.com example2101z.com example2102z.com]
}
```

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
