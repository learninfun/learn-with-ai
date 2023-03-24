# Set variables for the build
VERSION = 0.1.0
BINARY_NAME = apidocx
SOURCE_DIR = cmd/$(BINARY_NAME)
BUILD_DIR = build/$(VERSION)
ARCHIVE_DIR = archive/$(VERSION)

# Define the targets
.PHONY: all clean
all: create_dir linux-amd64 linux-arm64 windows-amd64 darwin-amd64 # deb

create_dir:
	mkdir -p $(BUILD_DIR)
	mkdir -p $(ARCHIVE_DIR)

# Build for Linux amd64
linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-linux-amd64 $(SOURCE_DIR)/main.go
	tar -czvf $(ARCHIVE_DIR)/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-$(VERSION)-linux-amd64

# Build for Linux arm64
linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-linux-arm64 $(SOURCE_DIR)/main.go
	tar -czvf $(ARCHIVE_DIR)/$(BINARY_NAME)-$(VERSION)-linux-arm64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-$(VERSION)-linux-arm64

# Build for Windows amd64
windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe $(SOURCE_DIR)/main.go
#	zip -j $(ARCHIVE_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe
	tar -a -c -f $(ARCHIVE_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe

# Build for macOS amd64
darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-macOS $(SOURCE_DIR)/main.go
#	zip -j $(ARCHIVE_DIR)/$(BINARY_NAME)-$(VERSION)-macOS.zip $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-macOS
	tar -a -c -f $(ARCHIVE_DIR)/$(BINARY_NAME)-$(VERSION)-macOS.zip $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-macOS

# Build a debian package
#deb:
#	# Create the debian package files here
#	mkdir -p $(BUILD_DIR)/debian
#	mkdir -p $(BUILD_DIR)/usr/bin
#	cp $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-linux-amd64 $(BUILD_DIR)/usr/bin/$(BINARY_NAME)
#	fpm -s dir -t deb -n $(BINARY_NAME) -v $(VERSION) --deb-no-default-config-files --deb-changelog CHANGELOG.md --description "My awesome Golang program" --maintainer "me@example.com" -C $(BUILD_DIR)/debian .

# Clean up the build artifacts
clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf $(ARCHIVE_DIR)/*

#RUN:
#	go build -o apidocx.exe ./cmd/apidocx/main.go