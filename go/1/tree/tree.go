package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	errReading := readDir(out , path, printFiles, false, 0, "")
	if errReading != nil {
		return errReading
	}
	return nil
}

func readDir(out io.Writer, path string, printFiles bool, forNext bool, tab int, tabSymbol string) error {
	allFiles, err := ioutil.ReadDir(path)
	if err != nil {
		err = errors.New("error: can't read dir - " + path)
		return err
	}

	var files []os.FileInfo
	if !printFiles{
		for _, file := range allFiles {
			if file.IsDir() {
				files = append(files, file)
			}
		}
	} else {
		files = allFiles
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for i, file := range files{
		var isLast bool
		if i == len(files) - 1{
			isLast = true
		}
		if file.IsDir() {
			tabSymbol, errPrinting := printLine(out, path, file, tab, isLast, forNext, tabSymbol)
			if errPrinting != nil {
				return errPrinting
			}
			errReading := readDir(out , path + "/" + file.Name(), printFiles, isLast,  tab + 1, tabSymbol)
			if errReading != nil {
				return errReading
			}
		} else if printFiles {
			_, errPrinting := printLine(out, path, file, tab, isLast, forNext, tabSymbol)
			if errPrinting != nil {
				return errPrinting
			}
		}
	}

	return nil
}

func printLine(out io.Writer, path string, file os.FileInfo, tab int, isLast bool, forNext bool, tabSymbol string) (string, error) {

	var format, nestingSymbol, sizeSymbol string
	nestingSymbol = "├───"
	sizeSymbol = ""

	if forNext{
		tabSymbol = tabSymbol + "\t"
	} else {
		tabSymbol = tabSymbol + "│\t"
	}

	if tab == 0{
		tabSymbol = ""
	}

	if isLast {
		nestingSymbol = "└───"
	}

	if !file.IsDir(){
		sizeSymbol = fmt.Sprintf(" (%db)", file.Size())
		if file.Size() == 0{
			sizeSymbol = " (empty)"
		}
	}

	format = fmt.Sprintf("%s%s%s%s\n", tabSymbol, nestingSymbol, file.Name(), sizeSymbol)
	_, err := fmt.Fprintf(out, format)
	if err != nil {
		if file.IsDir(){
			err = errors.New("error: can't print dir - " + path)
		} else {
			err = errors.New("error: can't print file '" + file.Name() + "' dir - " + path)
		}
		return tabSymbol, err
	}

	return tabSymbol, nil
}
