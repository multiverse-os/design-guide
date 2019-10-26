# Multiverse OS: Design Principals And Guidelines
Multiverse OS is intended to be a collaborative effort, the combined effort of 
a dedicated team sometimes able to work full time; and volunteers who
encouraged to participate and contribute as the project takes shape and grows.

*The first public release will completely abstract and hide all complexity
associated with clustering, identity, activity, and process compartmentalizaiton
using full hardware virtualizatized ephemeral virtual machines.*


## Go Language
While components of Multiverse OS will likely expand to be in `ruby`, `rust`, and 
`c`; a large portion of the codebase will inevtiably be in `go` language. Below is 
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
spinner. 

*The default example included with the binary program demonstrates this
specific benchmark, and to really get a better picture of the difference between
these two methods of checking for blank streams is by testing the benchmark on
the largest variety of hardware and foundational software.* 

If you decide to test this benchmark, please share your results with use so we 
can begin to build a statistical analysis program and datasheet to do better
analysis on this and future benchmark led design decisions.*

In this design, all the spinner packages are optional, and so if a
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
changes. And even after running the Multiverse OS maintained Go package `benchmark` 
for 12 hours, every test came out with the `int` comparison version using `len()` 
preforms better than the blank `string` comparison. 


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

## Inline declaration vs. Unix Package
Will update this later, but for now, the following is equivilent, so the 
better looking solution that provides more felxibility should be used:


```g
package main

import (
	"fmt"

	unix "golang.org/x/sys/unix"
)

func main() {
	// Write
	localVariable := unix.POLLIN
	// Read
	fmt.Println("localVariable:", localVariable)
}
```

The above code when compiled results in a `2MB` file. 

```go
package main

import (
	"fmt"
)

const (
	POLLIN = 0x0001
	//POLLPRI = 0x0002
	//POLLOUT = 0x0004
	//POLLERR = 0x0008
	//POLLHUP = 0x0010
	//POLLNVA = 0x0020
)

func main() {
	// Write
	localVariable := POLLIN
	// Read
	fmt.Println("localVariable:", localVariable)
}
```

The above code when compiled, will also result in a `2MB` file. 

## Comments And Variable Verbosity
Go Language is a compiled language, and the compiler is able to minify the
variables enabling developers to have whatever length without negatigvely
affecting preformance. 

With this being the case, Multiverse OS developers want to create a synergy
between a `ruby`-like scripting language that is very clear and easy to read, and
style of `go` programming that follows the established pattern. By always opting
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
examples, and pulling information from our `cli`, `api` and other locations in our
software where help text is already naturally being written. 

## Ordering Numeric Inequality Comparisons 
Multiverse OS developers should always opt for expressiveness so that the code 
is accessible the widest audience possible.  

This is accomplished by paying attention to small details that add up to more 
accessible code. 

When dealing with inequality comparisons: (`<`, `<=`, `==`, `>`, `>=`].

The expected lower value should always be first and the expected highest value  
last: 

```
if 0 < value && value < 100 {
  fmt.Println("value is between 0-100")
}

if value < 10 {
  fmt.Println("below 10") 
}
```

## Developers should never enforce preferences 
Multiverse OS devlepers should not enforce opinions or preferences with their design; 
for exmaple, a configuration file should be able to be loaded from any common data type 
[`TOML`, `JSON`, `YML`, `XML`].  

**The only time we force the users into a decision is when the design improves security, 
and we always default to a more secure design.**  






