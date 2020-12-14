// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Toolstash provides a way to save, run, and restore a known good copy of the Go toolchain
// and to compare the object files generated by two toolchains.
//
// Usage:
//
//	toolstash [-n] [-v] save [tool...]
//	toolstash [-n] [-v] restore [tool...]
//	toolstash [-n] [-v] [-t] go run x.go
//	toolstash [-n] [-v] [-t] [-cmp] compile x.go
//
// The toolstash command manages a ``stashed'' copy of the Go toolchain
// kept in $GOROOT/pkg/toolstash. In this case, the toolchain means the
// tools available with the 'go tool' command as well as the go, godoc, and gofmt
// binaries.
//
// The command ``toolstash save'', typically run when the toolchain is known to be working,
// copies the toolchain from its installed location to the toolstash directory.
// Its inverse, ``toolchain restore'', typically run when the toolchain is known to be broken,
// copies the toolchain from the toolstash directory back to the installed locations.
// If additional arguments are given, the save or restore applies only to the named tools.
// Otherwise, it applies to all tools.
//
// Otherwise, toolstash's arguments should be a command line beginning with the
// name of a toolchain binary, which may be a short name like compile or a complete path
// to an installed binary. Toolstash runs the command line using the stashed
// copy of the binary instead of the installed one.
//
// The -n flag causes toolstash to print the commands that would be executed
// but not execute them. The combination -n -cmp shows the two commands
// that would be compared and then exits successfully. A real -cmp run might
// run additional commands for diagnosis of an output mismatch.
//
// The -v flag causes toolstash to print the commands being executed.
//
// The -t flag causes toolstash to print the time elapsed during while the
// command ran.
//
// Comparing
//
// The -cmp flag causes toolstash to run both the installed and the stashed
// copy of an assembler or compiler and check that they produce identical
// object files. If not, toolstash reports the mismatch and exits with a failure status.
// As part of reporting the mismatch, toolstash reinvokes the command with
// the -S flag and identifies the first divergence in the assembly output.
// If the command is a Go compiler, toolstash also determines whether the
// difference is triggered by optimization passes.
// On failure, toolstash leaves additional information in files named
// similarly to the default output file. If the compilation would normally
// produce a file x.6, the output from the stashed tool is left in x.6.stash
// and the debugging traces are left in x.6.log and x.6.stash.log.
//
// The -cmp flag is a no-op when the command line is not invoking an
// assembler or compiler.
//
// For example, when working on code cleanup that should not affect
// compiler output, toolstash can be used to compare the old and new
// compiler output:
//
//	toolstash save
//	<edit compiler sources>
//	go tool dist install cmd/compile # install compiler only
//	toolstash -cmp compile x.go
//
// Go Command Integration
//
// The go command accepts a -toolexec flag that specifies a program
// to use to run the build tools.
//
// To build with the stashed tools:
//
//	go build -toolexec toolstash x.go
//
// To build with the stashed go command and the stashed tools:
//
//	toolstash go build -toolexec toolstash x.go
//
// To verify that code cleanup in the compilers does not make any
// changes to the objects being generated for the entire tree:
//
//	# Build working tree and save tools.
//	./make.bash
//	toolstash save
//
//	<edit compiler sources>
//
//	# Install new tools, but do not rebuild the rest of tree,
//	# since the compilers might generate buggy code.
//	go tool dist install cmd/compile
//
//	# Check that new tools behave identically to saved tools.
//	go build -toolexec 'toolstash -cmp' -a std
//
//	# If not, restore, in order to keep working on Go code.
//	toolstash restore
//
// Version Skew
//
// The Go tools write the current Go version to object files, and (outside
// release branches) that version includes the hash and time stamp
// of the most recent Git commit. Functionally equivalent
// compilers built at different Git versions may produce object files that
// differ only in the recorded version. Toolstash ignores version mismatches
// when comparing object files, but the standard tools will refuse to compile
// or link together packages with different object versions.
//
// For the full build in the final example above to work, both the stashed
// and the installed tools must use the same version string.
// One way to ensure this is not to commit any of the changes being
// tested, so that the Git HEAD hash is the same for both builds.
// A more robust way to force the tools to have the same version string
// is to write a $GOROOT/VERSION file, which overrides the Git-based version
// computation:
//
//	echo devel >$GOROOT/VERSION
//
// The version can be arbitrary text, but to pass all.bash's API check, it must
// contain the substring ``devel''. The VERSION file must be created before
// building either version of the toolchain.
//
package main // import "github.com/system-pclub/GCatch/GFix/dispatcher/tools/cmd/toolstash"

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var usageMessage = `usage: toolstash [-n] [-v] [-cmp] command line

Examples:
	toolstash save
	toolstash restore
	toolstash go run x.go
	toolstash compile x.go
	toolstash -cmp compile x.go

For details, godoc github.com/system-pclub/GCatch/GFix/dispatcher/tools/cmd/toolstash
`

