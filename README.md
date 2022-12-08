<div align="center">

<img src="https://user-images.githubusercontent.com/30717818/206506472-340a3497-207f-45a9-9bb7-efb57c4274bc.svg" width="100">

# dotenv

</div>

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/dotenv.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/dotenv/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/dotenv/actions/workflows/go.yml) [![CodeQL](https://github.com/gowizzard/dotenv/actions/workflows/codeql.yml/badge.svg)](https://github.com/gowizzard/dotenv/actions/workflows/codeql.yml) [![CompVer](https://github.com/gowizzard/dotenv/actions/workflows/compver.yml/badge.svg)](https://github.com/gowizzard/dotenv/actions/workflows/compver.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/dotenv.svg)](https://pkg.go.dev/github.com/gowizzard/dotenv) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/dotenv)](https://goreportcard.com/report/github.com/gowizzard/dotenv) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/dotenv)](https://github.com/gowizzard/dotenv/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/dotenv)](https://github.com/gowizzard/dotenv/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/dotenv)](https://github.com/gowizzard/dotenv/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/dotenv)](https://github.com/gowizzard/dotenv/blob/master/LICENSE)

With this lightweight library you can read a local file with environment variables into your project. Additionally, you can use functions to read the data and the variable will be returned directly in the desired type.

## Installation

First you have to install the package. You can do this as follows:

```shell
go get github.com/gowizzard/dotenv
```

## How to use

## Import variables

If you want to read your local `.env` file, so you can use these variables in your project, you can use this function for that.

With this function the data will be loaded from the file and set as local variables. After that you can read them with standard functions, or you can use the following functions.

[Here](https://regex101.com/r/SEDjKj/1) you can find the regex expression which is used to read the environment variables.

```go
err := dotenv.Import(value.path)
if err != nil {
    panic(err)
}
```

## Read variables

If you don't want to work with the standard golang function `func Getenv(key string) string` from the `os` package, because you want to output the variables directly in a certain type, then you can use the following different functions.

The functions check directly if the desired variable is available and returns the value. If the value is not available, then the default value of the type is returned.

You can also use these functions without the import function if the variables are already available locally.

### Boolean

With this function you can read an environment variable and return it directly as `type bool`.

```go
result := dotenv.Boolean("KEY")
fmt.Println(result)
```

### Float

With this function you can read an environment variable and return it directly as `type float64`. In this function you must not only specify the desired key, but also the bit size of the float type.

```go
result := dotenv.Float("KEY", 64)
fmt.Println(result)
```

### Integer

With this function you can read an environment variable and return it directly as `type integer`.

```go
result := dotenv.Integer("KEY")
fmt.Println(result)
```

### String

With this function you can read an environment variable and return it directly as `type string`.

```go
result := dotenv.String("KEY")
fmt.Println(result)
```

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).
