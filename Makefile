
install_path = /usr/local/bin
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

install: all
	sudo cp $(dist)/$(binaryName) $(install_path)/$(binaryName)
	sudo chmod +x $(install_path)/$(binaryName)
	sudo cp assets/dev.raccoon-tw.squaerere.plist /Library/LaunchAgents/dev.raccoon-tw.squaerere.plist
	sudo launchctl load /Library/LaunchAgents/dev.raccoon-tw.squaerere.plist

uninstall:
	@sudo launchctl unload /Library/LaunchAgents/dev.raccoon-tw.squaerere.plist
	sudo rm -r /Library/LaunchAgents/dev.raccoon-tw.squaerere.plist
	sudo rm -r $(install_path)/$(binaryName)