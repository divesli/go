/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file chan.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/04/16 17:05:53
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"fmt"
	"sync"
)

var c chan uint
var w sync.WaitGroup

func say(i int) {
	defer func() { <-c }()
	defer w.Done()
	fmt.Printf("This i=%d goruntime\n", i)
}

func main() {
	c = make(chan uint, 2)
	for i := 1; i <= 10; i++ {
		w.Add(1)
		fmt.Printf("want Put chan %d\n", i)
		c <- 1
		fmt.Printf("Put chan %d\n", i)
		go say(i)
		fmt.Printf("go call say %d\n", i)
	}
	w.Wait()
}
