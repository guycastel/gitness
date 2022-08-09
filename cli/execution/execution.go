// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package execution

import "gopkg.in/alecthomas/kingpin.v2"

// Register the command.
func Register(app *kingpin.Application) {
	cmd := app.Command("execution", "manage executions")
	registerFind(cmd)
	registerList(cmd)
	registerCreate(cmd)
	registerUpdate(cmd)
	registerDelete(cmd)
}
