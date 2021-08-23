#!/bin/bash

# exit once any command fails
set -e

if [[ $# -eq 0 ]]; then
	echo "Specify the domains for the certificaion."
	exit 1
fi

# sudo update-ca-certificates

script_dir=$(dirname "${BASH_SOURCE[0]}")
certs_dir=${script_dir}/certs

mkdir -p ${certs_dir}/
${script_dir}/mkcert -install
${script_dir}/mkcert -key-file ${certs_dir}/ssl.key -cert-file ${certs_dir}/ssl.crt $@ # localhost 127.0.0.1
