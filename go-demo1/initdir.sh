#!/usr/bin/env bash
if [ $1 != "" ];then
    mkdir go-$1
    cd go-$1
    touch learn.go work.go
else
    echo "请给参数"
fi