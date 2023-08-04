APP := "showhw"

# Default recipe (this list)
default:
    @just --list

# Delete generated binaries
clean:
    @echo todo

# Run local app
run:
    go run .

# Build app
build:
    go build .

# Build release version
release:
    go build -ldflags="-s -w" .
    upx {{APP}}

# Update go dependencies
update:
    go get -u -d ./...
    go mod tidy

# # Install fyne CLI
# install-fyne:
#     go install fyne.io/fyne/v2/cmd/fyne@latest

