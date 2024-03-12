# go-forum

[![build](https://github.com/linehk/go-forum/actions/workflows/build.yml/badge.svg "build")](https://github.com/linehk/go-forum/actions)
[![codecov](https://codecov.io/gh/linehk/go-forum/graph/badge.svg "codecov")](https://codecov.io/gh/linehk/go-forum)
[![go report](https://goreportcard.com/badge/github.com/linehk/go-forum "go report")](https://goreportcard.com/report/github.com/linehk/go-forum)

[English](./README.md "English") | 简体中文

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

## 参与贡献

如果您觉得这个项目有什么需要改进的地方，欢迎发起 Pull Request。

如果有重大变化，请先打开一个 Issue，讨论您想要改变的内容。

## 开源许可证

[MIT License](./LICENSE "MIT License")