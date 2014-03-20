package buckets

import (
	"labix.org/v2/mgo"
)

var (
	// 默认的数据库地址
	defaultDbAdress = "localhost:27017"
	// 默认的数据库名称
	defaultDbName = "mis"
)

// Bucket接口定义了一个与数据库交互的对象所要实现的方法
type Bucket interface {
	Add(interface{}) error
}

// OriginBucket提供最基本的数据库交互的方法
type OriginBucket struct {
	dbAddress      string
	dbName         string
	collectionName string
}

// Add方法向数据库相应的collection中添加一条记录
// 参数能是一个实例或map
// 返回err不为空时表示添加失败
func (bucket *OriginBucket) Add(obj interface{}) (err error) {
	session, err := mgo.Dial(bucket.dbAddress)
	if err != nil {
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(bucket.dbName).C(bucket.collectionName)
	err = c.Insert(obj)
	if err != nil {
		return err
	}

	return nil
}

// GetAll方法向数据库查询相应collection的所有记录
// 第一个返回值为储存了这些记录的slice；第二个返回值err不为nil表示查找失败
func (bucket *OriginBucket) GetAll() (list []interface{}, err error) {
	list = make([]interface{}, 10)
	return list, nil
}

// newBucket方法能创建一个bucket
// 参数为collection名称
// 返回一个Bucket接口
func newBucket(collectionName string) (bucket Bucket) {
	bucket = &OriginBucket{
		dbAddress:      defaultDbAdress,
		dbName:         defaultDbName,
		collectionName: collectionName,
	}

	return bucket
}
