package ledger_test

import (
	log "github.com/gomicro/ledger"
)

var lvl = log.InfoLevel

// ExampleExported demonstrates the default exported logger
func Example_exported() {
	log.Debug("We've got a debug line here")
	log.Info("We've got a info line here")
	log.Warn("We've got a warn line here")
	log.Error("We've got a error line here")
	log.Fatal("We've got a fatal line here")
}
