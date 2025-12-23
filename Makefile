build:
	GOOS=windows GOARCH=amd64 go build -o dist/mine-sync-win src/main.go
	go build -o dist/mine-sync-linux src/main.go
	chmod +X dist/mine-sync-linux