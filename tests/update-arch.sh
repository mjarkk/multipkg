#!/bin/bash

cd ..
go install

docker run \
  -v ~/go/bin/multipkg:/usr/bin/multipkg \
  --rm -ti pritunl/archlinux:latest /usr/bin/multipkg update

cd tests