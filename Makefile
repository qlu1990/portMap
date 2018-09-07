.PHONY: all
all : server node proxy

.PHONY: server
server:
	cmd/server/build.sh

.PHONY: node
node:
	cmd/node/build.sh

.PHONY: proxy
proxy:
	cmd/proxy/build.sh
	
.PHONY: proxyclient
proxy:
	cmd/proxyclient/build.sh
clean:
	@echo "remove all file in dir bin"
	rm -rf bin/*