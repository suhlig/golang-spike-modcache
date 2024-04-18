package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}
}

func mainE() error {
	repo, err := git.PlainOpen(".")

	if err != nil {
		return fmt.Errorf("opening the current directory as git repo: %w", err)
	}

	iter, err := repo.Log(&git.LogOptions{})

	if err != nil {
		return fmt.Errorf("getting the log: %w", err)
	}

	err = iter.ForEach(func(c *object.Commit) error {
		_, err = fmt.Println(c.Hash)
		return err
	})

	return err
}
