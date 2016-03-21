// Copyright 2016 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gomod provides functionality simular to the GNU modulefiles
package main

import (
	"log"
	"os"
	"strings"
)

// Env contains a map of the current environment
type Env map[string]string

// parseEnv parses the current environment into a type Env
func parseEnv() Env {
	env := make(Env)
	for _, s := range os.Environ() {
		items := strings.Split(s, "=")
		if len(items) != 2 {
			log.Print("failed to parse environmental variable ", s)
		}
		env[items[0]] = items[1]
	}
	return env
}
