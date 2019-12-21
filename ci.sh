#!/usr/bin/env bash

VERSION=$(perl -ne 'print if /(?<=##\s)\d+\.\d+\.\d+/' CHANGELOG.md | awk 'NR==1 {print $2}')

git tag -a "${VERSION}" -m "Release ${VERSION}"
git push origin master --tags