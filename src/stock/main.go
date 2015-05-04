/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file main.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/04 10:30:34
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
