// Package gomod provides functionality simular to the GNU modulefiles
//
// (C) Markus Dittrich, 2016
package main

import "fmt"

// description captures the module description
type description struct {
	Long  string
	Short string
}

// prependEnvSpec captures environmental variables to be appended to
type prependEnvSpec struct {
	Vars map[string]string
}

// setEnvSpec captures environmental variables to be set
type setEnvSpec struct {
	Vars map[string]string
}

// loadModulesSpec captures additional module to be loaded
type loadModulesSpec struct {
	Vars map[string]string
}

// Module captures the information for a given module
type Module struct {
	Desc        description
	PrependEnv  prependEnvSpec
	SetEnv      setEnvSpec
	LoadModules loadModulesSpec
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
		vMap := v.(map[string]interface{})
		switch k {
		case "description":
			if i, o := vMap["long"]; o {
				m.Desc.Long = i.(string)
			}
			if i, o := vMap["short"]; o {
				m.Desc.Short = i.(string)
			}
		case "prependEnv":
			m.PrependEnv.Vars, ok = parseVars(vMap)
			if !ok {
				return fmt.Errorf("parse error in [PrependEnv]")
			}
		case "setEnv":
			m.SetEnv.Vars, ok = parseVars(vMap)
			if !ok {
				return fmt.Errorf("parse error in [SetEnv]")
			}
		case "loadModules":
			m.LoadModules.Vars, ok = parseVars(vMap)
			if !ok {
				return fmt.Errorf("parse error in [LoadModules]")
			}
		}
	}
	return nil
}

func parseVars(m map[string]interface{}) (map[string]string, bool) {
	vars := make(map[string]string)
	var ok bool
	for kk, vv := range m {
		if vars[kk], ok = vv.(string); !ok {
			return nil, false
		}
	}
	return vars, true
}

//
