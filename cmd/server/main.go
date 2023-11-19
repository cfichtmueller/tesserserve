// Copyright 2023 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/cfichtmueller/jug"
	"github.com/cfichtmueller/tesserserve/internal/api"
	"github.com/cfichtmueller/tesserserve/internal/util"
	"log"
	"os"
)

func main() {

	bindIp := util.Nvl(os.Getenv("BIND_IP"), "127.0.0.1")
	port := util.Nvl(os.Getenv("PORT"), "8000")
	router := jug.New()

	router.POST("api/recognize", api.Recognize)

	router.ExpandMethods()

	address := bindIp + ":" + port

	log.Printf("Server listening on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
