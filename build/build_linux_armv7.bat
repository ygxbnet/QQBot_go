go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=arm
go env -w GOARM=7

go build

go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64
