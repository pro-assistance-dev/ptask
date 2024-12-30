include .env

main := main.go

run:
	go run -ldflags="-X main.ScriptsPath=${SCRIPTS_PATH}"  $(main)

deploy:
	go build  -ldflags="-X main.ScriptsPath=${SCRIPTS_PATH}" && sudo mv ptask /usr/bin
