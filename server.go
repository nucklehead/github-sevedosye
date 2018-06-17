/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"strings"
	"io/ioutil"
	"gopkg.in/src-d/go-git.v4"
)

type Build struct {
    ID string `json:"id"`
    Number string `json:"number"`
    Name string `json:"name"`
    GithubUsername string `json:"github_username"`
    FinishedAt int64 `json:"finished_at"`
    CommitID string `json:"commit_id"`
}

var repoKlone *git.Repository

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", "./", "the directory of static file to host")
	nonNouvoDosye := flag.String("n", "builds", "the directory of cloned repo")
	repo := flag.String("r", "github.com/nucklehead/patecho-tes-rapo", "the repo to clone")
	token := flag.String("t", "token", "github token")
	flag.Parse()

	os.RemoveAll(*directory + *nonNouvoDosye)

	kloneRepo(*directory + *nonNouvoDosye, *repo, *token)

	http.Handle("/", http.FileServer(http.Dir(*directory)))
	http.HandleFunc("/api/meteAJou", meteKodAJou)
	http.HandleFunc("/api/builds", retounenBuilds)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func kloneRepo(directory string, repo string, token string){
	var err error
	repoKlone, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL:               fmt.Sprintf("https://%s@%s", token,repo),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
			log.Fatal(err)
	}
}

func retounenBuilds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listDosye, err := ioutil.ReadDir("./builds")
  if err != nil {
      log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
			return
  }
	var builds []Build
  for _, dosye := range listDosye {
		if(dosye.IsDir() && !strings.Contains(dosye.Name(), ".git")){
			komitIDMounKiKomit := strings.Split(dosye.Name(), ".")
			builds = append(builds,
				Build{
				ID: komitIDMounKiKomit[0],
				Number: komitIDMounKiKomit[2],
				Name: dosye.Name(),
				GithubUsername: komitIDMounKiKomit[3],
				CommitID: komitIDMounKiKomit[1],
				FinishedAt: dosye.ModTime().Unix()})
		}
	}
	json.NewEncoder(w).Encode(builds)
}

func meteKodAJou(w http.ResponseWriter, r *http.Request) {
	dosyeRepo, err := repoKlone.Worktree()
	if err != nil {
      log.Fatal(err)
  }
	err = dosyeRepo.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
      log.Fatal(err)
  }
}
