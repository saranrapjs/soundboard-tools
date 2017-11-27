# soundboard-tools

two tools for helping to record & play back a series of recordings of individual words

## Installation

```
go install github.com/saranrapjs/soundboard-tools/cmd/soundboard-rec
go install github.com/saranrapjs/soundboard-tools/cmd/soundboard-play
```

This requires [sox](http://sox.sourceforge.net/); install on a Mac with `brew install sox`.

## soundboard-rec

Given a file input, interactively record a dictionary of words as wav files within the current directory.

Usage:

```
echo "hello" > file-with-words.txt
echo "there" >> file-with-words.txt
echo "this" >> file-with-words.txt
echo "rules" >> file-with-words.txt
soundboard-rec file-with-words.txt
```

## woundboard-play

Given a folder with wav files named in the pattern `word.wav`, play a string of text containing those words.

Usage w/ args:

```
soundboard-play "hello there"
```

or w/ stdin:

```
echo "hello there" | soundboard-play
```
