/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file db.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/04/14 10:14:49
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package db

import (
	"database/sql"
	//"errors"
	_ "go-sql-driver/mysql"
	//"reflect"
	//"strconv"
	"fmt"
)

type dbbase struct {
	dns string
	db  *sql.DB
}

//var instance *dbbase = nil

func Newdbbase() *dbbase {
	//if instance == nil {
	return &dbbase{dns: "", db: nil}
	//}
	//return instance
}

func (d *dbbase) initDns() {
	config, _ := conf.NewConf("ini", "ini.conf")
	host := config.GetString("db_host")
	port := config.GetString("db_port")
	user := config.GetString("db_user")
	pwd := config.GetString("db_pass")
	name := config.GetString("db_name")
	charset := config.GetString("db_charset")
	d.dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pwd, host, port, name, charset)
}

func (d *dbbase) open() (*sql.DB, error) {
	if d.db != nil {
		return d.db, nil
	}
	if len(d.dns) <= 0 {
		d.initDns()
	}
	db, err := sql.Open("mysql", d.dns)
	if err != nil {
		return nil, err
	}
	d.db = db
	return d.db, nil
}

func (d *dbbase) Close() {
	db, _ := d.open()
	if db != nil {
		db.Close()
	}
}

func (d *dbbase) Query(s string) (rst []map[string]string, err error) {
	db, err := d.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(s)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(columns))
	refs := make([]interface{}, len(columns))
	for i, _ := range values {
		refs[i] = &values[i]
	}
	rst = []map[string]string{}
	for rows.Next() {
		if err := rows.Scan(refs...); err != nil {
			return nil, err
		}
		kv := make(map[string]string, len(columns))
		for i, v := range values {
			val := string(v)
			kv[columns[i]] = val
		}
		rst = append(rst, kv)
	}
	return rst, nil
}
