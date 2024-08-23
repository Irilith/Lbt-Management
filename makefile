MAIN = ./cmd/Lbt-Management/main
OUT = ./bin/LbtM.exe
all: build

build:
	go build -o $(OUT) $(MAIN).go
