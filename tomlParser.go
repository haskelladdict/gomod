// Copyright 2016 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gomod provides functionality simular to the GNU modulefiles
package main

import "fmt"

// description captures the module description
type description struct {
	Long  string
	Short string
}

// stringMap is a simple string map for key value pairs
type stringMap map[string]string

// envSpec captures environmental variables to be acted on
type envSpec struct {
	Vars stringMap
}

// aliasSpec captures shell aliasea to be acted on
type aliasSpec struct {
	Vars stringMap
}

// modulesSpec captures a list of additional modules
type modulesSpec struct {
	Vars []string
}

// Module captures the information for a given module
type Module struct {
	Desc         description
	AppendEnv    envSpec
	PrependEnv   envSpec
	RemoveEnv    envSpec
	SetEnv       envSpec
	UnsetEnv     envSpec
	SetAlias     aliasSpec
	UnsetAlias   aliasSpec
	LoadMods     modulesSpec
	ConflictMods modulesSpec
	PrereqMods   modulesSpec
}

// UnmarshalTOML knows how to parse a module file description in toml format
func (m *Module) UnmarshalTOML(data interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("parse error %q", r.(error))
		}
	}()

	dataMap, _ := data.(map[string]interface{})
	var ok bool
	for k, v := range dataMap {
		switch k {
		case "longDescription":
			m.Desc.Long = v.(string)
		case "shortDescription":
			m.Desc.Short = v.(string)
		case "prependEnv":
			vMap := v.(map[string]interface{})
			if m.PrependEnv.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [prependEnv]")
			}
		case "appendEnv":
			vMap := v.(map[string]interface{})
			if m.AppendEnv.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [appendEnv]")
			}
		case "removeEnv":
			vMap := v.(map[string]interface{})
			if m.PrependEnv.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [removeEnv]")
			}
		case "setEnv":
			vMap := v.(map[string]interface{})
			if m.SetEnv.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [setEnv]")
			}
		case "unsetEnv":
			vMap := v.(map[string]interface{})
			if m.UnsetEnv.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [unsetEnv]")
			}
		case "setAlias":
			vMap := v.(map[string]interface{})
			if m.SetAlias.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [setAlias]")
			}
		case "unsetAlias":
			vMap := v.(map[string]interface{})
			if m.UnsetAlias.Vars, ok = parseMapVars(vMap); !ok {
				return fmt.Errorf("parse error in [unsetAlias]")
			}
		case "loadMods":
			vArr := v.([]interface{})
			if m.LoadMods.Vars, ok = parseArrayVars(vArr); !ok {
				return fmt.Errorf("parse error in [loadMods]")
			}
		case "conflictMods":
			vArr := v.([]interface{})
			if m.ConflictMods.Vars, ok = parseArrayVars(vArr); !ok {
				return fmt.Errorf("parse error in [conflictMods]")
			}
		case "prereqMods":
			vArr := v.([]interface{})
			if m.PrereqMods.Vars, ok = parseArrayVars(vArr); !ok {
				return fmt.Errorf("parse error in [prereqMods]")
			}
		}
	}
	return nil
}

// parseMapVars parses a map[string]string from a map[string]interface{}
func parseMapVars(m map[string]interface{}) (map[string]string, bool) {
	vars := make(map[string]string)
	var ok bool
	for kk, vv := range m {
		if vars[kk], ok = vv.(string); !ok {
			return nil, false
		}
	}
	return vars, true
}

// parseArrayVars parses a []string from a []interface{}
func parseArrayVars(a []interface{}) ([]string, bool) {
	vars := make([]string, len(a))
	var ok bool
	for i, vv := range a {
		if vars[i], ok = vv.(string); !ok {
			return nil, false
		}
	}
	return vars, true
}

//
