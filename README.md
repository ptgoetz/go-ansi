# Go-ANSI
Go-ANSI is an ANSI color, style, and cursor manipulation library for golang.

[![Go Report Card](https://goreportcard.com/badge/github.com/ptgoetz/go-ansi)](https://goreportcard.com/report/github.com/ptgoetz/go-ansi)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ptgoetz/go-ansi)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/ptgoetz/go-ansi)

## Rationale
There are a number of golang libraries for dealing with ANSI TTY terminals, but most seem to be too narrowly focused.
For example, there are a number of libraries that focus primarily on TTY color, but ignore styles and cursor 
navigation. Others overly abstract away, and block access to, the lower level ANSI code API.

### Existing Golang ANSI Libraries
* [https://github.com/mgutz/ansi](https://github.com/mgutz/ansi) (Color, Style)
* [https://github.com/k0kubun/go-ansi](https://github.com/k0kubun/go-ansi) (Color)
* [https://github.com/ahmetb/go-cursor](https://github.com/ahmetb/go-cursor) (Color, Cursor)

## Proposal
An OSS, liberally licensed (ASL v2), golang library that supports ANSI TTY color, style, and cursor navigation would 
benefit the golang community by providing a (potentially) commonly used utility.

It should provide:
* A low-level, simple API for dealing with low-level things such as ANSI constants, and functions related to color, 
style, and cursor navigation. Higher-level components can use this API to create components or widgets such as progress 
indicators, colorized loggers, etc.
* TTY detection. Provide an API for determining if the output is a an ANSI TTY, and if not, automatically disable ANSI 
formatting.
* Higher-level, reusable components/widgets such as colorized loggers, CLI applications, etc.
* Plenty of examples for developers to learn from and extend.

## Examples
The [examples](examples/) Directory provides some example usage:

### Panels

```shell
 go run examples/panel/panel.go
```

### Progress Indicator

```shell
 go run examples/progress-indicator/progress-indicator.go
```

### Prompts

```shell
 go run examples/prompt/prompt.go
```

### Terminal Screensaver

```shell
 go run examples/terminal-screensaver/terminal-screensaver.go
```

### Terminal Dimensions

```shell
go run examples/unix-info/unix-info.go
```
## Resources
[http://rrbrandt.dee.ufcg.edu.br/en/docs/ansi/cursor](http://rrbrandt.dee.ufcg.edu.br/en/docs/ansi/cursor)
