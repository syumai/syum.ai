.PHONY: dev
dev:
	wrangler dev

.PHONY: generate
generate:
	templ generate

.PHONY: watch
watch:
	templ generate --watch

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@v0.23.1
	tinygo build -o ./build/app.wasm -target wasm -no-debug ./...

.PHONY: deploy
deploy:
	wrangler deploy

.PHONY: install-tools
install-tools:
	go install github.com/a-h/templ/cmd/templ@latest
