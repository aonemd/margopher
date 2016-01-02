MarGopher
---

[Markov chain](http://www.wikiwand.com/en/Markov_chain) random text generator

## How it Works

**TL;DR** You give it some text, it gives you a random (and sane) chunk out of this text.

Basically MarGopher is based on Markov chain statistical model, except that it
starts by choosing a random prefix from the states dictionary to ensure more
randomness.

It starts by parsing the input text and putting it into states dictionary then
starts generating the output sentence.

The generator will keep generating words till it encounters a terminal word (a
word that ends in '.')

**States** is a map contains prefix as keys and suffix as values.

**Prefix** is an array of two consecutive words from the original text.

**Suffix** is a slice of all the words that occur after a given prefix.


## Installation

`
go get github.com/aasare/margopher
`

## Usage

1. Import the package

  `
  import "github.com/aasare/margopher"
  `

2. Create new margopher object

  `
  m := margopher.New()
  `

3. Read input text using one of three parsing methods:

  - ReadText(text string)

  `
  text := "I love cats. Cats love pizza."
  m.ReadText(text)
  `

  - ReadFile(filePath string)

  `
  filePath := "../file.txt"
  m.ReadFile(filePath)
  `

  - ReadURL(url string)

  `
  url := "https://github.com/AhmedZaleh/margopher"
  m.ReadURL(url)
  `

4. Print the sentence

  `
  fmt.Println(m.Generate())
  `

## API

To use the API, add the following:

  `
  func main() {
    // Create New margopher
    m := margopher.New()

    // Call the Api() function
    m.Api()
  }
  `

Then send POST reqeusts to `http://localhost:7878/` using one of these paths:

  - /readtext
  - /readfile
  - /readurl

Example

  `
  $ curl -d '"We eat sushi. They eat pasta."' http://localhost:7878/readtext
  `

## Feedback

I wrote this simple project particularly to learn Go so any feedback is more
than welcome. Also, I'm currently investigating if there's a way to run the
program concurrently, if you think it's possible please tell me or open an
issue/send pull request.

## License

See [LICENSE](https://github.com/AhmedZaleh/margopher/blob/master/LICENSE).
