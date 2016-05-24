#!/bin/bash

export SRCDIR=$(pwd)

cd nvml
make
cd ..

go build gonvml.go

