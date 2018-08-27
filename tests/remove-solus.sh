#!/bin/bash

cd ..
sudo eopkg install gitg --no-color
go install
echo "-----------------"
echo ""
sudo /home/mark/go/bin/multipkg remove gitg "$@"
cd tests