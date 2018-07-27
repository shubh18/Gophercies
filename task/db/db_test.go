package db

import (
	"fmt"
	"testing"
)

func TestInitDBNegative(t *testing.T) {
	_, err := InitDB("/home/cmd.db")
	if err == nil {
		t.Error("Expected err got,", nil)
	}

}
func TestInitDB(t *testing.T) {
	_, err := InitDB("/home/gslab/cmd.db")
	if err != nil {
		t.Error("Expected nil got,", err)
	}

}

func TestCreateTask(t *testing.T) {
	var i int
	i, _ = CreateTask("Work on Golang")
	if i != 0 {
		t.Error("Expected 0,nil got", i)
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
