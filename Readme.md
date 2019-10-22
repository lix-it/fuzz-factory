# Fuzz Factory
Makes data dirtier.

Accepts CSVs.
## Requirements
- Go >= 1.11

## Installation
*Build*

`go build .`

Or if you are feeling fancy and want to install it to your $PATH and run it from anywhere:
`go install ./...`
## Usage
`fuzz-factory [-d=dictionary.txt][-headers=false] [input] [output=output.csv]`

*flags*

    -d
        dictionary for random word replacements/additions/substitutions
    -headers
        input file has first row as headers



## TODO
- [x] header flag
- [ ] build the default dictionary into the binary
- [ ] add a fuzz metric. e.g. Levenshtein distance
- [ ] Input a desired fuzz factor on the command line and it will achieve that
