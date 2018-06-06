package dbclient

import (
	"github.com/callistaenterprise/goblog/accountservice/model"
	"log"
	"github.com/boltdb/bolt"
	"strconv"
	"encoding/json"
)

//从这里如下的struct和inteface的定义要掌握面向对象的编程思想

//这里定义接口的目的，是要将所有创建的db clien都归为一个接口
//这里针对database的操作方法来创建一个接口，接口中定义了databases的可操作方法
type IBoltClient interface {
	OpenBoltDb()  //创建一个database文件
	QueryAccount(accountId string) (model.Account, error)
	Seed() //在文件中添加数据
}

//这里创建一个bolt.DB类型的结构体，该结构体是一个连接database的实例，注意上边的接口是这个实例的通用类型
type BoltClient struct {
	boltDB *bolt.DB
}

//这里定义的是上边的实例的方法，这个方法用*指针的目的是，该方法会对该实例中的数据进行修改，所以这里要用到指针，否则无法进行结构体中数据的修改
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil) //调用open方法得到一个bolt.DB这样一个实例，该实例方法是接下来操作的基础
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

//该实例函数在得到了一个初始化的db实例而后，创建一个database仓库
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			log.Fatalln("create bucket failed: ", err)
		}
		return nil
	})
}

//针对上边创建的仓库，往仓库中传入数据
func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i<total; i++ {
		key := strconv.Itoa(10000 + i)
		acc := model.Account {
			Id:key,
			Name: "Person_" + strconv.Itoa(i),
		}

		jsonBytes,_ := json.Marshal(acc)

		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})

	}

	log.Printf("Seed %v fake accounts....\n", total)
}

func (bc *BoltClient) QueryAccount(accountId string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error { //该bolt类型数据库在进行查询操作的时候需要在view函数中进行
		b := tx.Bucket([]byte("AccountBucket"))

		accountBytes := b.Get([]byte(accountId)) //从数据库中查询数据
		if accountBytes == nil {
			log.Fatalln("No account found for ", accountId)
		}

		json.Unmarshal(accountBytes, &account) //将查询到的数据进行json反解析到一个结构体中
		return nil //这里想err 变量传递nil
	})

	if err != nil {
		return model.Account{}, err
	}
	return account, nil  //这里返回解析出的数据
}

