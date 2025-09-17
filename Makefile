.PHONY: run build

run:
	go run main.go

build:
	go build -o PersonalTaskScheduler main.go
