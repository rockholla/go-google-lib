#!/bin/bash

this_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $this_dir/../

if ! command -v mockery &> /dev/null; then
  go get github.com/vektra/mockery/v2/.../
fi
find ./mocks -mindepth 1 -maxdepth 1 -not -name .custom-mocks -exec rm -rf '{}' \;
mockery --all --keeptree
