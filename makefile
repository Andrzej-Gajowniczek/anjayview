build: main.go
	@echo "Removing old anjayview binary"
	@rm -f /usr/bin/anjayview
	@echo "Compiling binary..."
	/usr/local/go/bin/go build .
	@echo "Installing new anjayview binary"
	/usr/bin/install anjayview /usr/bin/anjayview
