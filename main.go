package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// Need to have url to clone
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}
	repo := os.Args[1]

	// Make sure git is there
	_, err := exec.LookPath("git")
	if err != nil {
		panic(err)
	}

	// Get the path where to clone to
	rootPath := os.Getenv("ERIO_PATH")
	if rootPath == "" {
		log.Fatal("ERIO_PATH was not defined...")
	}

	// Figure out namespace/name of repository
	repoNamespace := ""
	repoName := ""
	if len(strings.Split(repo, "@")) == 2 {
		// if "git@github.com:ipfs/notes.git" format
		fullRepoName := strings.Split(strings.Split(repo, ":")[1], ".")[0]
		repoNamespace = strings.Split(fullRepoName, "/")[0]
		repoName = strings.Split(fullRepoName, "/")[1]
	} else {
		// if "ipfs/notes" format
		repoNamespace = strings.Split(repo, "/")[0]
		repoName = strings.Split(repo, "/")[1]

		repo = "git@github.com:" + repoNamespace + "/" + repoName + ".git"
	}

	// Make namespace + reponame lowercase
	repoNamespace = strings.ToLower(repoNamespace)
	repoName = strings.ToLower(repoName)

	// Location where we want to clone to
	whereToCloneTo := path.Join(rootPath, repoNamespace, repoName)

	fmt.Println("Cloning " + repo + " into " + whereToCloneTo)

	// Setup the execution of git
	cmd := exec.Command("git", "clone", repo, whereToCloneTo)

	// Actually run the command
	out, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "already exists") {
			fmt.Println("Already existed")
		} else {
			fmt.Println(string(out))
			panic(err)
		}
	}

	// Write cloning path to clipboard
	err = clipboard.WriteAll(whereToCloneTo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Copied the clone destination path to your clipboard!")
}
