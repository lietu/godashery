@echo off

rem
rem Build and run the app
rem
rem So Windows stops nagging about firewall stuff all the time
rem like it does with "go run godashery.go"
rem

del godashery.exe
go build
godashery
