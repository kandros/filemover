package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type filemover struct {
	source            string
	destination       string
	destinationFolder string
}

func main() {
	if len(os.Args) != 2 {
		println("Expected a single argument [folder path]")
		os.Exit(1)
	}
	folderPath := os.Args[1]

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	movers := []filemover{}

	for _, file := range files {
		if file.Name() == "_.DS_Store" {
			continue
		}
		if file.IsDir() {
			continue
		}

		date := file.ModTime()

		formattedDate := fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
		destinationFolder := filepath.Join(folderPath, formattedDate)
		movers = append(movers, filemover{
			source:            filepath.Join(folderPath, file.Name()),
			destinationFolder: destinationFolder,
			destination:       filepath.Join(destinationFolder, file.Name()),
		})
	}

	for _, m := range movers {
		if _, err := os.Stat(m.destinationFolder); err != nil {
			if os.IsNotExist(err) {
				err := os.Mkdir(m.destinationFolder, 0777)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		err := os.Rename(m.source, m.destination)
		if err != nil {
			log.Fatal(err)
		}
	}
}
