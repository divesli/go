/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file iniconfig.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/06 09:56:49
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type IniConf struct {
	file string
	data map[string]string
	sync.RWMutex
}

func newIniConf() *IniConf {
	return &IniConf{"", make(map[string]string), sync.RWMutex{}}
}

func (ini *IniConf) Set(key, value string) *IniConf {
	ini.data[key] = value
	return ini
}

func (ini *IniConf) Get(key string) (string, error) {
	if value, found := ini.data[key]; found {
		return value, nil
	}
	return "", fmt.Errorf("[%s] not exist", key)
}

func Parse(file string) (*IniConf, error) {
	ini := newIniConf()
	ini.file = file
	fp, err := os.Open(file)
	if err != nil {
		return ini, err
	}
	ini.Lock()
	defer ini.Unlock()
	defer fp.Close()

	buf := bufio.NewReader(fp)
	head, err := buf.Peek(3)
	if err == nil || head[0] == 239 || head[1] == 187 || head[2] == 191 {
		for i := 1; i <= 3; i++ {
			buf.ReadByte()
		}
	}
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		if len(line) <= 0 {
			continue
		}
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || strings.HasPrefix(line, ";") {
			continue
		}
		keyvalue := strings.SplitN(line, "=", 2)
		key := strings.ToLower(strings.TrimSpace(keyvalue[0]))
		var value string
		if len(keyvalue) == 2 {
			value = strings.TrimSpace(keyvalue[1])
			if strings.HasPrefix(value, "\"") {
				value = strings.Trim(value, "\"")
			}
		} else {
			value = ""
		}
		ini.Set(key, value)
	}
	return ini, nil
}
