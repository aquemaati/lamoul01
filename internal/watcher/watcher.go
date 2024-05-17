package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watcher(testFunc func(string, string, string), testPath string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("🤮\033[31mError getting current directory path:\033[0m", err)
		return
	}
	fmt.Println("🤟\033[32mCurrent directory path:\033[0m", dir)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🤓\033[34mWatching directory for changes...\033[0m")

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("🤫\033[33mFile modified: %s\033[0m\n", event.Name)

				// Get the file name without extension
				fileName := filepath.Base(event.Name)
				fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))

				// Execute the test function with the file name without extension
				testFunc(fileNameWithoutExt, dir, testPath)

			}
		case err := <-watcher.Errors:
			if err != nil {
				log.Println("🤮\033[31mError:\033[0m", err)
			}
		}

		// Short pause before resuming watching to avoid system overload
		time.Sleep(1 * time.Second)
	}
}
