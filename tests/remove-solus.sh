#!/bin/bash

cd ..
sudo eopkg install screenfetch --no-color
go install
echo "-----------------"
echo ""
sudo /home/mark/go/bin/multipkg remove screenfetch "$@"
cd tests