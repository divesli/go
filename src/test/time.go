/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file time.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/13 10:43:29
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//	nt := time.NewTicker(time.Minute)
	//	for t := range nt.C {
	//		fmt.Println(t.Format("2006-01-02 15:04:05"))
	//	}
	now := time.Now()
	fmt.Printf("%v\n", now)
	// 计算下一个零点
	next := now.Add(time.Hour * 24)
	fmt.Printf("%v\n", next)
	next = time.Date(next.Year(), next.Month(), next.Day(), 10, 20, 30, 0, next.Location())
	fmt.Printf("%v\n", next)
	fmt.Printf("%v\n", next.Sub(now))
	//t := time.NewTimer(next.Sub(now))
	//<-t.C
	//fmt.Printf("%v\n", t)
}
