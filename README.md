# inflect

[![Circle CI](https://circleci.com/gh/martinusso/inflect.svg?style=svg)](https://circleci.com/gh/martinusso/inflect)
[![Build Status](https://travis-ci.org/martinusso/inflect.svg?branch=master)](https://travis-ci.org/martinusso/inflect)
[![GoDoc](https://godoc.org/github.com/martinusso/inflect?status.svg)](https://godoc.org/github.com/martinusso/inflect)



## IntoWords()

IntoWords convert numbers (float64) to words

```go func IntoWords(number float64) string ```

```go
got := IntoWords(42)  // -> forty-two
got := IntoWords(-147)  // -> Minus one hundred and forty-seven
```

## IsConsonant

`func IsConsonant(s string) bool`

IsConsonant check if a character is a consonant

## IsVowel

`func IsVowel(s string) bool`

IsVowel check if a character is a vowel

## Ordinal

`func Ordinal(number int) string`

Ordinal returns the ordinal suffix that should be added to a number to denote the position in an ordered sequence such as 1st, 2nd, 3rd, 4th...

```go
got := Ordinal(-1)  // -> st
got := Ordinal(42)  // -> nd
```

## Ordinalize

`func Ordinalize(number int) string`

Ordinalize turns a number into an ordinal string

```go
got := Ordinal(-1)  // -> -1st
got := Ordinal(42)  // -> 42nd
```

## Parameterize

`func Parameterize(s string) string`

Parameterize replaces special characters in a pretty string

```go
got := Parameterize("J. R. R. Tolkien")  // -> "j-r-r-tolkien"
```

### ParameterizeSep

`func ParameterizeSep(s, sep string) string`

ParameterizeSep replaces special characters, according to the separator, in a string

```go
got := ParameterizeSep("J. R. R. Tolkien")  // -> "j_r_r_tolkien"
```


## Pluralize

`func Pluralize(word string) string`

Pluralize generates the plurals of nouns

```go
got := Plural("word") // -> words
got := Plural("noun") // -> nouns
```

## License

This software is open source, licensed under the The MIT License (MIT). See [LICENSE](https://github.com/martinusso/inflect/blob/master/LICENSE) for details.
