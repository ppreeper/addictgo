dev:
	air

swagger:
	swag init --dir ./,./api


buildtemp:
	swag init --dir ./,./api
	go build -o ./tmp/main .

