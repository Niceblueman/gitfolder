package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/go-github/v53/github"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <GitHub_Folder_Link>")
		os.Exit(1)
	}

	folderLink := args[0]
	if !strings.Contains(folderLink, "/tree/") {
		fmt.Println("Invalid GitHub folder link.")
		os.Exit(1)
	}

	// Initialize GitHub client
	ctx := context.Background()
	client := newGitHubClient()

	// Parse owner, repository, and branch from the folder link
	parts := strings.Split(strings.TrimPrefix(folderLink, "https://github.com/"), "/")
	if len(parts) < 4 {
		fmt.Println("Invalid GitHub folder link format.")
		os.Exit(1)
	}

	owner := parts[0]
	repo := parts[1]
	branch := parts[3]
	lastPathComponent := parts[len(parts)-1]

	// Get the list of files in the folder
	fileLinks, err := getFileLinksFromGitHub(ctx, client, owner, repo, branch, strings.Join(parts[4:], "/"))
	if err != nil {
		fmt.Printf("Error getting file links: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d file(s) in the folder.\n", len(fileLinks))

	// Create a folder with the last path component name
	err = os.MkdirAll(lastPathComponent, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating folder: %s\n", err)
		os.Exit(1)
	}

	err = downloadFiles(lastPathComponent, fileLinks)
	if err != nil {
		fmt.Printf("Error downloading files: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Download successful!")
}

func newGitHubClient() *github.Client {
	return github.NewClient(nil)
}

func getFileLinksFromGitHub(ctx context.Context, client *github.Client, owner, repo, branch, path string) ([]string, error) {
	_, contents, _, err := client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{
		Ref: branch,
	})
	if err != nil {
		return nil, err
	}

	var fileLinks []string
	for _, content := range contents {
		if *content.Type == "file" {
			fileLinks = append(fileLinks, *content.DownloadURL)
		}
	}

	return fileLinks, nil
}

func downloadFiles(folderName string, fileLinks []string) error {
	for _, link := range fileLinks {
		resp, err := http.Get(link)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		fileName := path.Base(link)
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path.Join(folderName, fileName), data, 0644)
		if err != nil {
			return err
		}

		fmt.Printf("Downloaded: %s\n", fileName)
	}

	return nil
}
