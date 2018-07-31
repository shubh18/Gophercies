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

//InitDB initialises Database and creates task bucket if not already present.
func InitDB(DBString string) (*bolt.DB, error) {
	var err error
	dbconnect, err = bolt.Open(DBString, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return dbconnect, dbconnect.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

}

//CreateTask creates task in boltDB
func CreateTask(task string) error {
	var taskID int
	err := dbconnect.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskID64, _ := b.NextSequence()
		taskID = int(taskID64)
		key := IntToByte(taskID)
		return b.Put(key, []byte(task))
	})

	return err

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

	return tasks, err
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
