package minerset

import (
	"app/config"
	"app/constans"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)



type Minerset struct {
	ID       string `json:"_id" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title,omitempty"`
	Owner    string `json:"owner" bson:"owner,omitempty"`
	Url		 string `json:"url" bson:"url,omitempty"`
	Date     time.Time `json:"date" bson:"date,omitempty"`
}



func createMinerset(minerset *Minerset) (*Minerset, error) {
	config := config.Get()
	session, err := mgo.Dial(config.Mongodb.URI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)

	c := session.DB(constans.DATABASE).C("minersets")
	err = c.Insert(minerset)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return minerset, nil
}

func getMinerset(title string) (*Minerset, error) {
	config := config.Get()
	session, err := mgo.Dial(config.Mongodb.URI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)

	c := session.DB(constans.DATABASE).C("minersets")
	result := &Minerset{}
	err = c.Find(bson.M{"title": title}).One(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

func deleteMinerset(title string) (*Minerset, error) {

	config := config.Get()
	minerset, err := getMinerset(title)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	session, err := mgo.Dial(config.Mongodb.URI)
	c := session.DB(constans.DATABASE).C("minersets")
	err = c.RemoveId(minerset.ID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return minerset, nil
}

func listMinersets() ([]Minerset, error) {
	config := config.Get()
	session, err := mgo.Dial(config.Mongodb.URI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//session.SetMode(mgo.Monotonic, true)

	c := session.DB(constans.DATABASE).C("minersets")
	result := []Minerset{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}
