benchdir = ./bench/
bindir = ./bin/

.PHONY: deps_install
deps_install:
	@echo "\n$(GREEN)Installing Rust...$(RESET)"
	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- --profile minimal -y

	@echo "\n$(GREEN)Installing Go...$(RESET)"
	wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
	tar -C $$HOME -xzf go1.14.linux-amd64.tar.gz
	rm go1.14.linux-amd64.tar.gz

	echo "export PATH=\$$PATH:\$$HOME/go/bin" >> ~/.bashrc
	echo "export GOPATH=\$$HOME/golocal" >> ~/.bashrc
	export PATH=$$PATH:$$HOME/go/bin
	export GOPATH=$$HOME/golocal

	go get github.com/google/btree

	@echo "\n$(GREEN)Done!$(RESET)"

.PHONY: build
build:
	./build

.PHONY: listall
listall:
	@ls -1 $(benchdir)

.PHONY: clean
clean:
	rm -rf $(bindir)*
