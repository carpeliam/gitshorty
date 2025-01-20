# gitshorty
## Usage
A [Shortcut](https://www.shortcut.com) Command-Line tool that works within git branches that contain an associated story ID as the suffix, e.g. `branchname-sc-123`.

```
NAME:
   sc - command line client for Shortcut

USAGE:
   sc [global options] command [command options]

VERSION:
   0.0.1

COMMANDS:
   browse   open story associated with current branch in web browser
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
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
`gitshorty` requires an [API Token](https://help.shortcut.com/hc/en-us/articles/205701199-Shortcut-API-Tokens), defined as an environment variable:

```sh
SHORTCUT_API_TOKEN="1351e67f-993e-4656-b001-a3842d1354e2"
```
