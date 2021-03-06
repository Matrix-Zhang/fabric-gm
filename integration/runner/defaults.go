/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package runner

import (
	"time"

	"github.com/Matrix-Zhang/fabric-gm/integration/helpers"
)

const DefaultStartTimeout = 30 * time.Second

// DefaultNamer is the default naming function.
var DefaultNamer NameFunc = helpers.UniqueName

// A NameFunc is used to generate container names.
type NameFunc func() string
