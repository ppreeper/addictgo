.PHONY: htmx
htmx:
	@wget -qO ./internal/router/static/js/htmx.min.js "https://unpkg.com/htmx.org@latest/dist/htmx.min.js"

.PHONY: alpinejs
alpinejs:
	@wget -qO ./internal/router/static/js/alpine.min.js "https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"

.PHONY: air
clean:
	@find . -type f -iname "*_templ.go" -exec rm -f "{}" \;

## air templ tailwind commands when run together they create a dynamic dev environment
.PHONY: air
air:
	@air

.PHONY: templ-generate
templ-generate:
	@templ generate

.PHONY: templ-watch
templ-watch:
	@templ generate --watch
# --proxy=http://localhost:3001

.PHONY: tailwind-watch
tailwind-watch:
	@tailwindcss -c ./tailwind.config.js -i ./web/static/css/input.css -o ./internal/router/static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -c ./tailwind.config.js -i ./web/static/css/input.css -o ./internal/router/static/css/style.css --minify

.PHONY: tailwind-swagger
swagger:
	@swag init --dir ./cmd/addict,./internal/api

.PHONY: dev
dev:
	make -j 3 air templ tailwind
# go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build && make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go