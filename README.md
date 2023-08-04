# Reusable baseline CI workflows for Go [![CI](https://github.com/boinkor-net/ci-baseline-go/actions/workflows/ci.yml/badge.svg)](https://github.com/boinkor-net/ci-baseline-go/actions/workflows/ci.yml)

The workflows in this repo all help drive the CI in this organization's Golang repos. There are a few workflows defined here that are [reusable](https://docs.github.com/en/actions/using-workflows/reusing-workflows) in github actions.

All these workflows expect to test a modules-enabled go repo and take some common arguments:

* `module_base`: The directory containing the code base under test. Defaults to `.`
* `go_version`: The version of go under test. Defaults to `1.20`.

## `ci_baseline_go_tests.yml` - Tests

This workflow runs `go test` against a module.

Example workflow:

```yml
jobs:
  tests:
    uses: "boinkor-net/ci-baseline-go/.github/workflows/ci_baseline_go_tests.yml@main"
    with:
      module_base: "cmd/hello-world"
```

## `ci_baseline_go_build.yml` - Build

This workflow runs `go build` against a module. It has options:

* `subdir`: The directory containing the main.go file resulting in the executable. Defaults to `.`

Example workflow:

```yml
jobs:
  build:
    uses: "./.github/workflows/ci_baseline_go_build.yml"
    with:
      module_base: "go/"
      subdir: cmd/hello-world
```

## `ci_baseline_go_lints.yml` - Lints & formatting

This workflow runs `golangci-lint` and `goimports` against a module, as two separate / parallel jobs.

Example workflow:

```yml
  success_lints:
    uses: "./.github/workflows/ci_baseline_go_lints.yml"
    with:
      module_base: "cmd/hello-world"
```
