margopher
---

[Markov chain](http://www.wikiwand.com/en/Markov_chain) random text generator

## How it Works

**TL;DR** You give it some text, it gives you a random (and sane) chunk out of this text.

Basically margopher is based on Markov chain statistical model, except that it
starts by choosing a random prefix from the states dictionary to ensure more
randomness.

It starts by parsing the input text and putting it into states dictionary then
starts generating the output sentence.

The generator will keep generating words till it encounters a terminal word (a
word that ends in any of `.,:;?!`)

**States** is a map contains prefix as keys and suffix as values.

**Prefix** is an array of two consecutive words from the original text.

**Suffix** is a slice of all the words that occur after a given prefix.


## Installation

```sh
go get github.com/aaoiki/margopher
```

## Usage

1. Import the package

  ```go
  import "github.com/aaoiki/margopher"
  ```

2. Create new margopher object

  ```go
  m := margopher.New()
  ```

3. Read input text using one of three parsing methods:

  - ReadText(text string)

  ```go
  text := "I love cats. Cats love pizza."
  fmt.Println(m.ReadText(text))
  ```

  - ReadFile(filePath string)

  ```go
  filePath := "../file.txt"
  fmt.Println(m.ReadFile(filePath))
  ```

  - ReadURL(url string)

  ```go
  url := "https://github.com/aaoiki/margopher"
  fmt.Println(m.ReadURL(url))
  ```

4. You can see the input parsed into a dictionary of States using `margopher.StateDictionary()`:

  - It returns a dicitonary of this signature `map[[2]string][]string`.
  - The words are unordered because Go maps do not keep order.

  ```go
  fmt.Println(m.StateDictionary())
  ```

## Feedback

I wrote this simple project particularly to learn Go so any feedback is more
than welcome. If you have any, please open an issue or send a pull request.

## License

See [LICENSE](https://github.com/aaoiki/margopher/blob/master/LICENSE).
