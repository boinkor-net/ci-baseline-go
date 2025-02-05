# Reusable baseline CI actions for Go [![CI](https://github.com/boinkor-net/ci-baseline-go/actions/workflows/ci.yml/badge.svg)](https://github.com/boinkor-net/ci-baseline-go/actions/workflows/ci.yml)

The actions in this repo all help drive the CI in this organization's Golang repos.

All these actions expect to test a modules-enabled go repo and take some common arguments:

* `module_base`: The directory containing the code base under test. Defaults to `.`
* `go_version`: The version of go under test. If unset, this will use the version given in go.mod under `module_base`.

## `actions/test` - Tests

This action runs `go test` against a module.

Example workflow:

```yml
jobs:
  success_tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./actions/test
        with:
          module_base: "tests/success"
```

## `actions/build` - Build

This action runs `go build` against a module. It has options:

* `subdir`: The directory containing the main.go file resulting in the executable. Defaults to `.`

Example workflow:

```yml
jobs:
  success_build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./actions/build
        with:
          module_base: "tests/success"
          subdir: "./cmd/hello-world"
```

## `ci_baseline_go_lints.yml` - Lints & formatting

This action runs `golangci-lint` and `goimports` against a module, as two separate / parallel jobs.

Example workflow:

```yml
jobs:
  success_lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./actions/lint
        with:
          module_base: "tests/success"
```
