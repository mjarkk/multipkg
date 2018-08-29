#!/bin/bash

cd ..
go install
echo "-----------------"
echo ""
sudo /home/mark/go/bin/multipkg info screenfetch "$@"
cd tests