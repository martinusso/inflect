# inflect

[![Circle CI](https://circleci.com/gh/martinusso/inflect.svg?style=svg)](https://circleci.com/gh/martinusso/inflect)
[![Build Status](https://travis-ci.org/martinusso/inflect.svg?branch=master)](https://travis-ci.org/martinusso/inflect)
[![GoDoc](https://godoc.org/github.com/martinusso/inflect?status.svg)](https://godoc.org/github.com/martinusso/inflect)


## IntoWords()

IntoWords convert numbers (float64) to words

```go 
func IntoWords(number float64) string 
```

```go
got := IntoWords(42)  // -> forty-two
got := IntoWords(-147)  // -> Minus one hundred and forty-seven
```

## IsConsonant

IsConsonant check if a character is a consonant

```go 
func IsConsonant(s string) bool
```

## IsVowel

IsVowel check if a character is a vowel

```go 
func IsVowel(s string) bool
```

## Ordinal

Ordinal returns the ordinal suffix that should be added to a number to denote the position in an ordered sequence such as 1st, 2nd, 3rd, 4th...

```go 
func Ordinal(number int) string
```

```go
got := Ordinal(-1)  // -> st
got := Ordinal(42)  // -> nd
```

## Ordinalize

Ordinalize turns a number into an ordinal string

```go 
func Ordinalize(number int) string
```

```go
got := Ordinal(-1)  // -> -1st
got := Ordinal(42)  // -> 42nd
```

## Parameterize

Parameterize replaces special characters in a pretty string

```go 
func Parameterize(s string) string
```

```go
got := Parameterize("J. R. R. Tolkien")  // -> "j-r-r-tolkien"
```

### ParameterizeSep

ParameterizeSep replaces special characters, according to the separator, in a string

```go 
func ParameterizeSep(s, sep string) string
```

```go
got := ParameterizeSep("J. R. R. Tolkien")  // -> "j_r_r_tolkien"
```


## Pluralize

Pluralize generates the plurals of nouns

```go 
func Pluralize(word string) string
```

```go
got := Plural("word") // -> words
got := Plural("noun") // -> nouns
```

## License

This software is open source, licensed under the The MIT License (MIT). See [LICENSE](https://github.com/martinusso/inflect/blob/master/LICENSE) for details.
