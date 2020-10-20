/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"github.com/Matrix-Zhang/fabric-gm/core/chaincode/shim"
	"github.com/Matrix-Zhang/fabric-gm/examples/chaincode/go/example05"
)

func main() {
	err := shim.Start(new(example05.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
