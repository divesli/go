/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file color.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/04 14:12:20
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package lib

import (
	"fmt"
)

const (
	TxtBlack = iota + 30
	TxtRed
	TxtGreen
	TxtYellow
	TxtBlue
	TxtMagenta
	TxtCyan
	TxtWhite
)

func txtColor(color int, str string) string {
	switch color {
	case TxtBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtBlack, str)
	case TxtRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtRed, str)
	case TxtGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtGreen, str)
	case TxtYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtYellow, str)
	case TxtBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtBlue, str)
	case TxtMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtMagenta, str)
	case TxtCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtCyan, str)
	case TxtWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TxtWhite, str)
	default:
		return str
	}
}

func Black(str string) string {
	return txtColor(TxtBlack, str)
}

func Red(str string) string {
	return txtColor(TxtRed, str)
}

func Green(str string) string {
	return txtColor(TxtGreen, str)
}

func Yellow(str string) string {
	return txtColor(TxtYellow, str)
}

func Blue(str string) string {
	return txtColor(TxtBlue, str)
}

func Magenta(str string) string {
	return txtColor(TxtMagenta, str)
}

func Cyan(str string) string {
	return txtColor(TxtCyan, str)
}

func White(str string) string {
	return txtColor(TxtWhite, str)
}
