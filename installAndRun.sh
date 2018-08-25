#!/bin/bash

sudo eopkg remove gitg --no-color
go install
sudo /home/mark/go/bin/multipkg install gitg