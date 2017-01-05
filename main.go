package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

func main() {
	// Need to have url to clone
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}
	repo := os.Args[1]

	// Make sure git is there
	binary, err := exec.LookPath("git")
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

	// Location where we want to clone to
	whereToCloneTo := path.Join(rootPath, repoNamespace, repoName)

	// Arguments to git
	args := []string{"git", "clone", repo, whereToCloneTo}

	// Make use of already set location
	env := os.Environ()

	// Actually execute clone
	err = syscall.Exec(binary, args, env)
	if err != nil {
		panic(err)
	}
}
