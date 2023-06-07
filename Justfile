export GOBIN := join(env_var("PWD"), "bin")
export PATH := GOBIN + ":" + env_var("PATH")

tools:
	go install -tags tools github.com/google/ko

image: tools
	ko build -P

run:
	#!/usr/bin/env bash
	set -euxo pipefail
	image=`just image`
	kubectl run --rm  -it --restart=Never --image $image updatecli-demo
