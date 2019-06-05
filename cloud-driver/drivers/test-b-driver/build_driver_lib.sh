rm -rf /tmp/TestBDriver.so
go build -buildmode=plugin TestBDriver.go
chmod +x TestBDriver.so
mv ./TestBDriver.so /tmp
