/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file dbbase.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/08 11:31:13
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package dbbase

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
	"strings"
)

type dbBase struct {
	dns  string
	conn *sql.DB
	tx   *sql.Tx
}

func NewdbBase() *dbBase {
	return &dbBase{"", nil, nil}
}

func (db *dbBase) SetConfig(config map[string]string) *dbBase {
	host := config["db_host"]
	port := config["db_port"]
	user := config["db_user"]
	pwd := config["db_pwd"]
	name := config["db_name"]
	charset := config["db_charset"]
	db.dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pwd, host, port, name, charset)
	return db
}

func (db *dbBase) open() (*sql.DB, error) {
	if db.conn != nil {
		return db.conn, nil
	}
	if len(db.dns) <= 0 {
		return nil, fmt.Errorf("Database's dns failed")
	}
	dbConn, err := sql.Open("mysql", db.dns)
	panicerr(err)
	db.conn = dbConn
	return db.conn, nil
}

func (db *dbBase) Close() {
	if db.conn != nil {
		db.conn.Close()
		db.conn = nil
	}
}

func (db *dbBase) Query(sqlstr string) (res []map[string]string) {
	conn, err := db.open()
	panicerr(err)
	var rows *sql.Rows
	if db.tx != nil {
		rows, err = db.tx.Query(sqlstr)
	} else {
		rows, err = conn.Query(sqlstr)
	}
	panicerr(err)

	defer rows.Close()
	columns, err := rows.Columns()
	panicerr(err)

	colnum := len(columns)
	values := make([]sql.RawBytes, colnum)
	refs := make([]interface{}, colnum)

	for i, _ := range values {
		refs[i] = &values[i]
	}

	res = []map[string]string{}

	for rows.Next() {
		err := rows.Scan(refs...)
		panicerr(err)

		kv := make(map[string]string, colnum)
		for i, v := range values {
			val := string(v)
			kv[columns[i]] = val
		}
		res = append(res, kv)
	}
	return res
}

func (db *dbBase) Begin() *dbBase {
	conn, err := db.open()
	panicerr(err)
	db.tx, err = conn.Begin()
	panicerr(err)
	return db
}

func (db *dbBase) Commit() {
	if db.tx != nil {
		err := db.tx.Commit()
		db.tx = nil
		panicerr(err)
	}
}

func (db *dbBase) Rollback() {
	if db.tx != nil {
		err := db.tx.Rollback()
		db.tx = nil
		panicerr(err)
	}
}
func (db *dbBase) str2lower(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return str
}
func (db *dbBase) Insert(sqlstr string) (lastId int64) {
	sqlstr = db.str2lower(sqlstr)
	if !strings.HasPrefix(sqlstr, "insert") {
		panicerr(fmt.Errorf("It's not insert sql,[%s]", sqlstr))
	}
	res := db.Exec(sqlstr)
	lastInsertId, err := res.LastInsertId()
	panicerr(err)
	return lastInsertId
}

func (db *dbBase) Delete(sqlstr string) (n int64) {
	sqlstr = db.str2lower(sqlstr)
	if !strings.HasPrefix(sqlstr, "delete") {
		panicerr(fmt.Errorf("It's not delete sql,[%s]", sqlstr))
	}
	res := db.Exec(sqlstr)
	affect, err := res.RowsAffected()
	panicerr(err)
	return affect
}

func (db *dbBase) Update(sqlstr string) (n int64) {
	sqlstr = db.str2lower(sqlstr)
	if !strings.HasPrefix(sqlstr, "update") {
		panicerr(fmt.Errorf("It's not update sql,[%s]", sqlstr))
	}
	res := db.Exec(sqlstr)
	affect, err := res.RowsAffected()
	panicerr(err)
	return affect
}

func (db *dbBase) Exec(sqlstr string) (res sql.Result) {
	conn, err := db.open()
	panicerr(err)

	stmt, err := conn.Prepare(sqlstr)
	panicerr(err)
	if db.tx != nil {
		res, err = db.tx.Stmt(stmt).Exec()
	} else {
		res, err = stmt.Exec()
	}
	panicerr(err)

	return res
}

func panicerr(err error) {
	if err != nil {
		panic(err)
	}
}
