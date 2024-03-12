# go-forum

[![build](https://github.com/linehk/go-forum/actions/workflows/build.yml/badge.svg "build")](https://github.com/linehk/go-forum/actions)
[![codecov](https://codecov.io/gh/linehk/go-forum/graph/badge.svg "codecov")](https://codecov.io/gh/linehk/go-forum)
[![go report](https://goreportcard.com/badge/github.com/linehk/go-forum "go report")](https://goreportcard.com/report/github.com/linehk/go-forum)

English | [简体中文](./README-zh.md "简体中文")

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

## Contributing

If you feel that there is something to improve this project, please feel free to launch Pull Request.

For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT License](./LICENSE "MIT License")