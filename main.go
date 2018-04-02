package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	optNum := flag.Int("n", 5, "output line num")

	flag.Parse()

	for _, v := range flag.Args() {
		hp := HeadPrinter{path: v, n: *optNum, isPrintName: len(flag.Args()) > 1}
		hp.printHead()
	}
}

type HeadPrinter struct {
	path        string
	n           int
	isPrintName bool
}

func (hp *HeadPrinter) printHead() {
	file, err := os.Open(hp.path)
	if err != nil {
		return
	}
	defer file.Close()

	if hp.isPrintName {
		fmt.Printf("==> %s <==\n", file.Name())
	}

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan() && i <= hp.n; i++ {
		if err := sc.Err(); err != nil {
			break
		}
		fmt.Printf("%s\n", sc.Text())
	}
	fmt.Print("\n")
}
