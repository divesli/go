/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file reflect.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/04/15 11:01:23
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type MyStruct struct {
	name string
	age  int
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {
	//	s := "this is string"
	//	fmt.Println(reflect.TypeOf(s))
	//	fmt.Println("-------------------")
	//
	//	fmt.Println(reflect.ValueOf(s))
	//	var x float64 = 3.4
	//	fmt.Println(reflect.ValueOf(x))
	//	fmt.Println("-------------------")
	//
	//	a := new(MyStruct)
	//	a.name = "yejianfeng"
	//	typ := reflect.TypeOf(a)
	//
	//	fmt.Println(typ.NumMethod())
	//	fmt.Println("-------------------")

	//b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	//	a := new(MyStruct)
	//	a.name = "ddd"
	//	a.age = 100
	//	b := reflect.ValueOf(a)
	//	fmt.Printf("%v\n", b)
	//	fmt.Printf("%v\n", b.Kind())
	//	ind := reflect.Indirect(b)
	//	fmt.Printf("%v\n", ind.NumField())
	//	fmt.Printf("%v\n", ind.Type().Field(1).Name)
	//	fmt.Printf("%v\n", ind.Type().Field(1).Interface())
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("%v\n", mem)
	var mystruct interface{}
	mystruct = new(MyStruct)
	ty := reflect.TypeOf(mystruct)
	fmt.Println(ty)
	switch f := mystruct.(type) {
	default:
		fmt.Printf("%v\n", f)
		fmt.Printf("%v\n", mystruct)
	}
	runtime.ReadMemStats(&mem)
	fmt.Printf("%v\n", mem)

}
