.PHONY: dev
dev:
	go tool air .

.PHONY: generate
generate:
	go tool templ generate

.PHONY: watch
watch:
	go tool templ generate --watch

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.23.1
	tinygo build -o ./build/app.wasm -target wasm -panic=trap -no-debug -gc=leaking -opt=2 ./...

.PHONY: deploy
deploy:
	wrangler deploy

