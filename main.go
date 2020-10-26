package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const grammarFile = "bnf-cfg.txt"
const templateFile = "template.txt"

func main() {
	rand.Seed(time.Now().UnixNano())
	// seed, err := strconv.ParseInt(os.Args[1], 10, 32)
	// check(err)
	//rand.Seed(4)
	cfg := make(map[string][]string)
	parse(grammarFile, cfg, addRule)
	parse(templateFile, cfg, recursiveGenerateHelper)
	
}

func recursiveGenerateHelper(template string, cfg map[string][]string){
	line := recursiveGenerate(template, cfg)
	fmt.Println(line)
}

func recursiveGenerate(template string, cfg map[string][]string) string {
	if !strings.Contains(template, "<") {
		return template
	}

	generated := ""
	templateArray := strings.Split(template, " ")
	for i, templateElem := range templateArray {
		if strings.Contains(templateElem, "<") { // If contains a non-terminal, parse and recurse
			newtemplateElem := ""
			ntermArray := strings.FieldsFunc(templateElem, splitNterms) //Some elements may not be non-terminals
			for _, posNterm := range ntermArray {
				if ruleArray, present := cfg[posNterm]; present { // If it's a key in the map, it's an nterm
					rule := randomChoice(ruleArray)
					newtemplateElem += recursiveGenerate(rule, cfg)
				} else {
					newtemplateElem += posNterm
				}
				
			}
			templateElem = newtemplateElem
		}
		if i != 0 {
			generated+= " " //Add space between words
		}
		generated += templateElem
	}
	return generated
}

func randomChoice(ruleArray []string) string {
	var length int = len(ruleArray)
	return ruleArray[rand.Intn(length)]
}

func splitNterms(r rune) bool {
	return r == '<' || r == '>'
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addRule(line string, cfg map[string][]string) {
	endIndex := strings.Index(line, ">")
	nonTerm := line[1:endIndex]
	rulesString := line[endIndex+4:]
	cfg[nonTerm] = strings.Split(rulesString, "|")
}

func parse(fileName string, cfg map[string][]string, fn func(string, map[string][]string)) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line, cfg)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func demo() {
	rand.Seed(time.Now().UnixNano())
	cfg := make(map[string][]string)
	cfg["last-names"] = []string{"Smith", "Herrera", "Goel", "Ng", "Ngozi"}
	var length int = len(cfg["last-names"])
	for i := 0; i <= 10; i++ {
		index := rand.Intn(length)
		fmt.Println(cfg["last-names"][index])
	}
}
