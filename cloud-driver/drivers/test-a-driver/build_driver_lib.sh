rm -rf /tmp/TestADriver.so
go build -buildmode=plugin $CB_SPIDER_ROOT/cloud-driver/drivers/test-a-driver/TestADriver.go
chmod +x TestADriver.so
mv ./TestADriver.so /tmp
