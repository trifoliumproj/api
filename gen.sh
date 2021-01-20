#!/usr/bin/sh
# docker pull namely/protoc-all
docker run -v `pwd`:/defs namely/protoc-all \
	-f nhentai.proto \
	-l go \
	-o generated
