package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var dbconnect *bolt.DB

//Task struct
type Task struct {
	Key   int
	Value string
}

//InitDB initialises Database
func InitDB(DBString string) error {
	var err error
	dbconnect, err = bolt.Open(DBString, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return dbconnect.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

}

//CreateTask creates task in boltDB
func CreateTask(task string) (int, error) {
	var taskID int
	err := dbconnect.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskID64, _ := b.NextSequence()
		taskID = int(taskID64)
		key := IntToByte(taskID)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return 0, nil

}

//GetAllTasks returns all the tasks added
func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := dbconnect.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   ByteToInt(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//DeleteTasks delete tasks by id
func DeleteTasks(key int) error {
	return dbconnect.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(IntToByte(key))
	})
}

//IntToByte converts integer into byte
func IntToByte(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

//ByteToInt converts bytes into integer
func ByteToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
