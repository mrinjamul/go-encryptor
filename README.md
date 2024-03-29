# go-encryptor

[![build status](https://github.com/mrinjamul/go-encryptor/workflows/test/badge.svg)]()
[![build status](https://github.com/mrinjamul/go-encryptor/workflows/release/badge.svg)]()
[![go version](https://img.shields.io/github/go-mod/go-version/mrinjamul/go-encryptor.svg)](https://github.com/mrinjamul/go-encryptor)
[![GoReportCard](https://goreportcard.com/badge/github.com/mrinjamul/go-encryptor)](https://goreportcard.com/report/github.com/mrinjamul/go-encryptor)
[![Code style: standard](https://img.shields.io/badge/code%20style-standard-blue.svg)]()
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/mrinjamul/gpassmanager/blob/master/LICENSE)
[![Github all releases](https://img.shields.io/github/downloads/mrinjamul/go-encryptor/total.svg)](https://GitHub.com/mrinjamul/go-encryptor/releases/)

## Overview

a command-line application written in Go that allows for encrypting and decrypting data using AES 256 (GCM) and ChaCha20 poly 1305 encryption algorithms. It provides a simple and straightforward way to protect sensitive data and ensure that it can only be accessed by authorized users.

## Features

- Provides a simple and easy-to-use command-line interface
- Can encrypt any file or folder
- AES-256 encryption
- Password based encryption
- XChaCha20-Poly1305 encryption

**Tip**: To encrypt folders, create a tarball of the folder then encrypt it (**Recommanded**)

```sh
tar -cf folder.tar folder ...
```

## Build

To build the application, you need to have Go installed on your machine.

```sh
go build
```

## Usage

The following is an example of how to use go-encryptor to encrypt data using AES 256 (GCM):

```sh
go-encryptor encrypt -p "256Key-32Characters1234567890" -m "aes" "filename"
```

And the following is an example of how to use go-encryptor to encrypt data using ChaCha20-Poly1305:

```sh
go-encryptor encrypt -p "256Key-32Characters1234567890" -m "chacha20" "filename"
```

For decryption, the usage is similar, with the only difference being the use of the decrypt subcommand instead of encrypt.

For a full list of options and usage instructions, run `go-encryptor --help`

Use k flag in both `encrypt` & `decrypt` to keep the file.

```sh
go-encryptor encrypt -k "filename"
```

or

```sh
go-encryptor encrypt --keep "filename"
```

## More

Encrypt a file with a specific encryption method using a password,

```sh
go-encryptor encrypt --method "aes" --password "password" "filename"
```

or

```sh
go-encryptor encrypt -m "aes" -p "password" "filename"
```

Decrypt a file and print the output to stdout and pipe it to another command,

```sh
go-encryptor decrypt -p "password" --print "filename" | [command]
```

```

    go-encryptor: A file encryptor.
    go-encryptor is created to be as simple as possible to help you
    encrypt and decrypt files.

    Usage:
    go-encryptor [command]

    Available Commands:
    decrypt     Decrypt encrypted file
    encrypt     Encrypt file or folder
    help        Help about any command
    version     Prints version

    Flags:
    -h, --help   help for go-encryptor

    Use "go-encryptor [command] --help" for more information about a command.

```

## Installing

[ Download ](https://github.com/mrinjamul/go-encryptor/releases) for your platform

or Install from snap

```sh
sudo snap install go-encryptor
```

## Benchmarks

For AES-256 encryption (`time go-encryptor en -m "aes" -p "Password" alpine.iso`) ,

```
alpine.iso encrypted successfully.

________________________________________________________
Executed in    3.42 secs    fish           external
   usr time    3.19 secs  277.00 micros    3.19 secs
   sys time    0.32 secs   99.00 micros    0.32 secs

alpine.iso decrypted successfully.

________________________________________________________
Executed in    3.40 secs    fish           external
   usr time    3.18 secs  347.00 micros    3.18 secs
   sys time    0.25 secs  126.00 micros    0.25 secs

```

For XChaCha20-Poly1305 encryption (`time go-encryptor en -m "xchacha" -p "Password" alpine.iso`) ,

```
alpine.iso encrypted successfully.

________________________________________________________
Executed in  802.34 millis    fish           external
   usr time  898.45 millis  414.00 micros  898.04 millis
   sys time  296.73 millis  135.00 micros  296.60 millis


alpine.iso decrypted successfully.

________________________________________________________
Executed in  755.63 millis    fish           external
   usr time  804.23 millis    0.00 micros  804.23 millis
   sys time  365.06 millis  584.00 micros  364.48 millis

```

## License

This application is licensed under MIT, Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>

## Troubleshooting

If you encounter any errors while using go-encryptor, make sure that you have the correct key, algorithm and that the input and output files are correctly specified. If the issue persists, please file an issue on the GitHub repository for the project.

## Conclusion

go-encryptor is a simple and easy-to-use command-line application for encrypting and decrypting data using AES 256 (GCM) and ChaCha20-Poly1305 encryption algorithms. Its straightforward configuration and user-friendly interface make it a great choice for protecting sensitive information.
