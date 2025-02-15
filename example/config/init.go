// Code generated by genfig on 2019-07-24T22:16:50+02:00; DO NOT EDIT.

package config

import (
	"fmt"
	"os"
)

var (
	_ = os.Getenv
	_ = fmt.Printf
)

// Current is the current config, selected by the curren env and
// updated by the availalbe env vars
var Current *Config

// This init tries to retrieve the current environment via the
// common env var 'ENV' and applies activated plugins
func init() {
	Current, _ = Get(os.Getenv("ENV"))

	// calling plugin substitutor
	Current.Substitute()

	// calling plugin update_from_env
	if errs := Current.UpdateFromEnv(); errs != nil {
		fmt.Println(errs)
	}

}
