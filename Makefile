.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@latest -mode=go
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm .

.PHONY: publish
publish:
	wrangler publish
