# Ledger
[![Build Status](https://travis-ci.org/gomicro/ledger.svg)](https://travis-ci.org/gomicro/ledger)
[![GoDoc](https://godoc.org/github.com/gomicro/ledger?status.png)](https://godoc.org/github.com/gomicro/ledger)

Minimalist Golang log with levels added

# Example

## Standard Logger

```go
package main

import(
	log "github.com/gomicro/ledger"
)

func main(){
	log.Debug("We've got a debug line here")
	log.Info("We've got a info line here")
	log.Warn("We've got a warn line here")
	log.Error("We've got a error line here")
	log.Fatal("We've got a fatal line here")
}
```

## Global Logger

```go
package main

import(
	"os"

	"github.com/gomicro/ledger"
)

var (
	log Ledger
)

func init(){
	log = ledger.New(os.Stdout, ledger.DebugLevel)
}

func main(){
	log.Debug("We've got a debug line here")
	log.Info("We've got a info line here")
	log.Warn("We've got a warn line here")
	log.Error("We've got a error line here")
	log.Fatal("We've got a fatal line here")
}
```
