package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"backendbillingdashboard/autocrud/copy"
)

var moduleLocation = "../modules/"

// IF YOU WANT TO BUILD AS EXECUTABLE FILE, CHANGE moduleLocation to this one, so file can be placed on your root project directory
// var moduleLocation = "./modules/"

var stubFileName = "stub"
var stubProperName = "Stub"
var tempName = "modtemp"

func main() {
	var moduleName string
	var ERR string

	// stop execution if stub directory not exists
	if _, err := os.Stat(moduleLocation + stubFileName); os.IsNotExist(err) {
		panic("Cannot generate module, stub directory in modules is not found.")
	}

	// wait for user input until get the correct modulename target name
	for len(moduleName) == 0 || moduleName == tempName {
		if moduleName == tempName {
			moduleName = WaitForUserInput("Please use another name. Please type new module name")
		} else if ERR == "exists" {
			moduleName = WaitForUserInput("Module name is already exists. Please type another module name")
		} else if ERR == "toolong" {
			moduleName = WaitForUserInput("Module name is too long. Please type another module name")
		} else {
			moduleName = WaitForUserInput("Please type new module name")
		}

		if len(moduleName) == 0 {
			continue
		}
		if len(moduleName) > 10 {
			ERR = "toolong"
			moduleName = ""
			continue
		}

		// check moduleName
		if _, err := os.Stat(moduleLocation + moduleName); !os.IsNotExist(err) {
			// module is already exists
			ERR = "exists"
			moduleName = ""
			continue
		}

		// reset ERR value
		ERR = ""
	}

	// start generate module
	fmt.Printf("Module autocrud generation start\n")
	GenerateModule(moduleName)
	fmt.Printf("Module \"%s\" generation finish", moduleName)

}

func WaitForUserInput(s string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s : ", s)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.TrimSpace(response)
		return response
	}
}

func GenerateModule(moduleName string) {
	moduleFilename := strings.ToLower(moduleName)
	moduleProperName := strings.Title(moduleName)

	// copy from stub -> new module
	err := copy.Copy(moduleLocation+stubFileName, moduleLocation+moduleFilename)
	if err != nil {
		panic(err)
	}

	// listing module from stub
	fileList := ListingModuleFiles(moduleFilename)

	// rename new module filename
	err = RenameNewModule(fileList, moduleProperName)
	if err != nil {
		panic(err)
	}
}

func ListingModuleFiles(moduleFilename string) []string {
	//listing all files
	var fileList []string
	err := filepath.Walk(moduleLocation+moduleFilename, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return fileList
}

func RenameNewModule(pathLists []string, moduleProperName string) (err error) {
	for _, paths := range pathLists {
		lastext := paths[len(paths)-3:]
		if lastext != ".go" {
			continue
		}

		// rename this file
		newPath := strings.Replace(paths, stubProperName, moduleProperName, 1)
		if newPath != paths {
			// rename
			err = os.Rename(paths, newPath)
			if err != nil {
				panic(err)
			}

			ReplaceInnerContent(newPath, moduleProperName)
		} else {
			// no need to rename file, inner content only
			ReplaceInnerContent(newPath, moduleProperName)
		}
	}
	return
}

func ReplaceInnerContent(newPath, moduleProperName string) {
	moduleLowerName := strings.ToLower(moduleProperName)

	filecontent, err := ioutil.ReadFile(newPath)
	if err != nil {
		panic(err)
	}

	stringcontent := string(filecontent)
	stringcontent = strings.Replace(stringcontent, stubProperName, moduleProperName, -1)
	stringcontent = strings.Replace(stringcontent, stubFileName, moduleLowerName, -1)

	err = ioutil.WriteFile(newPath, []byte(stringcontent), 0644)
	if err != nil {
		panic(err)
	}
}
