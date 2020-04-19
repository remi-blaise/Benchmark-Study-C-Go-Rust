benchdir = ./bench/
bindir = ./bin/

RUST_OPTI=--release
CPP_OPTI=-O2

RUST_DEBUG=
CPP_DEBUG=-Wall

RUST=cargo build $(RUST_OPTI)
GO=go build
CPP=g++ $(CPP_OPTI) $(CPP_DEBUG) -std=c++11 -pthread -I bench/serialization

GREEN=\e[1m\e[32m
RESET=\e[0m

deps_install:
	@echo "\n$(GREEN)Installing Rust...$(RESET)"
	curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- --profile minimal -y

	@echo "\n$(GREEN)Installing Go...$(RESET)"
	wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
	tar -C $$HOME -xzf go1.14.linux-amd64.tar.gz
	rm go1.14.linux-amd64.tar.gz
	@echo "export PATH=\$$PATH:\$$HOME/go/bin" >> ~/.profile

	@echo "\n$(GREEN)Done!$(RESET)"

listall:
	@ls -1 $(benchdir)

build:
	cargo build --release
	mkdir -p $(bindir)$(test)/rust
	rm -f $(bindir)$(test)/rust/*
	for f in $(wildcard $(benchdir)$(test)/*.rs); do \
		filename=`echo $$f | sed 's:.*/::' | sed 's:.rs::'`; \
		cp target/release/$(test)-$$filename $(bindir)$(test)/rust/$$filename; \
	done
	@echo "$(GREEN)Built Rust!$(RESET)\n"
	mkdir -p $(bindir)$(test)/go
	for f in $(wildcard $(benchdir)$(test)/*.go) ; do \
		$(GO) -o $(bindir)$(test)/go $$f ; \
	done
	@echo "$(GREEN)Built Go!$(RESET)\n"
	mkdir -p $(bindir)$(test)/cpp
	for f in $(wildcard $(benchdir)$(test)/*.cpp) ; do \
		OUT=$$(basename $$f | sed -E "s/(.*).cpp/\1/"); \
		$(CPP) -o $(bindir)$(test)/cpp/$$OUT $$f ; \
	done
	@echo "$(GREEN)Built C++!$(RESET)\n"

clean:
	rm -rf $(bindir)*
