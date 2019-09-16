# Multiverse OS: Design Principals And Guidelines
=====

Multiverse OS is intended to be a collaborative effort, the combined effort of 
a dedicated team sometimes able to work full time; and volunteers who
encouraged to participate and contribute as the project takes shape and grows.

*The first public release will completely abstract and hide all complexity
associated with clustering, identity, activity, and process compartmentalizaiton
using full hardware virtualizatized ephemeral virtual machines.*


## Go Language
=====

While components of Multiverse OS will likely expand to be in Ruby, Rust, and 
C; a large portion of the codebase will inevtiably be in Go language.  Below is 
a evolving guide that all commits to the repositories should adhere to: 


#### No Global Variables
Global variables should not be used wtih the exception being small subpackages
that contain extensions, or extra variations, data or essentially instances of a
model defined by an interface. 

An example of this can be found in Multiverse OS's [`loader`
library](github.com/multiverse-os/loader) designed
to compliment the [`cli` framework](https://github.com/multiverse-os/cli).

In this exmaple, global variables are used for the spinner subpackages, which
contain a single global variable containing all the data needed to load the
spinner. In this design, all the spinner packages are optional, and so if a
spinner subpackage is only necessary to importat if the spinner is intended to
be used. And so the status of the global variable makes little to no differnece
in this case. 


#### Empty String Comparison 
When comparing against a blank string, `len(blankString)` should be compared
over a direct string comparison. The following may at first glance seem the
same: 

```Go
  if len("testString") == 0 {
    // This is the correct way to do blank string comparison
  }

  if "testString" == "" {
    // Do not do this type of blank string comparison
  }
```


Using the Multiverse OS [`benchmark`
library](https://github.com/multiverse-os/benchmark) Multiverse OS developers
regularly compare different methods of accomplishing a task. And regularly
re-check these conclusions after major updates to ensure there has been to major
changes. And even after running the gbenchmark for 12 hours, every test came out
with the `int` comparison version using `len()` preforms better than the blank
`string` comparison. 

```Bash
[benchmark] test took [ 207.695µs ms ] to complete 
Testing Method A: 'String' == ''
[benchmark] function being tested [ 0x495b80 ] with the name [ main.main.func1 ]
[benchmark] test has ran [ 25000 times ] with an average of [ 1.257µs  microseconds ]
Testing Method B: len('String') == 0
[benchmark] function being tested [ 0x495bb0 ] with the name [ main.main.func2 ]
[benchmark] test has ran [ 25000 times ] with an average of [ 572ns  microseconds ]
[benchmark] test has ran [ 25000 times ] with an average of [ 572ns  microseconds ]
Best function is: 0x495bb0

```


In addition, this also means that we may want to just create functions that will
be included in a `common` type library: 


```Go 

func Blank(s string) bool { return (len(s) == 0) } 
func NotBlank(s string) bool { return (len(s) > 0) }
```

## Comments And Variable Verbosity
Go Language is a compiled language, and the compiler is able to minify the
variables enabling developers to have whatever length without negatigvely
affecting preformance. 

With this being the case, Multiverse OS developers want to create a synergy
between a Ruby-like scripting language that is very clear and easy to read, and
style of Go programming that follows the established pattern. By always opting
for full names, focusing on human readability, function names that make sense,
and design that reads like a recipie book over convuluted or overly complex
designs; we often reduce the codebase of any library we refactor, despite our
long variable names and increase the speed, eliminate memory leaks, and patch
security issues. 

This style of Go programming also eliminates the requirement to put a sentence
above each function that explains what the function does, because the function
itself explains what it does. 

Leaving comments that actually 'pop', or stick out, giving them greater effect.
And for those who favor the single line above each function for documentaiton
creation, we feel documentation is better served by creating a variety of good
examples, and pulling information from our CLI, API and other locations in our
software where help text is already naturally being written. 




