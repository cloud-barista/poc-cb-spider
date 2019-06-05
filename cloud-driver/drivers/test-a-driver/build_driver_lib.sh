rm -rf /tmp/TestADriver.so
go build -buildmode=plugin TestADriver.go
chmod +x TestADriver.so
mv ./TestADriver.so /tmp
