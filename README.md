# imgrep

`imgrep` is a command-line utility in Go to search for keywords found
within images.

![coverage](https://img.shields.io/badge/coverage-70%25-green.svg)

## Installation

`imgrep` depends on
[Tesseract](https://github.com/tesseract-ocr/tesseract).
  * On Fedora: `sudo dnf install tesseract-devel`
  * On Debian: `sudo apt-get install libtesseract-dev`
  * On macOS: `brew install tesseract`

### Prerequisites
#### Linux: 

```
# Fedora
sudo dnf install tesseract tesseract-devel leptonica-devel golang

# Debian
sudo apt-get install tesseract-ocr libtesseract-dev libleptonica-dev golang
```

#### macOS:

 1. Install [Go](https://golang.org/dl/)
 2. Install [homebrew](https://brew.sh)

```
brew install tesseract
```

### Get `imgrep`
Make sure your `$GOPATH` is set, then run:

```
# fetch src and install binary
go get github.com/keeferrourke/imgrep
go install github.com/keeferrourke/imgrep
```

## Usage

`imgrep` like `grep`, searches file contents for text. `imgrep` works
exclusively on images however, using text extracted using OCR as the
search haystack.

`imgrep` comes with two interfaces; a CLI as one might expect, and a
web-UI graphical front-end.

### CLI

Because OCR is expensive to the CPU, `imgrep` can pre-process and
index files by keywords stored to a database. This database is queried
unless specified otherwise. To preindex an entire
directory (including subdirectories):

```
imgrep init # pre-process and index image files in working directory
```

By default, `imgrep` uses this database of pre-indexed files to perform
simple queries. Because of the nature of OCR, picked up keywords may not
be accurate, so a sort of "fuzzy-search" is employed here. To mimic the
original usage of `grep`, `imgrep` queries are case-sensitive. The `-i`
option is provided to ignore case specifiers of query strings &mdash;
it is nearly always recommended to run your queries with this option.

To check preindexed directories:

```
imgrep search -i QUERY
```

To use `imgrep` without checking against the database of preindexed
files, simply call

```
imgrep search -ni QUERY
```

Like the `grep` family of functions, `imgrep` is useful with Unix-pipes:

```
# Example: Count the number of images that contain the first line of a
#          plain-text file
head -n1 myfile | imgrep search -n -i - | wc -l

# Example: Open the first file that matches a search
xdg-open "$(imgrep s learn | head -n1)"
```

## Web UI
See [`imgrep-web`](https://github.com/keeferrourke/imgrep-web).

## License
`imgrep` is free software licensed under the MIT license.

Copyright (c) 2017 Keefer Rourke and Ivan Zhang.

See LICENSE for details.
