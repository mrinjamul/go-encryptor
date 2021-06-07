# go-encryptor

[![build status](https://github.com/mrinjamul/go-encryptor/workflows/test/badge.svg)]()
[![build status](https://github.com/mrinjamul/go-encryptor/workflows/release/badge.svg)]()
[![go version](https://img.shields.io/github/go-mod/go-version/mrinjamul/go-encryptor.svg)](https://github.com/mrinjamul/go-encryptor)
[![GoReportCard](https://goreportcard.com/badge/github.com/mrinjamul/go-encryptor)](https://goreportcard.com/report/github.com/mrinjamul/go-encryptor)
[![Code style: standard](https://img.shields.io/badge/code%20style-standard-blue.svg)]()
[![License: Apache 2](https://img.shields.io/badge/License-Apache%202-blue.svg)](https://github.com/mrinjamul/gpassmanager/blob/master/LICENSE)
[![Github all releases](https://img.shields.io/github/downloads/mrinjamul/go-encryptor/total.svg)](https://GitHub.com/mrinjamul/go-encryptor/releases/)

A encryptor to encrypt files using passwords.

## Features

- Can encrypt any files
- Can encrypt folder (**Experimental**, use at own risk)
- Only AES encryption is available right now.

**Tip**: To encrypt a folder, create a tarball of the folder then encrypt it (**Recommanded**)

```sh
tar -cf folder.tar folder
```

## Usage

Encrypt a file,

```sh
go-encryptor encrypt "filename"
```

Decrypt a file,

```sh
go-encryptor decrypt "filename"
```

Use k flag in both `encrypt` & `decrypt` to keep the file.

```sh
go-encryptor encrypt -k "filename"
```

or

```sh
go-encryptor encrypt --keep "filename"
```

## Installing

[ Download ](https://github.com/mrinjamul/go-encryptor/releases) for your platform

## TODO

- Implement others encryption methods

## License

This application is licensed under MIT, Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>
