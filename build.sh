#!/usr/bin/env bash

export GO111MODULE=on
REVISION=$(git rev-parse --short HEAD)
CHANGED=$(git diff-index --name-only HEAD --)
if [ ! -z "$CHANGED" ]; then
    SUFFIX="dirty"
else
    SUFFIX="clean"
fi

DATEVERSION=$(date +"%y.%m.%d")
VERSION="$DATEVERSION.$REVISION-$SUFFIX"

go build -v -ldflags "-X main.version=$VERSION"