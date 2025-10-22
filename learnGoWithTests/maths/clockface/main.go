package main

import (
	clockface "learnGoWithTests/maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)

}
