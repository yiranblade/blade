package main

import (
	"blade/table"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func IsFileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if stat.Mode()&os.ModeType == 0 {
			return true, nil
		}
		return false, errors.New(path + " exists but is not regular file")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ParseDataFile(path string) (table *table.Table, err error) {

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, table); err != nil {
		return nil, err
	}
	return

}

func main() {
	log.SetOutput(os.Stdout)
	var tableDataFilePath string
	var printVer bool

	flag.BoolVar(&printVer, "version", false, "print version")
	flag.StringVar(&tableDataFilePath, "c", "tableData.json", "specify config file")
	flag.StringVar(&tableDataFilePath, "c", "tableData.json", "specify config file")

	flag.Parse()

	if printVer {
		fmt.Println("V0.1")
		os.Exit(0)
	}

	exists, err := IsFileExists(tableDataFilePath)

	binDir := path.Dir(tableDataFilePath)
	if (!exists || err != nil) && binDir != "" && binDir != "." {
		oldConfig := tableDataFilePath
		tableDataFilePath = path.Join(binDir, "config.json")
		log.Printf("%s not found, try config file %s\n", oldConfig, tableDataFilePath)
	}
	tableData, err := ParseDataFile(tableDataFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", tableDataFilePath, err)
			os.Exit(1)
		}
	}


}
