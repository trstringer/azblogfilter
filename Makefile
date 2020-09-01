BIN_DIR=./bin
BIN=azblogfilter
BIN_DEBUG=$(BIN).debug
GCFLAGS_DEBUG="all=-N -l"
SYSTEMD_DIR=~/.config/systemd/user
INSTALL_LOCATION=~/bin

.PHONY: build build-debug test install install-systemd clean

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN)

build-debug:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_DEBUG) -gcflags=$(GCFLAGS_DEBUG)

test:
	go test -v ./...

install: build
	cp $(BIN_DIR)/$(BIN) $(INSTALL_LOCATION)/$(BIN)

install-systemd: install
	mkdir -p $(SYSTEMD_DIR)

clean:
	if [ -d $(BIN_DIR) ]; then rm -rf $(BIN_DIR); fi
	if [ -f $(INSTALL_LOCATION)/$(BIN) ]; then rm $(INSTALL_LOCATION)/$(BIN); fi
