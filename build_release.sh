#!/bin/bash
#
# build_release.sh of  TeleBeam from modul TeleBeam
# Created at 16.1.2022
# Created from: dpiening
# Last modified: 16.01.22, 18:15
# Copyright (C) 2021 - 2022 Dierk-Bent Piening & the TeleBeam Team.
#
#
# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
#
appname="TeleBeam"
version="$2"
echo ""
echo " GoReleaseBuilder - Compile and Release Tool"
echo " (C) 2021 - 2022 Dierk-Bent Piening"
echo """
# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
#
"""
echo " All Rights Reserved"
echo " E-Mail: dierk.bent.piening@icloud.com"
echo " -------------------------------------------"
echo ""
if [ -z "$1" ]; then
     echo "You need to specify Architecture to build for as first argument."
     exit 1 

elif [ -z "$2" ]; then 
    echo "You need to specify Version to build as second argument."
    exit 1
else
    RELEASEDIR="RELEASE-$version"
    if [ -d "$RELEASEDIR" ]; then
        echo " -> ERROR: $RELEASEDIR allready exists... "
        exit 1
    else
        echo " -> Creating Release Directory: $RELEASEDIR..."
        mkdir $RELEASEDIR
    fi
    echo " -> Cleaning Build Directory..."
    go clean
    echo ""
    echo " -> Loading Dependencies"
    go mod tidy
    echo ""
    echo " -> Cleaning Working Directory"
    echo ""
    echo " -> Starting to Compile and ZIP Release"
    echo ""
    echo " -> Build Windows version  $1..."
    GOOS=windows GOARCH=$1 go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-Windows-$1.zip"
    zip RELEASE/$zipname mailserv.exe conf/
    echo " -> Cleaning Build Root..."
    rm -rf mailserv.exe
    echo " Builded Windows Version..."
    echo ""
    echo " -> Build Linux version  $1..."
    GOOS=linux GOARCH=$1 go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-Linux-$1.zip"
    zip RELEASE/$zipname mailserv conf/
    echo " -> Cleaning Build Root..."
    rm -rf mailserv
    echo " -> Builded Linux version..."
    echo ""
    echo " -> Build MacOS version  $1..."
    GOOS=darwin GOARCH=$1 go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-MacOS-$1.zip"
    zip RELEASE/$zipname mailserv conf/
    echo " -> Cleaning Build Root..."
    rm -rf mailserv
    echo " -> Builded MacOS version..."
    echo ""
    echo " -> Build AIX version  $1..."
    GOOS=darwin GOARCH=$1 go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-AIX-$1.zip"
    zip RELEASE/$zipname mailserv conf/
    echo " -> Cleaning Build Root..."
    rm -rf mailserv
    echo " -> Builded AIX version..."
    echo ""
    echo " -> Build JS version WASM..."
    GOOS=js GOARCH=wasm go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-JS-WASM.zip"
    zip RELEASE/$zipname mailserv conf/
    echo " -> Cleaning Build Root..."
    rm -rf mailserv
    echo " -> Builded js version..."
    echo ""
    echo " -> Build Android version ARM 64Bit..."
    GOOS=android GOARCH=arm64 go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-Android-$1.zip" 
    zip RELEASE/$zipname mailserv conf/
    echo " -> Builded Android version..."
    echo ""
    echo " -> Build FreeBSD version  $1 ..."
    GOOS=freebsd GOARCH=$1 go build -a -v -ldflags "-s -w" .
    zipname="$appname-$version-FreeBSD-$1.zip"
    zip RELEASE/$zipname mailserv conf/
    echo " -> Cleaning Build Root..."
    rm -rf mailserv
    echo " -> Builded FreeBSD version..."
    echo ""
fi