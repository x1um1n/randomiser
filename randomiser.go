package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/x1um1n/checkerr"
)

type rnam struct {
	old, new string
}

// shamelessly stolen from https://gist.github.com/albrow/5882501
// updated to use checkerr
func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	checkerr.CheckFatal(err, "error reading response from command line")

	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

// shamelessly stolen from https://gist.github.com/albrow/5882501
// used by askForConfirmation
func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

// rename builds a slice containing all the filenames in the dir
func rename(dir string) {
	var outfiles []rnam              //slice used to store rename params
	rand.Seed(time.Now().UnixNano()) //seed the random num generator

	infiles, err := ioutil.ReadDir(dir) //attempt to build a slice containing the files in the dir
	checkerr.CheckFatal(err, "Error reading dir contents")

	l := len(infiles) //get the initial length of infiles
	for i := 0; i < l; i++ {
		ido := rand.Intn(len(infiles))                      //pick an element at random
		fnam := infiles[ido].Name()                         //get the current filename
		idx := fmt.Sprintf("%03d", i)                       //create a new 3-digit index marker
		s := strings.TrimLeftFunc(fnam, func(r rune) bool { //trim any leading runes that are not letters, which allows for re-sorting
			return !unicode.IsLetter(r)
		})
		f := rnam{
			old: fnam,
			new: idx + " " + s,
		}
		outfiles = append(outfiles, f)
		infiles = append(infiles[:ido], infiles[ido+1:]...) //remove the element from infiles so we don't get dupes
	}

	//get the human to confirm this is ok
	for _, r := range outfiles {
		fmt.Printf("Rename %s to %s\n", r.old, r.new)
	}
	fmt.Printf("\nIs this ok?\n")
	if askForConfirmation() {
		if strings.HasPrefix(dir, "/") { //absolute path
			os.Chdir(dir)
		} else { //relative path
			wd, _ := os.Getwd()
			os.Chdir(wd + "/" + dir)
		}
		for _, r := range outfiles {
			if _, err := os.Stat(r.new); os.IsNotExist(err) {
				fmt.Printf("Renaming %s to %s\n", r.old, r.new)
				err := os.Rename(r.old, r.new)
				checkerr.CheckFatal(err)
			} else {
				fmt.Printf("Can't rename %s to %s, %s already exists..\n", r.old, r.new, r.new)
				fmt.Printf("\nContinue to next file?\n")
				if !askForConfirmation() {
					log.Fatalln("Exiting...")
				}
			}
		}
	} else {
		fmt.Println("Exiting without applying changes...")
	}
}

func usage() {
	fmt.Printf("Usage: randomiser [-strip] </dir/to/be/sorted>\n\n")
}

func main() {
	stripP := flag.Bool("strip", false, "Strip index and restore original sort order.")
	flag.Parse()

	if len(os.Args) == 1 { //no arguments, so print usage and exit
		usage()
		flag.PrintDefaults()
		fmt.Printf("\n")
		os.Exit(1)
	}

	dir := os.Args[len(os.Args)-1]                  //get the dir to randomise from cli args
	if _, err := os.Stat(dir); os.IsNotExist(err) { //sanity check the input is actually a dir
		checkerr.CheckFatal(err, "Dir does not exist")
	}

	if *stripP {
		// strip the indexes
		os.Exit(0)
	}

	rename(dir)
}
