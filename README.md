# golyse
Go CLI to ping products and show data in graphics

## Install
### Windows
```bat
mkdir "%APPDATA%/golyse"

go build -o %APPDATA%/golyse/golyse-server.exe server.go
go build -o %APPDATA%/golyse/golyse.exe cli.go

setx /M PATH "%PATH%;%APPDATA%/golyse"
```
