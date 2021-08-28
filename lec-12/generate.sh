#!bin/bash

protoc greet/greetpb/greetpb.proto --go_out=plugins=grpc:.