func usage() {
	fmt.Fprint(os.Stderr, usageMessage)
	os.Exit(2)
}

var (
	goCmd   = flag.String("go", "go", "path to \"go\" command")
	norun   = flag.Bool("n", false, "print but do not run commands")
	verbose = flag.Bool("v", false, "print commands being run")
	cmp     = flag.Bool("cmp", false, "compare tool object files")
	timing  = flag.Bool("t", false, "print time commands take")
)

var (
	cmd       []string
	tool      string // name of tool: "go", "compile", etc
	toolStash string // path to stashed tool

	goroot   string
	toolDir  string
	stashDir string
	binDir   string
)

func canCmp(name string, args []string) bool {
	switch name {
	case "asm", "compile", "link":
		if len(args) == 1 && (args[0] == "-V" || strings.HasPrefix(args[0], "-V=")) {
			// cmd/go uses "compile -V=full" to query the tool's build ID.
			return false
		}
		return true
	}
	return len(name) == 2 && '0' <= name[0] && name[0] <= '9' && (name[1] == 'a' || name[1] == 'g' || name[1] == 'l')
}

var binTools = []string{"go", "godoc", "gofmt"}

func isBinTool(name string) bool {
	return strings.HasPrefix(name, "go")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("toolstash: ")

	flag.Usage = usage
	flag.Parse()
	cmd = flag.Args()

	if len(cmd) < 1 {
		usage()
	}

	s, err := exec.Command(*goCmd, "env", "GOROOT").CombinedOutput()
	if err != nil {
		log.Fatalf("%s env GOROOT: %v", *goCmd, err)
	}
	goroot = strings.TrimSpace(string(s))
	toolDir = filepath.Join(goroot, fmt.Sprintf("pkg/tool/%s_%s", runtime.GOOS, runtime.GOARCH))
	stashDir = filepath.Join(goroot, "pkg/toolstash")

	binDir = os.Getenv("GOBIN")
	if binDir == "" {
		binDir = filepath.Join(goroot, "bin")
	}

	switch cmd[0] {
	case "save":
		save()
		return

	case "restore":
		restore()
		return
	}

	tool = cmd[0]
	if i := strings.LastIndexAny(tool, `/\`); i >= 0 {
		tool = tool[i+1:]
	}

	if !strings.HasPrefix(tool, "a.out") {
		toolStash = filepath.Join(stashDir, tool)
		if _, err := os.Stat(toolStash); err != nil {
			log.Print(err)
			os.Exit(2)
		}

		if *cmp && canCmp(tool, cmd[1:]) {
			compareTool()
			return
		}
		cmd[0] = toolStash
	}

	if *norun {
		fmt.Printf("%s\n", strings.Join(cmd, " "))
		return
	}
	if *verbose {
		log.Print(strings.Join(cmd, " "))
	}
	xcmd := exec.Command(cmd[0], cmd[1:]...)
	xcmd.Stdin = os.Stdin
	xcmd.Stdout = os.Stdout
	xcmd.Stderr = os.Stderr
	err = xcmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func compareTool() {
	if !strings.Contains(cmd[0], "/") && !strings.Contains(cmd[0], `\`) {
		cmd[0] = filepath.Join(toolDir, tool)
	}

	outfile, ok := cmpRun(false, cmd)
	if ok {
		os.Remove(outfile + ".stash")
		return
	}

	extra := "-S"
	switch {
	default:
		log.Fatalf("unknown tool %s", tool)

	case tool == "compile" || strings.HasSuffix(tool, "g"): // compiler
		useDashN := true
		dashcIndex := -1
		for i, s := range cmd {
			if s == "-+" {
				// Compiling runtime. Don't use -N.
				useDashN = false
			}
			if strings.HasPrefix(s, "-c=") {
				dashcIndex = i
			}
		}
		cmdN := injectflags(cmd, nil, useDashN)
		_, ok := cmpRun(false, cmdN)
		if !ok {
			if useDashN {
				log.Printf("compiler output differs, with optimizers disabled (-N)")
			} else {
				log.Printf("compiler output differs")
			}
			if dashcIndex >= 0 {
				cmd[dashcIndex] = "-c=1"
			}
			cmd = injectflags(cmd, []string{"-v", "-m=2"}, useDashN)
			break
		}
		if dashcIndex >= 0 {
			cmd[dashcIndex] = "-c=1"
		}
		cmd = injectflags(cmd, []string{"-v", "-m=2"}, false)
		log.Printf("compiler output differs, only with optimizers enabled")

	case tool == "asm" || strings.HasSuffix(tool, "a"): // assembler
		log.Printf("assembler output differs")

	case tool == "link" || strings.HasSuffix(tool, "l"): // linker
		log.Printf("linker output differs")
		extra = "-v=2"
	}

	cmdS := injectflags(cmd, []string{extra}, false)
	outfile, _ = cmpRun(true, cmdS)

	fmt.Fprintf(os.Stderr, "\n%s\n", compareLogs(outfile))
	os.Exit(2)
}

func injectflags(cmd []string, extra []string, addDashN bool) []string {
	x := []string{cmd[0]}
	if addDashN {
		x = append(x, "-N")
	}
	x = append(x, extra...)
	x = append(x, cmd[1:]...)
	return x
}

func cmpRun(keepLog bool, cmd []string) (outfile string, match bool) {
	cmdStash := make([]string, len(cmd))
	copy(cmdStash, cmd)
	cmdStash[0] = toolStash
	for i, arg := range cmdStash {
		if arg == "-o" {
			outfile = cmdStash[i+1]
			cmdStash[i+1] += ".stash"
			break
		}
		if strings.HasSuffix(arg, ".s") || strings.HasSuffix(arg, ".go") && '0' <= tool[0] && tool[0] <= '9' {
			outfile = filepath.Base(arg[:strings.LastIndex(arg, ".")] + "." + tool[:1])
			cmdStash = append([]string{cmdStash[0], "-o", outfile + ".stash"}, cmdStash[1:]...)
			break
		}
	}

	if outfile == "" {
		log.Fatalf("cannot determine output file for command: %s", strings.Join(cmd, " "))
	}

	if *norun {
		fmt.Printf("%s\n", strings.Join(cmd, " "))
		fmt.Printf("%s\n", strings.Join(cmdStash, " "))
		os.Exit(0)
	}

	out, err := runCmd(cmd, keepLog, outfile+".log")
	if err != nil {
		log.Printf("running: %s", strings.Join(cmd, " "))
		os.Stderr.Write(out)
		log.Fatal(err)
	}

	outStash, err := runCmd(cmdStash, keepLog, outfile+".stash.log")
	if err != nil {
		log.Printf("running: %s", strings.Join(cmdStash, " "))
		log.Printf("installed tool succeeded but stashed tool failed.\n")
		if len(out) > 0 {
			log.Printf("installed tool output:")
			os.Stderr.Write(out)
		}
		if len(outStash) > 0 {
			log.Printf("stashed tool output:")
			os.Stderr.Write(outStash)
		}
		log.Fatal(err)
	}

	return outfile, sameObject(outfile, outfile+".stash")
}

func sameObject(file1, file2 string) bool {
	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	b1 := bufio.NewReader(f1)
	b2 := bufio.NewReader(f2)

	// Go object files and archives contain lines of the form
	//	go object <goos> <goarch> <version>
	// By default, the version on development branches includes
	// the Git hash and time stamp for the most recent commit.
	// We allow the versions to differ.
	if !skipVersion(b1, b2, file1, file2) {
		return false
	}

	lastByte := byte(0)
	for {
		c1, err1 := b1.ReadByte()
		c2, err2 := b2.ReadByte()
		if err1 == io.EOF && err2 == io.EOF {
			return true
		}
		if err1 != nil {
			log.Fatalf("reading %s: %v", file1, err1)
		}
		if err2 != nil {
			log.Fatalf("reading %s: %v", file2, err1)
		}
		if c1 != c2 {
			return false
		}
		if lastByte == '`' && c1 == '\n' {
			if !skipVersion(b1, b2, file1, file2) {
				return false
			}
		}
		lastByte = c1
	}
}

