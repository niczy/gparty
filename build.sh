#!/bin/bash
protoc -I proto/ proto/service.proto --go_out=plugins=grpc:server
