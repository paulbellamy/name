# name

Super-simple golang library for generating friendly Heroku-style names. I've copy-pasted it
around a few projects, and figured it's time to turn it into a proper library.

Some example names:

```
feisty-haze-34
gentle-meadow-66
twilight-silence-38
long-rain-82
summer-lake-10
giant-glade-02
black-moon-92
red-sun-54
still-frog-05
dark-wildflower-69
```

## Usage

Make sure you seed the random number generator before using this. For example:
```
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/paulbellamy/name"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf(name.Generate())
}
```
