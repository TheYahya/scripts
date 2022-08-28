package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Println("Execution time: ", duration)
	}()

	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalln("We need dir_path and ext you want to delete, e.g: remove-ext-files /home/user/pics/ .jpg")
	}

	err := removeFilesWithExt(args[0], args[1])
	if err != nil {
		log.Fatalln(err)
	}
}

func removeFilesWithExt(dir, ext string) error {
	var jpgSize int64 = 0
	var fullSize int64 = 0

	messages := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileExt := filepath.Ext(path)
		of, err := os.Open(path)
		if err != nil {
			return err
		}

		stat, err := of.Stat()
		if err != nil {
			return err
		}

		if strings.ToLower(fileExt) == strings.ToLower(ext) {
			err := os.Remove(path)
			if err != nil {
				messages = append(messages, fmt.Sprintf("can't delete file: %s, err: %s", path, err.Error()))
			}
			log.Printf("deleted %s\n", path)

			jpgSize += stat.Size()
		}
		fullSize += stat.Size()

		return nil
	})

	log.Printf("full: %dMB\n", fullSize/1024/1024)
	log.Printf("jpg: %dMB\n", jpgSize/1024/1024)
	return err
}
