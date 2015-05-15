/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file goruntimemanage.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/15 15:15:19
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package goruntimemanage

type GoRtManage struct {
	ncap uint
	gmc  chan uint
}

const (
	minDefaultNum = 1
)

func NewGoRtManage(n uint) *GoRtManage {
	if n <= 0 {
		n = minDefaultNum
	}
	return &GoRtManage{n, make(chan uint, n)}
}

func (gm *GoRtManage) Push() {
	gm.gmc <- 1
}

func (gm *GoRtManage) Pop() {
	<-gm.gmc
}

func (gm *GoRtManage) Total() uint {
	return uint(len(gm.gmc))
}

func (gm *GoRtManage) Left() uint {
	return gm.ncap - uint(len(gm.gmc))
}

func (gm *GoRtManage) Close() {
	if gm.gmc != nil {
		close(gm.gmc)
		gm.ncap = 0
	}
}
