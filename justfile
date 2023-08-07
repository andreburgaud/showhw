APP := "showhw"

alias b := build
alias c := clean
alias r := run
alias cc := cross-compile

# Default recipe (this list)
default:
    @just --list

# Delete generated binaries
clean:
    rm {{APP}}

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

# Cross compile
cross-compile:
    ./xbuild.sh {{APP}}
