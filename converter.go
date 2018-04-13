package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"os"
	"flag"
	"bufio"
	"strings"
)

var log = logrus.New()

// Tiny script to fast convert log files to csv
// Example usage: go run package -file={inputfullfilepath} -sep={logseparator} -outfile={outputfullfilepath}
// Preferable to put separator between "" to avoid unexpected results
func main() {
	start := time.Now()
	log.Out = os.Stdout

	fileName := flag.String("file", "foo", "a string")
	separator := flag.String("sep", ";", "a string")
	outputFile := flag.String("outfile", ";", "a string")
	flag.Parse()

	fileHandle, err := os.Open(*fileName)
	defer fileHandle.Close()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		slice := strings.Split(text, *separator)
		csvLine := strings.Join(slice[:], ";")
		err = writeLineToFile(*outputFile, csvLine+"\n")
		if err != nil {
			log.Error(err)
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Infof(elapsed.String())
}

func writeLineToFile(outputFile string, line string) error {
	var file, err = os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(line)

	return nil
}
