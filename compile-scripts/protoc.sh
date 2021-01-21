#!/usr/bin/sh
cd "$(dirname $0)"

# @trifoliumproj/nh
docker run \
	--rm \
	-v `pwd`:`pwd` \
	-w `pwd` \
	thethingsindustries/protoc \
	--go_out=./generated/nhentai \
	--go-grpc_out=./generated/nhentai \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	-I. \
	./nhentai.proto
