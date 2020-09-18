BIN_DIR=./bin
BIN=azblogfilter
BIN_DEBUG=$(BIN).debug
GCFLAGS_DEBUG="all=-N -l"
SYSTEMD_DIR=~/.config/systemd/user
UNIT=azblogfilter-notify.service
UNIT_IN=$(UNIT).in
UNIT_DIR=./systemd
TIMER=azblogfilter-notify.timer
NOTIFY_SCRIPT_DIR=./scripts
NOTIFY_SCRIPT=notify.sh
NOTIFY_SCRIPT_INSTALL_DIR=~
INSTALL_LOCATION=~/bin

.PHONY: build build-debug test install install-systemd clean release install-notify bin-dir

build: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s|LOCAL|$$(git rev-parse --short HEAD)|" ./cmd/version.go; \
		go build -o $(BIN_DIR)/$(BIN); \
		git checkout -- ./cmd/version.go; \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-debug: bin-dir
	sed -i "s|LOCAL|$(git rev-parse --short HEAD)|" ./cmd/version.go
	go build -o $(BIN_DIR)/$(BIN_DEBUG) -gcflags=$(GCFLAGS_DEBUG)

bin-dir:
	mkdir -p $(BIN_DIR)

test:
	go test -v ./...

install: build
	cp $(BIN_DIR)/$(BIN) $(INSTALL_LOCATION)/$(BIN)

install-systemd: install install-notify
	mkdir -p $(SYSTEMD_DIR)
	echo $(INSTALL_LOCATION)/$(BIN)
	sed "s|BIN_PATH|$$(realpath $(INSTALL_LOCATION)/$(BIN))|" $(UNIT_DIR)/$(UNIT_IN) | \
		sed "s|NOTIFY_SCRIPT_PATH|$$(realpath $(NOTIFY_SCRIPT_INSTALL_DIR)/$(NOTIFY_SCRIPT))|" | \
		sed "s|TARGET_USER|$$(whoami)|" \
		> $(SYSTEMD_DIR)/$(UNIT)
	cp $(UNIT_DIR)/$(TIMER) $(SYSTEMD_DIR)/$(TIMER)
	systemctl enable $(TIMER) --user
	systemctl start $(TIMER) --user

install-notify: bin-dir
	cp $(NOTIFY_SCRIPT_DIR)/$(NOTIFY_SCRIPT) $(NOTIFY_SCRIPT_INSTALL_DIR)/$(NOTIFY_SCRIPT)

clean: clean-systemd
	if [ -d $(BIN_DIR) ]; then rm -rf $(BIN_DIR); fi
	if [ -f $(INSTALL_LOCATION)/$(BIN) ]; then rm $(INSTALL_LOCATION)/$(BIN); fi
	if [ -f $(NOTIFY_SCRIPT_INSTALL_DIR)/$(NOTIFY_SCRIPT) ]; then rm $(NOTIFY_SCRIPT_INSTALL_DIR)/$(NOTIFY_SCRIPT); fi

clean-systemd:
	systemctl stop $(UNIT) --user || true
	systemctl disable $(UNIT) --user || true
	systemctl stop $(TIMER) --user || true
	systemctl disable $(TIMER) --user || true
	if [ -f $(SYSTEMD_DIR)/$(UNIT) ]; then rm $(SYSTEMD_DIR)/$(UNIT); fi
	if [ -f $(SYSTEMD_DIR)/$(TIMER) ]; then rm $(SYSTEMD_DIR)/$(TIMER); fi

release: build
	VERSION=$$($(BIN_DIR)/$(BIN) --version); \
	git tag $$VERSION;