func skipVersion(b1, b2 *bufio.Reader, file1, file2 string) bool {
	// Consume "go object " prefix, if there.
	prefix := "go object "
	for i := 0; i < len(prefix); i++ {
		c1, err1 := b1.ReadByte()
		c2, err2 := b2.ReadByte()
		if err1 == io.EOF && err2 == io.EOF {
			return true
		}
		if err1 != nil {
			log.Fatalf("reading %s: %v", file1, err1)
		}
		if err2 != nil {
			log.Fatalf("reading %s: %v", file2, err1)
		}
		if c1 != c2 {
			return false
		}
		if c1 != prefix[i] {
			return true // matching bytes, just not a version
		}
	}

	// Keep comparing until second space.
	// Must continue to match.
	// If we see a \n, it's not a version string after all.
	for numSpace := 0; numSpace < 2; {
		c1, err1 := b1.ReadByte()
		c2, err2 := b2.ReadByte()
		if err1 == io.EOF && err2 == io.EOF {
			return true
		}
		if err1 != nil {
			log.Fatalf("reading %s: %v", file1, err1)
		}
		if err2 != nil {
			log.Fatalf("reading %s: %v", file2, err1)
		}
		if c1 != c2 {
			return false
		}
		if c1 == '\n' {
			return true
		}
		if c1 == ' ' {
			numSpace++
		}
	}

	// Have now seen 'go object goos goarch ' in both files.
	// Now they're allowed to diverge, until the \n, which
	// must be present.
	for {
		c1, err1 := b1.ReadByte()
		if err1 == io.EOF {
			log.Fatalf("reading %s: unexpected EOF", file1)
		}
		if err1 != nil {
			log.Fatalf("reading %s: %v", file1, err1)
		}
		if c1 == '\n' {
			break
		}
	}
	for {
		c2, err2 := b2.ReadByte()
		if err2 == io.EOF {
			log.Fatalf("reading %s: unexpected EOF", file2)
		}
		if err2 != nil {
			log.Fatalf("reading %s: %v", file2, err2)
		}
		if c2 == '\n' {
			break
		}
	}

	// Consumed "matching" versions from both.
	return true
}

