# Chương 2: Cài đặt và Môi trường

## Cài đặt Go

### Windows

1. Tải Go từ [golang.org/dl](https://go.dev/dl/)
2. Chạy installer (.msi file)
3. Go sẽ tự động thêm vào PATH
4. Kiểm tra cài đặt:
```bash
go version
```

### macOS

#### Option 1: Homebrew (Khuyến nghị)
```bash
brew install go
```

#### Option 2: Tải từ website
1. Tải từ [golang.org/dl](https://go.dev/dl/)
2. Chạy installer (.pkg file)
3. Kiểm tra: `go version`

### Linux

#### Ubuntu/Debian
```bash
sudo apt update
sudo apt install golang-go
```

#### Fedora
```bash
sudo dnf install golang
```

#### Từ source (mới nhất)
```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

## Kiểm tra cài đặt

```bash
# Kiểm tra version
go version

# Xem thông tin Go environment
go env

# Xem các lệnh Go có sẵn
go help
```

## Cấu trúc Workspace Go

Go sử dụng workspace với cấu trúc:

```
go-workspace/
├── bin/          # Compiled binaries
├── pkg/          # Compiled packages
└── src/          # Source code
    └── github.com/
        └── username/
            └── project/
```

### GOPATH (Go 1.11 trước)

```bash
# Set GOPATH
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### Go Modules (Go 1.11+, Khuyến nghị)

Không cần GOPATH, mỗi project là một module:

```bash
# Tạo module mới
go mod init github.com/username/project

# Tải dependencies
go mod tidy

# Xem dependencies
go mod graph
```

## IDE và Editor

### 1. VS Code (Khuyến nghị)
- Cài extension: Go (by Go Team)
- Tự động format, lint, test

### 2. GoLand (JetBrains)
- IDE chuyên dụng cho Go
- Có bản trial và paid

### 3. Vim/Neovim
- Cài plugin: vim-go

### 4. Go Playground
- [play.golang.org](https://go.dev/play/)
- Chạy code online không cần cài đặt

## Các lệnh Go cơ bản

```bash
# Chạy Go program
go run main.go

# Build executable
go build main.go
./main  # hoặc main.exe trên Windows

# Build và chạy
go run .

# Format code
go fmt ./...

# Lint code
golangci-lint run

# Chạy tests
go test ./...

# Tải dependencies
go get package-name

# Xem documentation
go doc package-name
```

## Tạo project đầu tiên

```bash
# Tạo thư mục project
mkdir my-go-app
cd my-go-app

# Khởi tạo Go module
go mod init my-go-app

# Tạo file main.go
cat > main.go << EOF
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
EOF

# Chạy
go run main.go
```

## Go Tools

### gofmt
Format code tự động:
```bash
gofmt -w .
```

### go vet
Tìm lỗi tiềm ẩn:
```bash
go vet ./...
```

### golangci-lint
Linter mạnh mẽ:
```bash
# Cài đặt
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Chạy
golangci-lint run
```

## Environment Variables

```bash
# GOPATH (nếu dùng Go < 1.11)
export GOPATH=$HOME/go

# GOROOT (thường không cần set)
export GOROOT=/usr/local/go

# GOBIN
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

## Tóm tắt

- Go có thể cài đặt trên mọi hệ điều hành
- Go Modules là cách quản lý dependencies hiện đại
- VS Code với Go extension là lựa chọn tốt
- Các lệnh Go đơn giản và mạnh mẽ

**Tiếp theo**: [Chương 3: Cú pháp cơ bản](./03-basic-syntax.md)

