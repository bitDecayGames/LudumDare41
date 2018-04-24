export GOOS=windows
go build -o Tanxz-server.exe

export GOOS=linux
go build -o Tanxz-server-linux

export GOOS=darwin
go build -o Tanxz-server-osx