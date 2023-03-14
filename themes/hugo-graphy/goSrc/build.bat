SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
cd themes\hugo-graphy\goSrc
go build hugo-graphy.go