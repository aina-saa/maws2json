#!/usr/bin/env bash
version="0.0.2"
time=$(date)
hash=$(git log -n 1 --pretty=format:"%H")
hostname=$(hostname)
go build -ldflags="\
    -X 'github.com/aina-saa/maws2json/version.BuildVersion=$version' \
    -X 'github.com/aina-saa/maws2json/version.BuildTime=$time' \
    -X 'github.com/aina-saa/maws2json/version.BuildHost=$hostname' \
    -X 'github.com/aina-saa/maws2json/version.BuildSha=$hash'" .
./maws2json --version

# eof