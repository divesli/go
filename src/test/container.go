/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file container.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/07/24 17:25:04
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"container/list"
	"fmt"
)

func main() {
	lister := list.New()
	lister.Init()
	lister.PushBack(1)
	lister.PushBack(2)
	lister.PushBack(3)
	lister.PushBack(4)
	lister.PushBack(5)

	// has bug, not delete all elements
	fmt.Println("Before remove list:")
	for e := lister.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := lister.Front(); e != nil; e = e.Next() {
		lister.Remove(e)
	}
	fmt.Println("After remove list:")
	for e := lister.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("==== Fixed bug ====")
	fmt.Println("Before remove list:")
	var n *list.Element
	for e := lister.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := lister.Front(); e != nil; e = n {
		n = e.Next()
		lister.Remove(e)
	}
	fmt.Println("After remove list:")
	for e := lister.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
