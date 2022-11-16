package user

import (
	"encoding/json"
	"errors"
	errs "github.com/PavelMo/user-api/err"
	"github.com/PavelMo/user-api/models"
	"io/fs"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"
)

const store = `users.json`

type UserStorage struct {
	mapMu     sync.RWMutex
	fileMu    sync.RWMutex
	userStore *UserStore
}

func NewUserStorage() *UserStorage {
	return &UserStorage{mapMu: sync.RWMutex{}, fileMu: sync.RWMutex{}, userStore: getStorage()}
}

func getStorage() *UserStore {
	f, err := ioutil.ReadFile(store)
	if err != nil {
		log.Fatalln("couldn't get storage", err)
	}

	s := UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		log.Fatalln(err)
	}
	return &s
}

func (u *UserStorage) updateStorage() error {
	u.fileMu.Lock()
	defer u.fileMu.Unlock()

	b, err := json.Marshal(u.userStore)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(store, b, fs.ModePerm)
}

func (u *UserStorage) GetAllUsers() models.UserList {
	u.mapMu.RLock()
	defer u.mapMu.RUnlock()

	list := make(models.UserList, len(u.userStore.List))
	for st, u := range u.userStore.List {
		list[st] = models.User{
			CreatedAt:   u.CreatedAt,
			DisplayName: u.DisplayName,
			Email:       u.Email,
		}
	}

	return list
}

func (u *UserStorage) CreateUser(name, email string) (string, error) {
	u.mapMu.Lock()
	defer u.mapMu.Unlock()
	u.userStore.Increment++
	id := strconv.Itoa(u.userStore.Increment)

	user := User{
		CreatedAt:   time.Now(),
		DisplayName: name,
		Email:       email,
	}

	u.userStore.List[id] = user

	return id, u.updateStorage()
}

func (u *UserStorage) GetUser(id string) (models.User, error) {
	u.mapMu.RLock()
	defer u.mapMu.RUnlock()
	if _, ok := u.userStore.List[id]; !ok {
		return models.User{}, errors.New(errs.ErrNotFound)
	}

	userFromStorage := u.userStore.List[id]

	return models.User{
		CreatedAt:   userFromStorage.CreatedAt,
		DisplayName: userFromStorage.DisplayName,
		Email:       userFromStorage.Email,
	}, nil
}

func (u *UserStorage) UpdateUser(id, newName string) error {
	u.mapMu.Lock()
	defer u.mapMu.Unlock()

	if _, ok := u.userStore.List[id]; !ok {
		return errors.New(errs.ErrNotFound)
	}
	user := u.userStore.List[id]
	user.DisplayName = newName
	u.userStore.List[id] = user

	return u.updateStorage()
}

func (u *UserStorage) DeleteUser(id string) error {
	if _, ok := u.userStore.List[id]; !ok {
		return errors.New(errs.ErrNotFound)
	}
	u.mapMu.Lock()
	defer u.mapMu.Unlock()
	delete(u.userStore.List, id)

	return u.updateStorage()
}
