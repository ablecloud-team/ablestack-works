#!/bin/bash
version=test-04
srcpath=samba-ad
image=smb
if [ $(which yum) ]; then
  PM="yum"
elif [ $(which apt) ]; then
  PM="apt"
fi

if [ $(which docker) ]; then
  CM="docker"
elif [ $(which podman) ]; then
  CM="podman"
else
  echo "No container build system found."
  echo "please install podman or docker"
  exit
fi

if [ ! $(which go) ]; then
  echo "not found go"
  if [ $PM=="yum" ]; then
    echo "install with yum"
    yum install golang -y
  elif [ $PM=="apt" ]; then
    echo "install with apt"
    apt install golang -y
  fi
fi


GOOS=linux GOARCH=amd64 go build -o $srcpath/samba-ad .
pushd $srcpath

if [ $CM=="podman" ]; then
    podman build . -t $image:$version
elif [ $CM=="docker" ]; then
    docker build . -t $image:$version
fi
