#!/bin/bash

git pull
path=$PWD

cd $path/service
num=$(netstat -nlpt | grep ":::3000" | grep -v "grep" | wc -l)
if [ $num -eq 1 ]; then
    kill -9 $(ps -ef | grep "job-hunting" | grep -v "grep" | awk '{print $2}')
    sleep 3
fi
nohup go run job-hunting.go &
sleep 3

num=$(netstat -nlpt | grep ":::3000" | grep -v "grep" | wc -l)
if [ $num -eq 1 ]; then
    echo "run job-hunting success"
else
    cd $path/app
    tail -n 10 nohup.out
    echo "run job-hunting fail"
fi