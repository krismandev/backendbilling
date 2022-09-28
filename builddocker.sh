#!/bin/bash
if [ -z ${1} ]; then
  echo "Need version param";
  exit 1;
else
  ver=$1;
fi
builddate="$(date '+%Y-%m-%d_%H:%M:%S')"
echo $builddate

docker build -t backendbillingdashboard:$ver --build-arg VER=$ver --build-arg BUILDDATE=$builddate .
docker tag backendbillingdashboard:$ver mygit.imitra.com:5020/backendbillingdashboard:$ver
docker tag backendbillingdashboard:$ver mygit.imitra.com:5020/backendbillingdashboard:latest
docker push mygit.imitra.com:5020/backendbillingdashboard:$ver
docker push mygit.imitra.com:5020/backendbillingdashboard:latest
docker image prune --filter label=stage=builder --force
