MarGopher
---

Simple [Markov chain](http://www.wikiwand.com/en/Markov_chain) random text generator in Go

## Installation

```
go get github.com/AhmedZaleh/margopher
```

## Usage

```
import "github.com/AhmedZaleh/margopher"

text := "I love cats. Cats love pizza"
sentenceLength := 6

margopher := margopher.NewMargopher()
margopher.ParseText(text)
fmt.Println(margopher.Generate(sentenceLength))
```

## TODO

- [ ] Remove 'Simple' from the description and add more features

## License

See [LICENSE](https://github.com/AhmedZaleh/margopher/blob/master/LICENSE).
