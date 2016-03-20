// Package gomod provides functionality simular to the GNU modulefiles
//
// (C) Markus Dittrich, 2016
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatal("file name required")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("error opening file ", os.Args[1], ": ", err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("error reading file", os.Args[1], ": ", err)
	}

	var mod Module
	err = toml.Unmarshal(content, &mod)
	if err != nil {
		log.Fatal("failed to parse module file ", err)
	}
	fmt.Println(mod.Desc.Short)
	fmt.Println(mod.SetEnv.Vars)
	fmt.Println(mod.UnsetEnv.Vars)
}

//
