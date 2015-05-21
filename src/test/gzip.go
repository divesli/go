/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file gzip.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/21 11:55:23
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	//tarZip()
	unTarZip()
	//zip()
	//unzip()
}
func unzip() {
	fr, err := os.Open("test.zip")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	fw, err := os.Create(gr.Name)
	if err != nil {
		panic(err)
	}
	ln, err := io.Copy(fw, gr)
	if err != nil {
		panic(err)
	}
	fmt.Println("unzip file success,", ln)
}

func zip() {
	fw, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	gw.Name = "log"
	fr, err := os.Open("log.txt")

	if err != nil {
		panic(err)
	}
	defer fr.Close()

	l, err := io.Copy(gw, fr)
	if err != nil {
		panic(err)
	}

	fmt.Println("zip file success,", l)
}

func unTarZip() {
	// file read
	fr, err := os.Open("test.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	// tar read
	tr := tar.NewReader(gr)

	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// 显示文件
		fmt.Println(h.Name)

		// 打开文件
		fw, err := os.OpenFile("test/"+h.Name, os.O_CREATE|os.O_WRONLY, 0755 /*os.FileMode(h.Mode)*/)
		//fw, err := os.OpenFile("test/"+h.Name, os.O_CREATE|os.O_RDWR, 0755 /*os.FileMode(h.Mode)*/)
		if err != nil {
			panic(err)
		}
		defer fw.Close()

		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("un tar.gz ok")
}

func tarZip() {
	fw, err := os.Create("test.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	dir, err := os.Open("class/")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		//fr, err := os.Open(dir.Name() + "/" + fi.Name())
		fr, err := os.Open(dir.Name() + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		fmt.Println(fr.Name())

		h := new(tar.Header)
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()

		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("tar gzip file success")
}
