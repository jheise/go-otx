go-otx
---

An implementation of the AlienVault Open Threat Exchange API in Go

Example
```
package main

import (
    "fmt"
    otx "github.com/jheise/go-otx/api"
)

func main() {
    apikey := "APIKEYGOESHERE"
    client := otx.NewClient(apikey)

    ipv4data := client.GetIPv4General("4.2.2.1")
    fmt.Println(ipv4data)
}
