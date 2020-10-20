/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/Matrix-Zhang/fabric-gm/bccsp/factory"
	"github.com/Matrix-Zhang/fabric-gm/cmd/common"
	"github.com/Matrix-Zhang/fabric-gm/discovery/cmd"
)

func main() {
	factory.InitFactories(nil)
	cli := common.NewCLI("discover", "Command line client for fabric discovery service")
	discovery.AddCommands(cli)
	cli.Run(os.Args[1:])
}
