package dao

import (
	"log"
	"time"

	. "movie-rest-api/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	username   = "rully"
	password   = "secret"
	collection = "movies"
)

// Establish a connection to database
func (m *MoviesDAO) Connect() {
	// session, err := mgo.Dial(m.Server)

	info := &mgo.DialInfo{
		Addrs:    []string{m.Server},
		Timeout:  60 * time.Second,
		Database: m.Database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(collection).Find(bson.M{}).All(&movies)
	return movies, err
}

// Find a movie by its id
func (m *MoviesDAO) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie Movie) error {
	err := db.C(collection).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie Movie) error {
	err := db.C(collection).Remove(&movie)
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie Movie) error {
	err := db.C(collection).UpdateId(movie.ID, &movie)
	return err
}
