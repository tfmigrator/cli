#!/usr/bin/env bash

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

git ls-files | grep -E ".*\.go$" | xargs gofumpt -l -s -w
