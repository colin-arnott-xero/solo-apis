#!/bin/bash

set -e

# Sync the proto files from the Gloo Repository checked out at ../gloo
rsync -ax --exclude 'solo-kit.json' --exclude 'grpc/v*'  ../gloo/projects/gloo/api/  ./api/gloo/gloo.apis
rmdir api/gloo/gloo.apis/grpc
rsync -ax --exclude 'solo-kit.json'  ../gloo/projects/gateway/api/  ./api/gloo/gateway

# Create Enterprise Gloo directory and rename extauth.proto to auth_config.proto
mkdir -p ./api/gloo/enterprise.gloo.apis/v1
mv ./api/gloo/gloo.apis/v1/enterprise/options/extauth/v1/extauth.proto ./api/gloo/enterprise.gloo.apis/v1/auth_config.proto

#mv ./api/gloo/gloo ./api/gloo/gloo.apis
# Fix paths
for file in $(find api/gloo -type f | grep ".proto")
do

# Gloo API changes
  sed "s|gloo/projects/gloo/api|solo-apis/api/gloo/gloo.apis|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gloo/pkg/api|github.com/solo-io/solo-apis/pkg/api/gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
 # sed -e "s|/gloo/gloo|/gloo/gloo.apis|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s| gloo\.solo\.io| gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s|/gloo\.solo\.io|/gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s| \.gloo\.solo\.io| .gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
# Gateway API changes
  sed "s|gloo/projects/gateway/api|solo-apis/api/gloo/gateway|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|github.com/solo-io/gloo/projects/gateway/pkg/api|github.com/solo-io/solo-apis/pkg/api/gateway.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"

# Gloo Enterprise API changes
  sed "s|gloo.apis/v1/enterprise/options/extauth/v1/extauth.proto|enterprise.gloo.apis/v1/auth_config.proto|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed "s|gloo.solo.io/v1/enterprise/options/extauth|enterprise.gloo.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s| enterprise\.gloo\.solo\.io| enterprise.gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s|/enterprise\.gloo.solo.io|/enterprise.gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s|/enterprise\.gloo/|/enterprise.gloo.apis/|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
  sed -e "s| \.enterprise.gloo\.solo\.io| .enterprise.gloo.apis.solo.io|g" "$file" > "$file".tmp && mv "$file".tmp "$file"
done

# convert protos used by skv1 into protos that can be used by skv2.
go run hack/convert.go