package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	optNum := flag.Int("n", 1, "output num")

	flag.Parse()

	fmt.Println("n : ", *optNum)
	fmt.Println("files:", flag.Args())

	// execute
	for _, v := range flag.Args() {
		file, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		sc := bufio.NewScanner(file)

		fmt.Printf("%s\n", file.Name())
		//TODO nをintにしたい
		for i := 1; sc.Scan() && i <= 2; i++ {
			if err := sc.Err(); err != nil {
				break
			}
			fmt.Printf("%s\n", sc.Text())
		}
		fmt.Print("\n")
	}
}
