#!/usr/bin/env bash

set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd "$script_dir/../.." && pwd)"
dist_dir="${DIST_DIR:-$repo_root/dist/lambda}"
binary_name="appbuilder-grader-lambda"
zip_path="$dist_dir/${binary_name}.zip"

mkdir -p "$dist_dir"

pushd "$repo_root" >/dev/null
GOOS=linux GOARCH=amd64 go build -o "$dist_dir/bootstrap" ./cmd/lambda
pushd "$dist_dir" >/dev/null
zip -q -j "$zip_path" bootstrap
popd >/dev/null
popd >/dev/null

echo "Created $zip_path"