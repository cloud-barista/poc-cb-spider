rm -rf /tmp/TestBDriver.so
go build -buildmode=plugin $CB_SPIDER_ROOT/cloud-driver/drivers/test-b-driver/TestBDriver.go
chmod +x TestBDriver.so
mv ./TestBDriver.so /tmp
