#go run CloudDriverManager.go -driver=/tmp/Aws2Driver.so
echo "driver re-build"
$CB_SPIDER_ROOT/cloud-driver/drivers/aws2/build_driver_lib.sh

echo "Call DynamicPluginDriverPoc.go"
go run DynamicPluginDriverPoc.go -driver=/tmp/Aws2Driver.so
