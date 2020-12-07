package db

import (
	"log"
	"testing"
)

func TestNewTables(t *testing.T) {
	Newdb()
	//NewTables()
}

func TestQueryRate(t *testing.T) {
	Newdb()
	//数据入库
	log.Println(QueryRate("00ff25000010", "051a7000001"))
}

func TestChedckyssjDataUpdate(t *testing.T) {
	Newdb()

	log.Println(ChedckyssjDataUpdate("320102000111032020042916134000000386"))
}
