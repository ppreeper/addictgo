dev:
	air

swagger:
	swag init --dir ./../api


buildtemp:
	swag init
	go build -o ./tmp/main .
