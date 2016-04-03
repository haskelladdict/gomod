// Copyright 2016 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gomod provides functionality simular to the GNU modulefiles
package main

import "strings"

var sep = ":"

// prependToEnv prepends to the existing environment for the requested
// environmental variables
func prependToEnv(env, newEnv Env, pre envSpec) {
	for k, v := range pre.Vars {
		vars := curEnvValue(k, env, newEnv)
		newEnv[k] = strings.Join(append([]string{v}, vars...), sep)
	}
}

// appendToEnv appends to the existing environment for the requested
// environmental variables
func appendToEnv(env, newEnv Env, app envSpec) {
	for k, v := range app.Vars {
		vars := curEnvValue(k, env, newEnv)
		newEnv[k] = strings.Join(append(vars, v), sep)
	}
}

// removFromEnv appends to the existing environment for the requested
// environmental variables
func removeFromEnv(env, newEnv Env, rem envSpec) {
	for k, v := range rem.Vars {
		vars := curEnvValue(k, env, newEnv)
		if len(vars) == 0 {
			return
		}
		i := findInSlice(vars, v)
		if i == -1 {
			return
		}
		newEnv[k] = strings.Join(append(vars[:i], vars[i+1:]...), sep)
	}
}

// curEnvValue returns a list with elements of the current value for the
// environmental variable k. If not is currently defined returns an empty
// list
func curEnvValue(k string, env, newEnv Env) []string {
	var val string
	if v, ok := newEnv[k]; ok {
		val = v
	} else if v, ok := env[k]; ok {
		val = v
	}
	if val == "" {
		return []string{}
	}
	return strings.Split(val, sep)
}

// findInSlice looks for a string in a string slice and returns the index
// of the first occurrence or -1 if not found
func findInSlice(arr []string, s string) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}
