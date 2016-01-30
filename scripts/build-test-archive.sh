#!/bin/sh

set -ex

pushd `dirname $0` > /dev/null
base=$(pwd -P)
popd > /dev/null

export GOPATH=$base/gogo
mkdir -p $GOPATH

###

go get github.com/venicegeo/pz-logger

go get github.com/stretchr/testify/suite

#go test -v github.com/venicegeo/pz-logger

go install github.com/venicegeo/pz-logger

###

exe=$GOPATH/bin/pz-logger

# gather some data about the repo
source $base/vars.sh

# do we have this artifact in s3? If not, upload it.
aws s3 ls $S3URL || aws s3 cp $exe $S3URL

# success if we have an artifact stored in s3.
aws s3 ls $S3URL
