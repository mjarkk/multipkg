#!/bin/bash

cd ..
go install

sudo systemctl start docker
docker run \
  -v ~/go/bin/multipkg:/usr/bin/multipkg \
  --rm -ti pritunl/archlinux:latest /bin/bash

cd tests