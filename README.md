# Ledger
[![Build Status](https://travis-ci.org/gomicro/ledger.svg)](https://travis-ci.org/gomicro/ledger)
[![Coverage](http://gocover.io/_badge/github.com/gomicro/ledger)](http://gocover.io/github.com/gomicro/ledger)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/ledger)](https://goreportcard.com/report/github.com/gomicro/ledger)
[![GoDoc](https://godoc.org/github.com/gomicro/ledger?status.png)](https://godoc.org/github.com/gomicro/ledger)

Ledger is a threadsafe, minimalist layer on top of native Go logging with the ability to write to more than standard out and honor log level thresholds.

# Example

## Standard Logger

```
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

```
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
