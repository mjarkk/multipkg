#!/bin/bash

cd ..
sudo eopkg remove gitg --no-color
go install
echo "-----------------"
echo ""
sudo /home/mark/go/bin/multipkg install gitg "$@"
cd tests