# go-encryptor

A encryptor to encrypt files using passwords.

## Features

- Only AES encryption is available right now.

## Usage

Encrypt a file,

```sh
go-encryptor encypt "filename"
```

Decrypt a file,

```sh
go-encryptor decypt "filename"
```

Use k flag in both `encrypt` & `decrypt` to keep the file.

```sh
go-encryptor encypt -k "filename"
```

or

```sh
go-encryptor encypt --keep "filename"
```

## Installing

[ Download ](https://github.com/mrinjamul/go-encryptor/releases) for your platform

## TODO

- Implement others encryption methods

## License

This application is licensed under MIT, Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>
