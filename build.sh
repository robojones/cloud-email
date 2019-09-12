#!/usr/bin/env bash

docker run -v `pwd`:/defs namely/protoc-all -d def/email -l go -o lib/email --go-source-relative
