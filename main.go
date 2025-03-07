package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func readFileAndPopulateMap(fileName string, m map[string]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		word := Scanner.Text()
		_, ok := m[word]
		if ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func outputTopNWords(n int, m map[string]int) {
	if len(m) <= n {
		for key, value := range m {
			fmt.Printf("%s: %d\n", key, value)
		}
	} else {
		for i := 0; i < n; i++ {
			if len(m) == 0 {
				return
			}

			var maxKey string
			maxNumber := 0

			for key, value := range m {
				if value >= maxNumber {
					maxKey = key
					maxNumber = value
				}
			}

			fmt.Printf("%s: %d\n", maxKey, maxNumber)
			delete(m, maxKey)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "uniq"
	app.Usage = "count uniq words in file and show top N"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "number",
			Usage: "top N to output, default 10",
			Value: 10,
		},
	}
	app.Action = func(c *cli.Context) error {
		fileName := c.Args().Get(0)
		m := make(map[string]int)

		readFileAndPopulateMap(fileName, m)
		outputTopNWords(c.Int("number"), m)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
