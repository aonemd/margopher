MarGopher
---

[Markov chain](http://www.wikiwand.com/en/Markov_chain) random text generator in Go

## Installation

```
go get github.com/AhmedZaleh/margopher
```

## Usage

1. Import the package

```
import "github.com/AhmedZaleh/margopher"
```

2. Create new margopher object

```
margopher := margopher.NewMargopher()
```

3. Read input text using one of three parsing methods:

- ReadText(text string)

```
text := "I love cats. Cats love pizza"
margopher.ReadText(text)
```

- ReadFile(filePath string)

```
filePath := "../file.txt"
margopher.ReadFile(filePath)
```

- ReadURL(url string)

```
url := "https://github.com/AhmedZaleh/margopher"
margopher.ReadURL(url)
```

4. Print the sentence

```
sentenceLength := 3   //specify number of words
fmt.Println(margopher.Generate(sentenceLength))
```

## License

See [LICENSE](https://github.com/AhmedZaleh/margopher/blob/master/LICENSE).
