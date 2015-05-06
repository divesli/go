/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file run.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/04 10:54:50
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"stock/lib"
	"stock/request"
	"stock/sigl"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	et  = 1
	sec = 1000
	lp  = 30
	sp  = ","
)

var w sync.WaitGroup
var f string
var c chan uint

func main() {
	f = strings.Repeat("-", 100)
	var confile string
	if len(os.Args) > 1 {
		confile = os.Args[1]
	} else {
		confpath := getpath()
		confile = filepath.Join(confpath, "config.ini")
	}

	if _, err := os.Stat(confile); err != nil {
		fmt.Printf("%v\n", err)
	}

	ini, _ := lib.Parse(confile)
	ids, err := ini.Get("stock_code")
	if err != nil {
		ids = "0601377,0601318,0600030,0600036"
	}
	stsp, err := ini.Get("stock_sp")
	if err != nil {
		stsp = sp
	}
	var stlp int
	st_lp, err := ini.Get("stock_refresh")
	if err != nil {
		stlp = lp
	} else {
		stlp, _ = strconv.Atoi(st_lp)
	}

	stlp *= sec
	sl := sigl.NewSigl()
	sl.Register(syscall.SIGINT, siglhandler)
	sl.Register(syscall.SIGKILL, siglhandler)
	c = make(chan uint)
	go sl.Run()
	go run(ids, stsp, stlp)
	go getinput()
	<-c // 从chan 取值,取不到即阻塞等待只到取到值
}

func siglhandler(sg os.Signal) {
	c <- et
}

func getinput() {
	reader := bufio.NewReader(os.Stdin)
	char, err := reader.ReadString('\n')
	if err == io.EOF || char == "\n" {
		c <- et // 向chan写入 1
	}
}

func getpath() string {
	workpath, _ := os.Getwd()
	workpath, _ = filepath.Abs(workpath)
	confpath := filepath.Join(workpath, "config")
	return confpath
}

func pftitle() {
	out := fmt.Sprintf("%10s %10s %10s %10s %10s %10s %10s\n%s", "名称", "当前价格", "涨跌幅", "最高价格", "最低价格", "成交量", "成交额", f)
	fmt.Println(out)
}

func run(ids, stsp string, stlp int) {
	url := "http://api.money.126.net/data/feed/"
	slids := strings.Split(ids, stsp)
	for {
		// 清屏命令
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		pftitle()

		req := request.Get(url + ids)
		jsonstr, err := req.GetBody()
		defer req.Close()
		if err != nil {
			fmt.Println(err)
		}

		reg, err := regexp.Compile("\\((.*)\\)")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		match := reg.FindStringSubmatch(jsonstr)
		if len(match) != 2 {
			fmt.Println("Match failed")
			fmt.Println(match)
			os.Exit(1)
		}

		str := match[1]
		var data map[string]map[string]interface{}
		if err := json.Unmarshal([]byte(str), &data); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, id := range slids {
			if v, found := data[id]; found {
				w.Add(1)
				go spline(v)
			}
		}

		w.Wait()

		time.Sleep(time.Duration(stlp) * time.Millisecond)
	}
}

func spline(v map[string]interface{}) {
	defer w.Done()
	name := fmt.Sprintf("%s", v["name"])
	price := fmt.Sprintf("%.2f", v["price"])
	per, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", v["percent"]), 64)
	perc := per * 100
	percent := lib.Red(fmt.Sprintf("%.2f%%", perc))
	high := fmt.Sprintf("%.2f", v["high"])
	low := fmt.Sprintf("%.2f", v["low"])
	vol, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", v["volume"]), 64)
	volume := fmt.Sprintf("%.2f万手", vol/1000000)
	turn, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", v["turnover"]), 64)
	turnover := fmt.Sprintf("%.2f亿", turn/100000000)
	out := fmt.Sprintf("%10s %10s %26s %13s %13s %16s %10s\n%s", name, price, percent, high, low, volume, turnover, f)
	fmt.Println(out)
}
