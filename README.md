# gitshorty
A [Shortcut](https://www.shortcut.com) Command-Line tool that works within git branches that contain an associated story ID as the suffix, e.g. `branchname-sc-123`.
## Usage
Usage for individual commands can be found in the [wiki](https://github.com/carpeliam/gitshorty/wiki/Usage).
```
NAME:
   sc - command-line client for Shortcut

USAGE:
   sc [global options] command [command options]

VERSION:
   0.0.3

COMMANDS:
   browse   open story associated with current branch in web browser
   clean    delete local branches associated with delivered stories
   tasks    display tasks associated with the current branch's story
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api-token value  API Token to use with Shortcut API [$SHORTCUT_API_TOKEN]
   --verbose, -V      enable verbose logging (default: false)
   --help, -h         show help
   --version, -v      print the version
```

## Installation
### Via Homebrew
```
brew install carpeliam/brew/gitshorty
```

### From source
The source is at least compatible with Go v1.23.1, and can be built as follows:
```
go build -o sc .
```

## Configuration
`gitshorty` requires an [API Token](https://help.shortcut.com/hc/en-us/articles/205701199-Shortcut-API-Tokens). You can either define it as an environment variable or as a global option for `sc`:

```sh
# As a global option
sc --api-token 1351e67f-993e-4656-b001-a3842d1354e2 browse
# As an environment variable
export SHORTCUT_API_TOKEN="1351e67f-993e-4656-b001-a3842d1354e2"
sc browse
```
