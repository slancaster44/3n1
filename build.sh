#!/usr/bin/bash

cd 3n1-0.1_amd64/usr/src/3n1
go build
mv 3n1 ../../bin
cd ../../../..

dpkg-deb --build --root-owner-group 3n1-0.1_amd64