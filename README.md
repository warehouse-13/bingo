## bingo

A CLI binary bootstrapper for go. (`bin`-`go`, geddit?).

This is a pattern I tend to use a lot, so I am making it easier for myself to
get started.

It sets up a new CLI boilerplate, based on [`urfave/cli/v2`](https://github.com/urfave/cli).

After removing some placeholders, it should work out of the box.

### Installation

Get a [released binary](https://github.com/warehouse-13/bingo/releases), or

```bash
go install github.com/warehouse-13/bingo@latest
```

### Usage

```bash
mkdir new-cli
cd new-cli

git init
go mod init github.com/some-user/new-cli

bingo bootstrap --name woohoo
```

Output:
```
Bootstrapping CLI for `woohoo`...

Created directory pkg/command/
Created directory pkg/flags/
Created directory pkg/config/
Created directory pkg/version/
Written file main.go
Written file pkg/command/app.go
Written file pkg/command/example.go
Written file pkg/command/version.go
Written file pkg/flags/flags.go
Written file pkg/config/config.go
Written file pkg/version/version.go


Done!
Check each file and follow instructions within.
Run `go mod tidy` and start using.
```
