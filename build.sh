# TODO: it will be used in cd/cis

GIT_COMMIT=$(git rev-parse HEAD | cut -c1-6)
VERSION="0.1.0"
BUILD_TIME=$(date "+%Y-%m-%d_%H:%M:%S")

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X 'github.com/manboster/manboster/internal/config.BuildCommit=${GIT_COMMIT}' -X 'github.com/manboster/manboster/internal/config.BuildTime=${BUILD_TIME}'" -trimpath -o manboster_${VERSION}_windows_amd64.exe ./cmd/manboster/main.go
GOOS=windows GOARCH=arm64 go build -ldflags="-s -w -X 'github.com/manboster/manboster/internal/config.BuildCommit=${GIT_COMMIT}' -X 'github.com/manboster/manboster/internal/config.BuildTime=${BUILD_TIME}'" -trimpath -o manboster_${VERSION}_windows_arm64.exe ./cmd/manboster/main.go
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X 'github.com/manboster/manboster/internal/config.BuildCommit=${GIT_COMMIT}' -X 'github.com/manboster/manboster/internal/config.BuildTime=${BUILD_TIME}'" -trimpath -o manboster_${VERSION}_darwin_arm64 ./cmd/manboster/main.go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'github.com/manboster/manboster/internal/config.BuildCommit=${GIT_COMMIT}' -X 'github.com/manboster/manboster/internal/config.BuildTime=${BUILD_TIME}'" -trimpath -o manboster_${VERSION}_linux_amd64 ./cmd/manboster/main.go
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X 'github.com/manboster/manboster/internal/config.BuildCommit=${GIT_COMMIT}' -X 'github.com/manboster/manboster/internal/config.BuildTime=${BUILD_TIME}'" -trimpath -o manboster_${VERSION}_linux_arm64 ./cmd/manboster/main.go