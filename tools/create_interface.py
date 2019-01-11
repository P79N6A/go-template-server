#! /usr/bin/env python

'''
Step for add new interface in grpc
1. write proto file
2. protoc -I . --go_out=plugins=grpc:. hello.proto
'''