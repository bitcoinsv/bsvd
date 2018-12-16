// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpctest

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	// compileMtx guards access to the executable path so that the project is
	// only compiled once.
	compileMtx sync.Mutex

	// executablePath is the path to the compiled executable. This is the empty
	// string until bsvd is compiled. This should not be accessed directly;
	// instead use the function bsvdExecutablePath().
	executablePath string
)

// bsvdExecutablePath returns a path to the bsvd executable to be used by
// rpctests. To ensure the code tests against the most up-to-date version of
// bsvd, this method compiles bsvd the first time it is called. After that, the
// generated binary is used for subsequent test harnesses. The executable file
// is not cleaned up, but since it lives at a static path in a temp directory,
// it is not a big deal.
func bsvdExecutablePath() (string, error) {
	compileMtx.Lock()
	defer compileMtx.Unlock()

	// If bsvd has already been compiled, just use that.
	if len(executablePath) != 0 {
		return executablePath, nil
	}

	testDir, err := baseDir()
	if err != nil {
		return "", err
	}

	// Determine import path of this package. Not necessarily bitcoinsv/bsvd if
	// this is a forked repo.
	_, rpctestDir, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("Cannot get path to bsvd source code")
	}
	bsvdPkgPath := filepath.Join(rpctestDir, "..", "..", "..")
	bsvdPkg, err := build.ImportDir(bsvdPkgPath, build.FindOnly)
	if err != nil {
		return "", fmt.Errorf("Failed to build bsvd: %v", err)
	}

	// Build bsvd and output an executable in a static temp path.
	outputPath := filepath.Join(testDir, "bsvd")
	if runtime.GOOS == "windows" {
		outputPath += ".exe"
	}
	cmd := exec.Command("go", "build", "-o", outputPath, bsvdPkg.ImportPath)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Failed to build bsvd: %v", err)
	}

	// Save executable path so future calls do not recompile.
	executablePath = outputPath
	return executablePath, nil
}
