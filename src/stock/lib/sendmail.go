/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file sendmail.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/05/06 11:30:52
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package lib

import (
	"net/smtp"
	"strings"
)

func SendMail(user, pwd, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, pwd, hp[0])
	var conttype string
	if mailtype == "html" {
		conttype = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		conttype = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	message := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + conttype + "\r\n\r\n" + body)
	tos := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, tos, message)
	return err
}
