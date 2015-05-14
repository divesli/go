/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file runtime.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/07 11:30:37
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var w sync.WaitGroup

func main() {
	fmt.Printf("start = %d\n", time.Now().Unix())
	var num int
	num = runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	for y := 0; y < num; y++ {
		w.Add(1)
		go func() {
			defer w.Done()
			var x, a, b, c, d, e, f, g int
			for i := 0; i < 100000000; i++ {
				x = x + i - (i - 1)
				a = a * i / (i + 1)
				b = b / (i + 1) * (i + 1)
				c = c / (i + 1) * (i + 1)
				d = d / (i + 1) * (i + 1)
				e = e / (i + 1) * (i + 1)
				f = f / (i + 1) * (i + 1)
				g = g / (i + 1) * (i + 1)
			}
		}()
	}
	w.Wait()
	fmt.Printf("end = %d\n", time.Now().Unix())
}
