APP := "showhw"
VERSION := "0.1.0"

alias b := build
alias c := clean
alias r := run
alias v := version
alias cc := cross-compile
alias ghp := github-push

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
    -upx {{APP}}  # Optional

# Update go dependencies
update:
    go get -u -d ./...
    go mod tidy

# Cross compile
cross-compile:
    ./xbuild.sh {{APP}}_{{VERSION}}

# Push and tag the code to Github
github-push: version
    @git push
    @git tag -a {{VERSION}} -m "Version {{VERSION}}"
    @git push origin --tags

# Display the version
version:
    @echo "version {{VERSION}}"
