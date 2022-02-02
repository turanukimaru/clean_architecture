package dummies

import (
	"fmt"
	"github.com/turanukimaru/gormstart/pkg/dummydb"
)

type Dummy struct {
}

func (d *Dummy) Hello() (err error) {
	fmt.Println("Hello Dummies !")
	// お試しという事でDBアクセスだけ別モジュールにしてみる。
	// ローカルを参照するには
	//	go mod edit -replace github.com/turanukimaru/gormstart=../gormstart
	// なお、バージョンの都合で次のエラーが発生することがある。ので go get -u でバージョンを合わせる。
	//go: finding module for package github.com/turanukimaru/goastart/dummydb
	//github.com/turanukimaru/goastart/calcapi imports
	//	github.com/turanukimaru/goastart/dummydb: no matching versions for query "latest"
	// go get -u github.com/turanukimaru/gormstart
	err = dummydb.DbAccess()
	return
}

func (d *Dummy) Allow() (err error) {
	fmt.Println("Hello Dummies !")
	err = dummydb.DbAccess()
	return
}
