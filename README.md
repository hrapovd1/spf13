# spf13
CLI App in Go – OSCON 2017

### run tests with lcov coverage files
#### install extension for VSCodium
go get -u github.com/jandelgado/gcov2lcov
#### usage extension
go test -v -coverprofile=coverage.cover && \
gcov2lcov -infile=coverage.cover -outfile=coverage.lcov
