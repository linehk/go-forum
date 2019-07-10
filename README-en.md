# go-forum

[![Build Status](https://travis-ci.org/linehk/go-forum.svg?branch=master)](https://travis-ci.org/linehk/go-forum)
[![codecov](https://codecov.io/gh/linehk/go-forum/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/go-forum)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/go-forum)](https://goreportcard.com/report/github.com/linehk/go-forum)

English | [简体中文](./README.md "简体中文")

go-forum is a simple forum web application that uses the `database/sql` and `http` packages in the standard library, with `toml` as the configuration file format and `MySQL` as the database.

## Installation

```bash
git clone https://github.com/linehk/go-forum.git
```

Create database:

```bash
mysql -uroot -proot < model/setup.sql
```

Fill in the configuration options in `config.toml`.

Then, build it:

```bash
go build -o go-forum
```

And, run it：

```bash
./go-forum
```

## Usages

Visit `http://localhost:8080/`.

## License

[MIT License](./LICENSE "MIT License")