Snowball
========

A [Go](http://golang.org) implementation of the
[Snowball stemmer](http://snowball.tartarus.org/)
for natural language processing.  The project currently only includes
the [English stemmer](http://snowball.tartarus.org/algorithms/english/stemmer.html).


## Usage

The `snowball` package has a single exported function `snowball.Stem`,
which is defined in `snowball/snowball.go`.  Each language also exports
a `Stem` function: e.g. `english.Stem`, which is defined in
`snowball/english/stem.go`.

Here is a minimal Go program that uses this package in order
to stem a single word.

```go
package main
import (
	"fmt"
	"github.com/kljensen/snowball"
)
func main(){
	stemmed, err := snowball.Stem("Accumulations", "english", true)
	if err == nil{
		fmt.Println(stemmed) // Prints "accumul"
	}
}
```


## Status

Only the English stemmer is implemented; however, I'd like to add others.
The English stemmer produces the same output as the stemmer Snowball
language stemmer given the sample vocabulary
[here](http://snowball.tartarus.org/algorithms/english/stemmer.html).


## Implementation

I would like to mention here a few details about
the manner in which the stemmers (currently only English) are implemented.

* In order to ensure the code is easily extended to non-English lanuages,
  I avoided using bytes and byte arrays, and instead perform all operations
  on runes.  See `snowball/snowballword/snowballword.go` and the 
  `SnowballWord` struct.
* In order to avoid casting strings into slices of runes numerous times,
  this implementation uses a single slice of runes stored in the `SnowballWord`
  struct for each word that needs to be stemmed.
* Instead of carrying around the word regions R1 and R2 as separate strings
  (or slices or runes, or whatever), we carry around the index where each of
  these regions begins.  These are stored as `R1start` and `R2start` on the 
  `SnowballWord` struct. I believe this is a relatively efficient way of
  storing R1 and R2.
* I tried to avoided all maps and regular expressions 1) for kicks, and 2) because
  I thought they'd negatively impact the performance. 


## Future work

I'd like to implement the Snowball stemmer for other lanuages, particularly Spanish.
If you can help, I would greatly appreciate it: please fork the project and send
a pull request!

(Also, if you are interested in creating a larger NLP project for Go, please get in touch.)


## Contributors

* Kyle Jensen (kljensen@gmail.com, @DataKyle)
* Your name here should be here, seriously.


## License (MIT)

Copyright (c) 2013 the Contributors (see above)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.