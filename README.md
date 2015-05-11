# language-filter
Golang program to filter different languages out of text files using the lexicon repository

## Status
Ready to use

## Environment variables
### LEXICON_DATA
Set the LEXICON_DATA directory to the location of the lexicon data files, for example:
export LEXICON_DATA=~/go/src/github.com/BluntSporks/lexicon/data

## Usage
Usage:

    language-filter [options] FILENAME

Options:

    lang=LANG        Name of language file to inspect
    lexdir=DIR       Location of lexicon data directory
    percent=PERCENT  Minimum percentage of language to require to be included
