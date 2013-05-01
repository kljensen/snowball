Snowball
========

Pure Go implementation of the Snowball stemmer for English

## Caveats

This is a work in progress.  There has not been much progress.

## Thanks

This is based almost entirely on the [NLTK](http://nltk.org/)
version of
[Snowball](https://raw.github.com/nltk/nltk/master/nltk/stem/snowball.py),
which is described [here](http://snowball.tartarus.org/algorithms/english/stemmer.html).

## Warnings and Notes

I may have variously treated strings as utf8 and byte arrays.  This needs
to be remedied.  The reason that I've bothered with unicode is because
I'd like to (some day) implement the stemmers for other languages and 
I figured some of the code could be reused.

I tried to avoid maps and regular expressions for 1) kicks and 2) because
I thought they'd negatively impact the speed.


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