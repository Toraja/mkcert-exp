#!/bin/bash

# exit once any command fails
set -e

# seems like libnss3-tools is not necessary if binary file is downloaded from github

version=${1:-v1.4.3}
cpu_arch=${2:-amd64}
script_dir=$(dirname "${BASH_SOURCE[0]}")
curl -fsSL -o ${script_dir}/mkcert https://github.com/FiloSottile/mkcert/releases/download/${version}/mkcert-${version}-linux-${cpu_arch}
chmod 755 ${script_dir}/mkcert
