#!/bin/bash

cd ..
sudo eopkg remove screenfetch --no-color
go install
echo "-----------------"
echo ""
sudo /home/mark/go/bin/multipkg install screenfetch "$@"
cd tests