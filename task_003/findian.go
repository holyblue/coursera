package main

import (
	//"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	//"os"

	s "strings"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func main() {
	// fmt.Println("Please enter some string")
	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Printf("Input was: %q\n", line)
	// 	line = s.ToLower(line)

	// 	if checkRule(line) {
	// 		fmt.Println("Found!")
	// 	} else {
	// 		fmt.Println("Not Found!")
	// 	}
	// }
	//main4()
	readFile()
}

func checkRule(input string) bool {
	if s.HasPrefix(input, "i") && s.HasSuffix(input, "n") && s.Contains(input, "a") {
		return true
	}

	return false
}

func main2() {
	s := strings.Replace("ni|ni|ni|ni|ni", "ni", "in", 5)
	fmt.Println(s)
}

func main3() {
	x := [...]int{1, 2, 3, 4, 5, 6, 8, 9}

	for i, v := range x {
		fmt.Printf("index=%d ---- value=%d \n", i, v)
	}

}

func main4() {
	var idMap = make(map[string]int)

	idMap["blues"] = 99

	for key, value := range idMap {
		fmt.Printf("%s %d", key, value)
	}
}

func decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func readFile() {
	content, err := ioutil.ReadFile("/Users/blues/Desktop/big5")
	if err != nil {
		log.Fatal(err)
	}

	unicodeContent, _ := decode(content)
	fmt.Printf("File contents: %s", unicodeContent)
}
