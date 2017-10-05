[![Build Status](https://api.travis-ci.org/crazcalm/flash-cards.svg?branch=master)](https://travis-ci.org/crazcalm/flash-cards)[![Go Report Card](https://goreportcard.com/badge/github.com/crazcalm/flash-cards)](https://goreportcard.com/report/github.com/crazcalm/flash-cards) ![cover.run go](https://cover.run/go/github.com/crazcalm/flash-cards/cards.svg)

# Minimal Terminal Flashcard App

![](img/flashcard_app.png)

## How it works
### Manage your flashcards with a csv file
- The app takes a csv file as input.

The csv file should be of the form:

	Note: The headers "front", "back", and "hint" must exist
|front|back|hint|
|-----|----|----|
|f    |b   |h   |
|你好  |ni3hao3|hello|

### Interface

	Usage of ./flash-cards:
 	-f string
		file: path to csv file
  	-s 	Shuffle the cards (default true)
