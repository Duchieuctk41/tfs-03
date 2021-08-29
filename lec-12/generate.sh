#!bin/bash

protoc greet/greetpb/greetpb.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculatorpb.proto --go_out=plugins=grpc:.