package db

import (
	"fmt"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestInitDBNegative(t *testing.T) {
	_, err := InitDB("/")
	if err == nil {
		t.Error("Expected err got,", nil)
	}

}
func TestInitDB(t *testing.T) {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "cmd.db")
	_, err := InitDB(DbPath)
	if err != nil {
		t.Error("Expected nil got,", err)
	}

}

func TestCreateTask(t *testing.T) {

	err := CreateTask("Work on Golang")
	if err != nil {
		t.Error("Expected nil got", err)
	}
}

func TestGetAllTasks(t *testing.T) {
	tasks, _ := GetAllTasks()
	if tasks == nil {
		t.Error("Expected tasks got,", tasks)
	}
}

func TestDeleteTasks(t *testing.T) {
	err := DeleteTasks(1)
	fmt.Println(err)
	if err != nil {
		t.Error("Expected nil got,", err)
	}
}

func TestIntToByte(t *testing.T) {
	bytes := IntToByte(10)
	fmt.Println(bytes)
	if bytes == nil {
		t.Error("Expected nil got", bytes)

	}
}

func TestByteToInt(t *testing.T) {
	var s []byte
	s = make([]byte, 8, 12)
	s = []byte{0, 0, 0, 0, 0, 0, 0, 10}
	val := ByteToInt(s)
	if val == 0 {
		t.Error("Expected int got", val)
	}
}

/*
func TestAddTask(t *testing.T) {
	DbPath := "/home/gslab/tasks.db"
	db, _ := InitDB(DbPath)
	err := CreateTask("testing123")
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestInit(t *testing.T) {
	DbPath := "/home/gslab/tasks.db"
	db, err := InitDB(DbPath)
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestInitNegative(t *testing.T) {
	DbPath := "/home/gslab123/tasks.db"
	_, err := InitDB(DbPath)
	if err == nil {
		t.Errorf("Expected result error, But got NO Error")
	}
}

func TestListTasks(t *testing.T) {
	DbPath := "/home/gslab/tasks.db"
	db, _ := InitDB(DbPath)
	_, err := GetAllTasks()
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestDeleteTask(t *testing.T) {
	DbPath := "/home/gslab/tasks.db"
	db, _ := InitDB(DbPath)
	err := DeleteTasks(10)
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}
*/
