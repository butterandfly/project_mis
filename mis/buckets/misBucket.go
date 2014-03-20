package buckets

import (
	// "labix.org/v2/mgo"
	"../beers"
)

// const MisResourceUrl "/mis"

// AddMis能向数据库添加一个Mis。
// 返回err不为nil的时候，即添加失败。
func (misBucket *MisBucket) AddMis(mis *beers.Mis) (err error) {
	return misBucket.Add(mis)
}

type MisBucket struct {
	Bucket
}

func newMisBucket() (misBucket *MisBucket) {
	misBucket = &MisBucket{
		newBucket("mis"),
	}

	return misBucket
}

var SharedMisBucket = newMisBucket()

/*
func GetAMis() (mis *Mis, err error) {
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    return nil, err
  }
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)

  c := session.DB("test").C("mis")
  err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
  &Person{"Cla", "+55 53 8402 8510"})
  if err != nil {
  panic(err)
  }

  result := Person{}
  err = c.Find(bson.M{"name": "Ale"}).One(&result)
  if err != nil {
  panic(err)
  }
}  */
