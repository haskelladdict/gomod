#!/bin/bash

export MODULEPATH=/Users/markus/programming/go/src/markus_local/gomod/modulefiles

module() {
  eval $(./gomod bash "$@")
}
