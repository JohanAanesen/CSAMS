#!/bin/bash

export PEER_AUTH="5243980712315079823517089"

echo -e "\n\n-------------------GOLINT-------------------"
golint ./...

echo -e "\n\n-------------------GO VET-------------------"
go vet ./webservice/...
go vet ./peerservice/...

echo -e "\n\n-------------------GO FMT-------------------"
go fmt ./webservice/...
go fmt ./peerservice/...

echo -e "\n\n-------------------GOCYCLO-------------------"
gocyclo ./webservice/...
gocyclo ./peerservice/...

echo -e "\n\n-------------------GO TEST WEBSERVICE-------------------"
go test ./webservice/... -cover

echo -e "\n\n-------------------GO TEST PEERSERVICE-------------------"
go test ./peerservice/... -cover

echo -e "\n\n-------------------GO TEST SCHEDULERSERVICE-------------------"
go test ./schedulerservice/... -cover