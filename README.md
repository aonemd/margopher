MarGopher
---

[Markov chain](http://www.wikiwand.com/en/Markov_chain) random text generator in Go

## How it Works

It's based on Markov chain statistical model, except that it starts by choosing a random prefix to 
ensure more randomness in the genereated sentence.

The generator will keep generating words till it encounters a terminal word (a 
word that ends in '.')

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
  margopher := margopher.New()
  ```

3. Read input text using one of three parsing methods:

  - ReadText(text string)

  ```
  text := "I love cats. Cats love pizza."
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
  fmt.Println(margopher.Generate())
  ```

## License

See [LICENSE](https://github.com/AhmedZaleh/margopher/blob/master/LICENSE).
