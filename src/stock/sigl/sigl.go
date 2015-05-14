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
	//	"syscall"
)

type signalHandler func(sg os.Signal)

type Sigl struct {
	m map[os.Signal]signalHandler
}

func NewSigl() *Sigl {
	return &Sigl{m: make(map[os.Signal]signalHandler)}
}

func (s *Sigl) Register(handler signalHandler, sigs ...os.Signal) *Sigl {
	for _, sig := range sigs {
		if _, found := s.m[sig]; !found {
			s.m[sig] = handler
		}
	}
	return s
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
	var sigs []os.Signal
	for sig, _ := range s.m {
		sigs = append(sigs, sig)
	}
	go func() {
		//signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL)
		signal.Notify(ch, sigs...)
		//signal.Notify(ch)
		sg := <-ch
		s.handler(sg)
	}()
}
