# imgrep

`imgrep` is a command-line utility written in Go to search for keywords
found within images.

`imgrep`'s server scans the user Pictures directory and processes files
using Tesseract. File/keyword associations are stored in a small sqlite
database.
