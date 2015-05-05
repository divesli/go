/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file sigl.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/05 10:43:41
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package sigl

import (
	"os"
	"os/signal"
	"syscall"
)

type signalHandler func(sg os.Signal)

type Sigl struct {
	m map[os.Signal]signalHandler
}

func NewSigl() *Sigl {
	return &Sigl{m: make(map[os.Signal]signalHandler)}
}

func (s *Sigl) Register(sg os.Signal, handler signalHandler) {
	if _, found := s.m[sg]; !found {
		s.m[sg] = handler
	}
}

func (s *Sigl) handler(sg os.Signal) (err error) {
	if _, found := s.m[sg]; found {
		s.m[sg](sg)
		return nil
	}
	return nil
}

func (s *Sigl) Run() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL)
	//signal.Notify(ch)
	sg := <-ch
	s.handler(sg)
}
