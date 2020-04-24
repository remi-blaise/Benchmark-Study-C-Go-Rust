benchdir = ./bench/
bindir = ./bin/

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

clean:
	rm -rf $(bindir)*
