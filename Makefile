VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1` | cut -c2-6)
DESCRIPTION="Command line to easily send messages to RabbitMQ"
PACKAGE=amqpsender

gox:
	gox -output "build/{{.OS}}_{{.Arch}}/{{.Dir}}" -osarch "darwin/amd64 linux/386 linux/amd64"

packages:
	fpm -f -a amd64 -p build/linux_amd64/  -s dir -t deb -n $(PACKAGE) --description $(DESCRIPTION) -v $(VERSION) ./build/linux_amd64/amqpsender=/usr/bin/amqpsender
	fpm -f -a i386 -p build/linux_386/  -s dir -t deb -n $(PACKAGE) --description $(DESCRIPTION) -v $(VERSION) ./build/linux_386/amqpsender=/usr/bin/amqpsender
	fpm -f -a x86_64 -p build/linux_amd64/  -s dir -t rpm -n $(PACKAGE) --description $(DESCRIPTION) -v $(VERSION) ./build/linux_amd64/amqpsender=/usr/bin/amqpsender
	fpm -f -a i386 -p build/linux_386/  -s dir -t rpm -n $(PACKAGE) --description $(DESCRIPTION) -v $(VERSION) ./build/linux_386/amqpsender=/usr/bin/amqpsender
	tar -czvf build/linux_amd64/linux_amd64.tar.gz -C build/linux_amd64 amqpsender
	tar -czvf build/linux_386/linux_386.tar.gz -C build/linux_386 amqpsender
	tar -czvf build/darwin_amd64/darwin_amd64.tar.gz -C build/darwin_amd64 amqpsender

packagecloud:
	package_cloud push rodrigosaito/pkgs/ubuntu/trusty ./build/linux_amd64/*.deb
	package_cloud push rodrigosaito/pkgs/ubuntu/trusty ./build/linux_386/*.deb

clean:
	rm -rf build

test:
	go test
