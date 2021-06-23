package main

import (
	"bufio"
	"os"
)

// can't be a const compile don't run functions
// even though once function runs it won't change again that function is only
// ran at runtime
var reader = bufio.NewReader(os.Stdin)
