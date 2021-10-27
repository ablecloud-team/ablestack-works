#!/bin/bash
version=v3
srcpath=/root/samba-ad4/container
image=works-ad

if command -v yum &> /dev/null; then
  PM="yum"
elif command -v apt &> /dev/null; then
  PM="apt"
fi

if command -v docker &> /dev/null; then
  CM="docker"
elif command -v podman &> /dev/null; then
  CM="podman"
else
  echo "No container build system found."
  echo "please install podman or docker"
  exit
fi

if ! command -v go &> /dev/null; then
  echo "not found go"
  if [ $PM == "yum" ]; then
    echo "install with yum"
    yum install golang -y
  elif [ $PM == "apt" ]; then
    echo "install with apt"
    apt install golang -y
  fi
fi


GOOS=linux GOARCH=amd64 go build -o $srcpath/samba-ad .
chmod +x $srcpath/*

pushd $srcpath

if [ $CM == "podman" ]; then
    podman build . -t $image:$version
elif [ $CM == "docker" ]; then
    docker build . -t $image:$version
fi
