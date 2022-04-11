#!/bin/bash
if [ -z ${1} ]; then
  echo "Need version param";
  exit 1;
else
  ver=$1;
fi
builddate="$(date '+%Y-%m-%d_%H:%M:%S')"
echo $builddate

docker build -t backendbillingdashbord:$ver --build-arg VER=$ver --build-arg BUILDDATE=$builddate .
docker tag backendbillingdashbord:$ver mygit.imitra.com:5030/backendbillingdashbord:$ver
docker tag backendbillingdashbord:$ver mygit.imitra.com:5030/backendbillingdashbord:latest
docker push mygit.imitra.com:5030/backendbillingdashbord:$ver
docker push mygit.imitra.com:5030/backendbillingdashbord:latest
docker image prune --filter label=stage=builder --force
