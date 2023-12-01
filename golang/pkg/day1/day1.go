package day1

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var letterLookup map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func processLines(lines []string, regexp *regexp.Regexp, sum *int64, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
    localSum := int64(0)

	for _, line := range lines {
        first := ""
        last := ""
		for i := 0; i < len(line); i++ {
			loc := regexp.FindStringIndex(line[i:])
			if loc != nil {
				start, end := loc[0]+i, loc[1]+i
				match := line[start:end]

				n, ok := letterLookup[match]
                if ok {
					match = n
				}

                if first == "" {
                    first = match
                } else {
                    last = match
                }
                if ok {
                    i = start + 1
                } else {
                    i = end - 1
                }
			}
		}

        if last == "" {
            last = first
        }
        asNum, err := strconv.ParseInt(first + last, 10, 64)
        if err == nil {
            localSum += asNum
        }
	}


	mu.Lock()
	*sum += localSum
	mu.Unlock()
}


func Solve() {
	inBuffer, err := ioutil.ReadFile("day1-input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(inBuffer)
	asLines := strings.Split(input, "\n")
	numbersRegexp, err := regexp.Compile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var sum int64

	batchSize := 1000
	for i := 0; i < len(asLines); i += batchSize {
		end := i + batchSize
		if end > len(asLines) {
			end = len(asLines)
		}

		wg.Add(1)
		go processLines(asLines[i:end], numbersRegexp, &sum, &wg, &mu)
	}

	wg.Wait()
	fmt.Println(sum)
}
