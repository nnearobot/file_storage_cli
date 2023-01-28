# FILESTORAGE CLI
A CLI (command line interface) for interacting with a file storage that implements a functions listed below:

1. List all the files on the filestorage server
2. Upload a file to this server
3. Delete a file by providing its name


## Usage
Int the **./out** directory, there are several binaries that are built for different OSs and architecture.

As part of the testing this assignment, you can directly use one of these files, corresponding to the architecture of your computer.


### File listing
```
./fstorage list
```

### File uploading
```
./fstorage upload ~/files/testfile.txt
```

### File deleting
```
./fstorage delete testfile.txt
```

### Help info
```
./fstorage -h
```


## Developing points
There is an open-source Go-module to build powerful CLI, [Cobra](https://github.com/spf13/cobra). It has great features built-in, which we can use to create clean, user-friendly interfaces. That is why this CLI is created on Golang using Cobra module.



## Building the application
Go has an incredible build system for creating a binary executable files for multiple OS and architectures.

### Linux target build
```
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o out/fstorage_linux-arm64-calc -ldflags="-extldflags=-static" # linux, arm64 arch
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o out/fstorage_linux-amd64-calc -ldflags="-extldflags=-static" # linux, amd64 arch
```

### MacOS (Darwin) target build
```
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o out/fstorage_darwin-arm64-calc -ldflags="-extldflags=-static" # mac, arm64 arch
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o out/fstorage_darwin-amd64-calc -ldflags="-extldflags=-static" # mac, amd64 arch
```

### Windows target build
```
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o out/fstorage_windows-arm64-calc -ldflags="-extldflags=-static" # windows, arm64 arch
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o out/fstorage_windows-amd64-calc -ldflags="-extldflags=-static" # windows, amd64 arch
```


## Distribution

For public distribution, we should place it in the repositories of such package managers as, Homebrew (MacOS), Snapcraft (Linux), Scoop (Windows) and other.

Also there is a [GoReleaser](https://github.com/goreleaser/goreleaser) project that makes all the build and release process automatically.
