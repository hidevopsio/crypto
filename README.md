# crypto
crypto is a command line tool that encrypt / decrypt text

## Installation

Type below command to install `crypto`

```bash
go get -u github.com/hidevopsio/crypto
go install github.com/hidevopsio/crypto
```

## How to use

```bash
Usage:
  crypto [flags]

Examples:

crypto rsa -h
crypto rsa -e -s "text to encrypt"
crypto rsa -d -s "text to decrypt"


Flags:
  -d, --decrypt         run with option --decrypt or -d for text encryption
  -e, --encrypt         run with option --encrypt or -e for text encryption
  -h, --help            help for crypto
  -k, --key string      run with option --key or -k for rsa key
  -s, --source string   run with option --source=source text to encrypt or encrypt
```

## Encrypt text

```bash
crypto rsa -e -s hello

HZ62JK0istPirbTe4M14foFdE88hvB1oH5HruOEVwwAGU7tT2fy39grAnFriZFJVgM0yAgoqgppW3d40IuUBZMFajzIPzyhM9s2vOWNmvTYFkXF0vYfQporA/FGVttOfD27Tji3XAnoc1nlOAPjRLnF1vGbSLDUWkg3/LsCcwGY=
```

## Decrypt text

```bash
crypto rsa -d -s HZ62JK0istPirbTe4M14foFdE88hvB1oH5HruOEVwwAGU7tT2fy39grAnFriZFJVgM0yAgoqgppW3d40IuUBZMFajzIPzyhM9s2vOWNmvTYFkXF0vYfQporA/FGVttOfD27Tji3XAnoc1nlOAPjRLnF1vGbSLDUWkg3/LsCcwGY=

hello
```