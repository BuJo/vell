package deb

import (
	"github.com/rkcpi/vell/repos"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type repository struct {
	store *repoStore
	name  string
}

func newRepository(store *repoStore, name string) repos.AnyRepository {
	return &repository{store, name}
}

func (r *repository) path() string {
	return filepath.Join(r.store.base, r.name)
}

// Add a file to a repository
func (r *repository) Add(filename string, f io.Reader) {
	log.Printf("Adding %s to repository %s", filename, r.path())
	tempPath, _ := ioutil.TempDir("", "vell")
	defer os.RemoveAll(tempPath)

	destinationPath := filepath.Join(tempPath, filename)
	destination, err := os.Create(destinationPath)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(destination, f)
	if err != nil {
		panic(err)
	}
	destination.Close()

	log.Printf("Executing `reprepro includedeb %s %s`", r.name, destinationPath)
	cmd := exec.Command("reprepro", "--base", r.path(), r.name, destinationPath)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func (r *repository) Update() error {
	return nil
}

func (r *repository) IsValid() bool {
	return true
}

func (r *repository) ListPackages() []repos.Package {
	files, _ := ioutil.ReadDir(filepath.Join(r.path(), "pool"))
	packages := make([]repos.Package, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			p := repos.Package{
				Name:      file.Name(),
				Timestamp: file.ModTime().Format(time.RFC3339),
				Size:      file.Size(),
			}
			packages = append(packages, p)
		}
	}
	return packages
}
