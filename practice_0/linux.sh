#!/bin/sh

protoc --go_out=plugins=grpc:. proto/blockchain.proto

