#!/bin/bash

outputDir=dist
projectName=fablab-door

rm -Rf ${outputDir}/*

env GOOS=linux GOARCH=arm go build -o ${outputDir}/${projectName}-arm
env GOOS=windows GOARCH=amd64 go build -o ${outputDir}/${projectName}-windows.exe