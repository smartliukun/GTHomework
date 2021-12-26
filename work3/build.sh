#!/usr/bin/env bash

RUN_NAME='work3'


mkdir -p output/bin output/conf
cp script/bootstrap.sh 2>/dev/null

go build -o output/bin/${RUN_NAME}