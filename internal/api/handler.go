// Copyright 2023 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/cfichtmueller/jug"
	"github.com/cfichtmueller/tesserserve/internal/ocr"
	"log"
)

func Recognize(c jug.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Printf(err.Error())
		c.HandleError(err)
		return
	}
	request := ocr.Request{
		Lang:          c.Query("lang"),
		IteratorLevel: c.Query("iterator-level"),
		Data:          data,
	}
	if err := request.Validate(); err != nil {
		c.RespondBadRequestE(err)
		return
	}

	res, err := ocr.Recognize(request)
	if err != nil {
		log.Printf(err.Error())
		c.HandleError(err)
		return
	}

	c.RespondOk(res)
}
