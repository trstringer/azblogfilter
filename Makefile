BIN_DIR=./bin
BIN=azblogfilter
BIN_DEBUG=$(BIN).debug
GCFLAGS_DEBUG="all=-N -l"
SYSTEMD_DIR=~/.config/systemd/user
INSTALL_LOCATION=~/bin

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN)

build_debug:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_DEBUG) -gcflags=$(GCFLAGS_DEBUG)

test:
	go test -v ./...

install: build
	cp $(BIN_DIR)/$(BIN) $(INSTALL_LOCATION)/$(BIN)

install_systemd: install
	mkdir -p $(SYSTEMD_DIR)
