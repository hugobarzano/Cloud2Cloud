package users

import (
	"app/config"
	"app/constans"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	ID       string `json:"_id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
}



func createUser(user *User) (*User, error) {
	config := config.Get()
	session, err := mgo.Dial(config.Mongodb.URI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)

	c := session.DB(constans.DATABASE).C("users")
	err = c.Insert(user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil
}

func getUser(username string) (*User, error) {
	config := config.Get()
	session, err := mgo.Dial(config.Mongodb.URI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)

	c := session.DB(constans.DATABASE).C("users")
	result := &User{}
	err = c.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

func deleteUser(username string) (*User, error) {

	config := config.Get()
	user, err := getUser(username)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	session, err := mgo.Dial(config.Mongodb.URI)
	c := session.DB(constans.DATABASE).C("users")
	err = c.RemoveId(user.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}
