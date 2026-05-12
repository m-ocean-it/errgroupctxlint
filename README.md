# errgroupctxlint


## About

This linter catches cases when, within an error-group goroutine, a non-errgroup context is passed to a function or method, while there is a context, specifically attached to the said errgroup. In most cases, you want that specific context to be passed to functions invoked within the errgroup's `Go` methods.

```go
eg, egCtx := errgroup.WithContext(ctx)

eg.Go(func() error {
	return doSmth(ctx) // want `errgroup callback should probably not reference outer context "ctx", use the errgroup-derived context "egCtx"`
})

eg.Go(func() error {
	return doSmth(egCtx) // Correctly uses the context returned by "errgroup.WithContext"
})

eg.TryGo(func() error {
	return doSmth(ctx) // want `errgroup callback should probably not reference outer context "ctx", use the errgroup-derived context "egCtx"`
})
```

A *lot* more cases are covered in the [`examples.go`](analyzer/testdata/base/examples.go) file!


## Installation
```sh
go install 'github.com/m-ocean-it/errgroupctxlint/cmd/errgroupctxlint@latest'
```


## Usage
```sh
errgroupctxlint ./...      # lint the whole project
errgroupctxlint <package>  # or specific packages
```

You can specify alternative `errgroup`-packages separated with commas (useful in case you project uses some fork(s)/wrapper(s)/alternative(s) to the original `errgroup` package):
```sh
errgroupctxlint -pkgs 'golang.org/x/sync/errgroup,github.com/johejo/semerrgroup,some.org/platform/errgroup/v2' ./...
```


## [Golangci-lint](https://github.com/golangci/golangci-lint) plugin guide

As of now (2026-05-12), the linter is not integrated into `golangci-lint`, however, you can build a custom version of `golangci-lint` with `errgroupctxlint` supported as a plugin. If you're interested in that, read below.

Read the [official guide](https://golangci-lint.run/docs/plugins/module-plugins/).

Prepare a `.custom-gcl.yml` file:
```yml
version: v2.11.4
plugins:
  - module: "github.com/m-ocean-it/errgroupctxlint"
    import: "github.com/m-ocean-it/errgroupctxlint"
    version: latest
```

Run
```sh
golangci-lint custom -v
```

A custom binary of `golangci-lint` would appear at path `./custom-gcl`.

Prepare a `.golangci.yml` config file (or amend the existing one):
```yml
version: "2"

linters:
  default: none
  enable:
    - errgroupctxlint
  settings:
    custom:
      errgroupctxlint:
        type: "module"
        settings:
          pkgs:
            # Specify alternative errgroup packages here, if needed, like so:
            # - golang.org/x/sync/errgroup
            # - errgroup1
            # - foobar/errgroup2
```

Run the resulted binary like the original `golangci-lint`:
```sh
./custom-gcl run ./...
```
