/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file main.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/04/15 16:18:06
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("test md5"))
	str := md5Ctx.Sum(nil)
	fmt.Println(str)
	fmt.Println(hex.EncodeToString(str))
}
