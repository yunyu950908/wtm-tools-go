.PHONY: build

dev: lint codegen
	go run cmd/main.go

build: lint codegen
	CGO_ENABLED=0 go build -trimpath -ldflags -extldflags=-static -o bin/wtm-tools-go cmd/main.go

lint:
	golangci-lint run

codegen:
	rm -rf internal/sc
	mkdir -p internal/sc
	abigen --abi assets/abis/erc20.abi --pkg sc --type ERC20 --out internal/sc/erc20.go
	abigen --abi assets/abis/erc4626.abi --pkg sc --type ERC4626 --out internal/sc/erc4626.go
	abigen --abi assets/abis/aggregatorV3.abi --pkg sc --type AggregatorV3 --out internal/sc/aggregatorV3.go
	abigen --abi assets/abis/infraredVault.abi --pkg sc --type Vault --out internal/sc/infraredVault.go
