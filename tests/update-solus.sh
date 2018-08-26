#!/bin/bash

cd ..
go install
sudo /home/mark/go/bin/multipkg update "$@"
cd tests