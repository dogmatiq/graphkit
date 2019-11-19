# Dogma Graph Toolkit

[![Build Status](https://github.com/dogmatiq/graphkit/workflows/CI/badge.svg)](https://github.com/dogmatiq/graphkit/actions?workflow=CI)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/graphkit/master.svg)](https://codecov.io/github/dogmatiq/graphkit)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/graphkit.svg?label=semver)](https://semver.org)
[![GoDoc](https://godoc.org/github.com/dogmatiq/graphkit?status.svg)](https://godoc.org/github.com/dogmatiq/graphkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/dogmatiq/graphkit)](https://goreportcard.com/report/github.com/dogmatiq/graphkit)

Generate message-flow diagrams for [Dogma](https://github.com/dogmatiq/dogma) applications.

This library produces [Graphviz](https://www.graphviz.org/) diagrams in the [DOT
language](https://www.graphviz.org/doc/info/lang.html). Use one of the Graphviz
renderers to produce a graphic.

For an example, see the [`graph` binary in the example
application](https://github.com/dogmatiq/example/blob/master/cmd/graph/main.go),
and the associated recipe in the [Makefile](https://github.com/dogmatiq/example/blob/master/Makefile).
