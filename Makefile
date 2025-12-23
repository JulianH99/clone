##
# Project Title
#
# @file
# @version 0.1

run:
	@go run main.go


build:
	@go build -o build/clone .

test:
	@go test ./...

# end
