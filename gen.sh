#!/usr/bin/sh
cd "$(dirname $0)"
docker pull namely/protoc-all
docker run -v `pwd`:/defs namely/protoc-all \
	-f nhentai.proto \
	-l go \
	-o generated/nhentai
docker run -v `pwd`:/defs namely/protoc-all \
	-f nhentai.proto \
	-l node \
	-o generated/nhentai \
	--with-typescript
