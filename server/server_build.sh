mkdir -p artifacts/

export GOOS=windows
go build -o ./artifacts/Tanxz-server.exe

export GOOS=linux
go build -o ./artifacts/Tanxz-server-linux

export GOOS=darwin
go build -o ./artifacts/Tanxz-server-osx
