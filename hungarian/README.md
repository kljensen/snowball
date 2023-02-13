Snowball Hungarian
================

This package implements the
[Hungarian language Snowball stemmer](https://snowballstem.org/algorithms/hungarian/stemmer.html)
algorithm by [atordai@science.uval.nl](Anna Tordai).

## Implementation

The Hungarian language stemmer comprises preprocessing, a number of steps,
and postprocessing.  Each of these is defined in a separate file in this
package.  All of the steps operate on a `SnowballWord` from the
`snowballword` package and *modify the word in place*.

