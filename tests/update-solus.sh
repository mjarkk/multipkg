#!/bin/bash

cd ..
go install
echo "-----------------"
echo ""
sudo /home/mark/go/bin/multipkg update "$@"
cd tests