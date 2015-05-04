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
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"stock/lib"
	"stock/request"
	"strconv"
	"strings"
	"sync"
	"time"
)

var w sync.WaitGroup
var f string

func main() {
	f = strings.Repeat("-", 100)
	for i := 0; i < 10; i++ {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		pftitle()
		list := []string{"0601169", "0601377", "0601318", "0600030", "0600036"}
		for _, id := range list {
			w.Add(1)
			go run(id)
		}
		w.Wait()
		time.Sleep(2000 * time.Millisecond)
	}
	//w.Wait()
}

func pftitle() {
	out := fmt.Sprintf("%10s %10s %10s %10s %10s %10s %10s", "名称", "当前价格", "涨跌幅", "最高价格", "最低价格", "成交量", "成交额")
	fmt.Println(out)
	fmt.Println(f)
}

func run(id string) {
	defer w.Done()
	url := "http://api.money.126.net/data/feed/" + id
	req := request.Get(url)
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
	for _, v := range data {
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
}
