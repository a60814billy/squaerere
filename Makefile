
dist = ./dist
binaryName = squaerere

mainSrcFile = ./cmd/main.go

.PHONY: all clean run

all: $(mainSrcFile)
	go build -o $(dist)/$(binaryName) $(mainSrcFile)

clean:
	rm -rf $(dist)/*

run: all
	sudo $(dist)/$(binaryName)