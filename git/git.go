package git

import (
	"log"
	"os"
	"os/exec"
)

// Init initializes a git repo in the given path.
// If the path does not already exists, it will be created.
func Init(path string) {
	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		log.Fatal(err)
	}
	err := exec.Command("git", "init", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Add wraps the git add call and will change into the path of the git repo
// and add the file given by name.
func Add(path, filename string) {
	os.Chdir(path)
	err := exec.Command("git", "add", filename).Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Commit will change into the path of the git repo and execute git commit.
// In addition a message and date for the commit are specified.
func Commit(path, message, date string) {
	os.Chdir(path)
	err := exec.Command("git", "commit", "-m", message, "--date", date).Run()
	if err != nil {
		log.Fatal(err)
	}
}