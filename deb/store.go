package deb

import (
	"github.com/rkcpi/vell/repos"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type repoStore struct {
	base string
}

// Instanciate a new Debian based repository store by providing a base path.
func NewRepositoryStore(base string) repos.RepositoryStore {
	return &repoStore{base}
}

func (store *repoStore) Get(name string) repos.AnyRepository {
	return newRepository(store, name)
}

func (store *repoStore) Initialize(name string) error {
	log.Printf("Initializing repository %s", name)
	path := store.ensureExists(name)

	store.ensureExists(name, "conf")

	err := ioutil.WriteFile(
		store.distributionsPath(name),
		[]byte("Origin: "+name+"\n"+
			"Label: "+name+"\n"+
			"Codename: "+name+"\n"+
			"Architectures: amd64\n"+
			"Components: main\n"+
			"Description: Apt repository\n"),
		0644)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(
		store.optionsPath(name),
		[]byte("verbose\n"+
			"basedir "+path+"\n"),
		0644)

	return err
}

func (store *repoStore) ListRepositories() []repos.Repository {
	files, err := ioutil.ReadDir(store.base)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	reps := make([]repos.Repository, 0, len(files))
	for _, file := range files {
		repo := repos.Repository{Name: file.Name()}
		reps = append(reps, repo)
	}

	return reps
}

func (store *repoStore) ensureExists(name ...string) string {
	path := filepath.Join(append([]string{store.base}, name...)...)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Creating repository directory %s", path)
		os.MkdirAll(path, 0755)
	}
	return path
}

func (store *repoStore) distributionsPath(name string) string {
	return filepath.Join(store.base, name, "conf", "distributions")
}
func (store *repoStore) optionsPath(name string) string {
	return filepath.Join(store.base, name, "conf", "options")
}
