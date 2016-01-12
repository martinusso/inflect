# inflect

[![Circle CI](https://circleci.com/gh/martinusso/inflect.svg?style=svg)](https://circleci.com/gh/martinusso/inflect)
[![Build Status](https://travis-ci.org/martinusso/inflect.svg?branch=master)](https://travis-ci.org/martinusso/inflect)
[![GoDoc](https://godoc.org/github.com/martinusso/inflect?status.svg)](https://godoc.org/github.com/martinusso/inflect)


## IntoWords()

Convert numbers (float64) to words

```go
got := IntoWords(42)
// got -> forty-two
```

```go
got := IntoWords(-147)
// got -> Minus one hundred and forty-seven
```

## Plural("word") `func Plural(word string) string`

Plural generates the plurals of nouns

```go
got := Plural("word")
// got -> words

got := Plural("noun")
// got -> nouns
```

## License

This software is open source, licensed under the The MIT License (MIT). See [LICENSE](https://github.com/martinusso/inflect/blob/master/LICENSE) for details.
