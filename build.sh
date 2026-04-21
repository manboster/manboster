# TODO: it will be used in cd/cis
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o manboster_0.0.1_windows_amd64.exe ./cmd/manboster/main.go
GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o manboster_0.0.1_windows_arm64.exe ./cmd/manboster/main.go
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o manboster_0.0.1_darwin_arm64 ./cmd/manboster/main.go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o manboster_0.0.1_linux_amd64 ./cmd/manboster/main.go
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o manboster_0.0.1_linux_arm64 ./cmd/manboster/main.go