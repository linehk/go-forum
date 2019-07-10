# go-forum

[![Build Status](https://travis-ci.org/linehk/go-forum.svg?branch=master)](https://travis-ci.org/linehk/go-forum)
[![codecov](https://codecov.io/gh/linehk/go-forum/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/go-forum)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/go-forum)](https://goreportcard.com/report/github.com/linehk/go-forum)

[English](./README-en.md "English") | 简体中文

go-forum 是一个简单的论坛 Web 程序，主要使用标准库中的 `database/sql` 和 `http` 包实现，以 `toml` 为配置文件格式和以 `MySQL` 为数据库。

## 安装

```bash
git clone https://github.com/linehk/go-forum.git
```

建立数据库：

```bash
mysql -uroot -proot < model/setup.sql
```

在 `config.toml` 中填写配置选项。

然后进行编译：

```bash
go build -o go-forum
```

再运行：

```bash
./go-forum
```

## 使用

访问 `http://localhost:8080/`。

## 开源许可证

[MIT License](./LICENSE "MIT License")