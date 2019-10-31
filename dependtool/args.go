// Copyright 2019 The UNICORE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file
//
// Author: Gaulthier Gain <gaulthier.gain@uliege.be>

package dependtool

import (
	. "github.com/akamensky/argparse"
	"os"
	u "tools/common"
)

const (
	PROGRAM     = "program"
	TEST_FILE   = "testFile"
	OPTIONS     = "options"
	WAIT_TIME   = "waitTime"
	SAVE_OUTPUT = "saveOutput"
	FULL_DEPS   = "fullDeps"
)

// parseLocalArguments parses arguments of the application.
func parseLocalArguments(p *Parser, args *u.Arguments) {

	args.InitArgParse(p, args, u.STRING, "p", PROGRAM,
		&Options{Required: true, Help: "Program name"})
	args.InitArgParse(p, args, u.STRING, "t", TEST_FILE,
		&Options{Required: false, Help: "Path of the test file"})
	args.InitArgParse(p, args, u.STRING, "o", OPTIONS,
		&Options{Required: false, Default: "", Help: "Extra options for " +
			"launching program"})

	args.InitArgParse(p, args, u.INT, "w", WAIT_TIME,
		&Options{Required: false, Default: 60, Help: "Time wait (" +
			"sec) for external tests (default: 60 sec)"})

	args.InitArgParse(p, args, u.BOOL, "", SAVE_OUTPUT,
		&Options{Required: false, Default: false,
			Help: "Save results as TXT file and graphs as PNG file"})
	args.InitArgParse(p, args, u.BOOL, "", FULL_DEPS,
		&Options{Required: false, Default: false, Help: "Show dependencies of" +
			"dependencies"})

	_ = p.Parse(os.Args)
}
