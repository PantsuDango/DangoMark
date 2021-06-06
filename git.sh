#!/bin/bash

git pull
path=$PWD

cd $path/service
num=$(netstat -nlpt | grep ":::3000" | grep -v "grep" | wc -l)
if [ $num -eq 1 ]; then
    kill -9 $(ps -ef | grep "DangoMark" | grep -v "grep" | awk '{print $2}')
    sleep 3
fi
nohup go run DangoMark.go &
sleep 3

num=$(netstat -nlpt | grep ":::3000" | grep -v "grep" | wc -l)
if [ $num -eq 1 ]; then
    echo "run DangoMark success"
else
    cd $path/app
    tail -n 10 nohup.out
    echo "run DangoMark fail"
fi