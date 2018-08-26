#!/bin/bash

cd ..
sudo eopkg remove gitg --no-color
go install
sudo /home/mark/go/bin/multipkg install gitg "$@"
cd tests