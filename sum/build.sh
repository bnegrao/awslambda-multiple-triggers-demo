#!/bin/bash

function die {
    declare MSG="$@"
    echo -e "$0: Error: $MSG">&2
    exit 1
}

appNames=(api-gateway-handler direct-invoke-handler)

rootAppDir=$(pwd)

for app in ${appNames[@]}; do
    echo "Building app $app"
    appDir="handlers/$app/main"
    cd "$appDir" || die "dir '$appDir' does not exist"
    go get
    go build -o ${app}.goex ${app}.go || die "go build has failed"
    cd $rootAppDir || die "can't cd to $rootAppDir"
done