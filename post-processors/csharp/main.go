package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("exactly one directory name must be provided to run post-processing on")
	}
	dirPath := os.Args[1]
	log.Printf("running post processing for C# SDK on directory %v", dirPath)

	files = make(map[string]fs.FileInfo)
	err := filepath.Walk(dirPath, walkFiles)
	if err != nil {
		return fmt.Errorf("could not read directory %v: %v", dirPath, err)
	}

	for path, file := range files {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read file %v: %v", file.Name(), err)
		}
		fileContents := string(fileBytes)
		if len(fileContents) < 1 {
			return fmt.Errorf("input file %v must not be empty", file.Name())
		}

		if file.Name() == "AutolinksPostRequestBody.cs" {
			fileContents = fixStringToBool(fileContents)
		}

		if file.Name() == "WebhookConfigInsecureSsl.cs" {
			fileContents = fixUsing(fileContents)
		}

		fileContents = fixThumbsUpThumbsDownProperties(fileContents)
		fileContents = fixPagesPaths(fileContents)

		err = os.WriteFile(path, []byte(fileContents), 0600)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO(kfcampbell) fix this sleazy global
var files map[string]fs.FileInfo

func walkFiles(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("error walking files: %v", err)
	}

	// skip directories
	if info.IsDir() {
		return nil
	}
	files[path] = info
	return nil
}

func fixStringToBool(inputFile string) string {
	// find: IsAlphanumeric = "true";
	// replace: IsAlphanumeric = true;
	inputFile = strings.ReplaceAll(inputFile, "IsAlphanumeric = \"true\";", "IsAlphanumeric = true;")
	return inputFile
}

func fixUsing(inputFile string) string {
	// find: using Microsoft.Kiota.Abstractions;
	// replace: using Microsoft.Kiota.Abstractions.Serialization;
	inputFile = strings.ReplaceAll(inputFile, "using Microsoft.Kiota.Abstractions;", "using Microsoft.Kiota.Abstractions.Serialization;")
	return inputFile
}

func fixThumbsUpThumbsDownProperties(inputFile string) string {
	/*
		One,
		One,
	*/
	toReplace := `{
        One,
        One,`

	replaceWith := `{
        One,
        MinusOne,`

	if strings.Contains(inputFile, toReplace) {
		inputFile = strings.ReplaceAll(inputFile, toReplace, replaceWith)
	}

	return inputFile
}

func fixPagesPaths(inputFile string) string {
	// find: Path = PagesPostRequestBody_source_path.;
	// replace: Path = PagesPostRequestBody_source_path.Root;
	inputFile = strings.ReplaceAll(inputFile, "Path = PagesPostRequestBody_source_path.;", "Path = PagesPostRequestBody_source_path.Docs;")
	return inputFile
}
