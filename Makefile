run: build
	@./bin/nextgen

build:
	templ generate
	@go build -o ./bin/nextgen ./cmd/main.go

css:
	tailwindcss -i view/css/app.css -o public/styles.css --watch

