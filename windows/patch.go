package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// generate glob search path
func glob(root string, ext string) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func exit() {
	println("Press Enter to exit...")
	fmt.Scanln()
}

func main() {
	// print error message
	println("Report issues to https://github.com/mchangrh/NoWhatsNew/issues")
	// retreive CSS class name from JS
	cssName, err := findJsIndicator()
	if err != nil {
		println("error: " + err.Error())
	}
	println("Found CSS class name: " + cssName)
	// find and read matching CSS file
	cssFileContent, cssFileName, err := walkTester("C:\\Program Files (x86)\\Steam\\steamui\\css", ".css", "."+cssName)
	if err != nil {
		println("error: " + err.Error())
		exit()
		return
	}
	// patch CSS file
	patchedCss, err := patchCss(cssFileContent, cssName)
	if err != nil {
		println("error: " + err.Error())
		exit()
		return
	}
	// write patched CSS file
	err = writeFile(cssFileName, patchedCss)
	if err != nil {
		println("error patching file: " + err.Error())
		exit()
		return
	} else {
		println("Patched " + cssFileName)
		println("Restart Steam to see changes")
		exit()
		return
	}
}

func writeFile(file string, newContent string) error {
	writeFileErr := os.WriteFile(file, []byte(newContent), 0644)
	if writeFileErr != nil {
		return errors.New("error writing file " + writeFileErr.Error())
	}
	return nil
}

func findJsIndicator() (cssClassName string, err error) {
	// indicator for WhatsNew contianer
	indicator := "UpdatesContainer"
	fileContent, _, err := walkTester("C:\\Program Files (x86)\\Steam\\steamui", ".js", indicator)
	if (err != nil) {
		return "", err
	}
	// find indicator in file
	jsDefRegex := regexp.MustCompile(`UpdatesContainer:"(.+?)"`)
	matchArr := jsDefRegex.FindSubmatch([]byte(fileContent))
	if len(matchArr) < 2 {
		return "", errors.New("JS indicator not found - Please open an issue on GitHub")
	} else {
		cssClassName = string(matchArr[1])
		return cssClassName, nil
	}
}

func walkTester(root string, ext string, pattern string) (fileContents string, filename string, err error){
	files := glob(root, ext)
	for _, file := range files {
		readFileContent, readFileErr := os.ReadFile(file)
		if readFileErr != nil {
			return "", "", errors.New("error reading file " + readFileErr.Error())
		}
		fileContents = string(readFileContent)
		if strings.Contains(fileContents, pattern) {
			return fileContents, file, nil
		}
	}
	return "", "", errors.New("no matching " + ext + " files found")
}

func patchCss(fileContents string, className string) (newFileContents string, err error) {
	// create patch based on className
	patch := "div:has(>." + className + "){display:none}"
	// check if already patched
	if strings.Contains(fileContents, patch) {
		return "", errors.New("already patched")
	}
	// continue with patching
	// find .BasicUI child selector
	candidateLineRegex := regexp.MustCompile("\\.BasicUI \\." + className + "{.+?}")
	candidateLine := candidateLineRegex.FindString(fileContents)
	if (candidateLine == "") {
		return "", errors.New("candidate line not found - the file might be missing or changed")
	}
	// find difference in length
	paddingLength := len(candidateLine) - len(patch)
	patch = patch + strings.Repeat(" ", paddingLength)
	// replace in file
	fileContents = strings.Replace(fileContents, candidateLine, patch, 1)
	// Write file
	return fileContents, nil
}