func runCmd(cmd []string, keepLog bool, logName string) (output []byte, err error) {
	if *verbose {
		log.Print(strings.Join(cmd, " "))
	}

	if *timing {
		t0 := time.Now()
		defer func() {
			log.Printf("%.3fs elapsed # %s\n", time.Since(t0).Seconds(), strings.Join(cmd, " "))
		}()
	}

	xcmd := exec.Command(cmd[0], cmd[1:]...)
	if !keepLog {
		return xcmd.CombinedOutput()
	}

	f, err := os.Create(logName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "GOOS=%s GOARCH=%s %s\n", os.Getenv("GOOS"), os.Getenv("GOARCH"), strings.Join(cmd, " "))
	xcmd.Stdout = f
	xcmd.Stderr = f
	defer f.Close()
	return nil, xcmd.Run()
}

func save() {
	if err := os.MkdirAll(stashDir, 0777); err != nil {
		log.Fatal(err)
	}

	toolDir := filepath.Join(goroot, fmt.Sprintf("pkg/tool/%s_%s", runtime.GOOS, runtime.GOARCH))
	files, err := ioutil.ReadDir(toolDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if shouldSave(file.Name()) && file.Mode().IsRegular() {
			cp(filepath.Join(toolDir, file.Name()), filepath.Join(stashDir, file.Name()))
		}
	}

	for _, name := range binTools {
		if !shouldSave(name) {
			continue
		}
		src := filepath.Join(binDir, name)
		if _, err := os.Stat(src); err == nil {
			cp(src, filepath.Join(stashDir, name))
		}
	}

	checkShouldSave()
}

func restore() {
	files, err := ioutil.ReadDir(stashDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if shouldSave(file.Name()) && file.Mode().IsRegular() {
			targ := toolDir
			if isBinTool(file.Name()) {
				targ = binDir
			}
			cp(filepath.Join(stashDir, file.Name()), filepath.Join(targ, file.Name()))
		}
	}

	checkShouldSave()
}

func shouldSave(name string) bool {
	if len(cmd) == 1 {
		return true
	}
	ok := false
	for i, arg := range cmd {
		if i > 0 && name == arg {
			ok = true
			cmd[i] = "DONE"
		}
	}
	return ok
}

func checkShouldSave() {
	var missing []string
	for _, arg := range cmd[1:] {
		if arg != "DONE" {
			missing = append(missing, arg)
		}
	}
	if len(missing) > 0 {
		log.Fatalf("%s did not find tools: %s", cmd[0], strings.Join(missing, " "))
	}
}

func cp(src, dst string) {
	if *verbose {
		fmt.Printf("cp %s %s\n", src, dst)
	}
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(dst, data, 0777); err != nil {
		log.Fatal(err)
	}
}