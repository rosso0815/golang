package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {

	fmt.Println("start ", os.Args[0])
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("i=", i, " args+", os.Args[i])
	}

	// For our example we'll exec `ls`. Go requires an
	// absolute path to the binary we want to execute, so
	// we'll use `exec.LookPath` to find it (probably
	// `/bin/ls`).
	//	binary, lookErr := exec.LookPath("ls")
	//	if lookErr != nil {
	//		panic(lookErr)
	//	}

	// `Exec` requires arguments in slice form (as
	// apposed to one big string). We'll give `ls` a few
	// common arguments. Note that the first argument should
	// be the program name.
	//	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	//	env := os.Environ()

	// Here's the actual `syscall.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.
	//	execErr := syscall.Exec(binary, args, env)
	//	if execErr != nil {
	//		fmt.Println("error")
	//		panic(execErr)
	//	}

	//_, err := os.Stat("tmp")
	//if err != nil {
	//	fmt.Println("error mkdir")
	//}

	//if os.IsNotExist(err) {
	//		return false, nil
	//	}

	//	fmt.Println("mkdir")
	//	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
	//		err = syscall.Mkdir("tmp", 0777)
	//		if err != nil {
	//			panic(err.Error)
	//		} else {
	//			fmt.Println("mkdir created passed")
	//		}
	//	}
	//	fmt.Println("mkdir passed")

	cmd := "/Users/pfistera/workspace/GO/bin/exec01"
	binary, lookErr := exec.LookPath(cmd)
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println(binary)
	os.Remove("/tmp/stdin")
	os.Remove("/tmp/stdout")
	os.Remove("/tmp/stderr")

	fstdin, err1 := os.Create("/tmp/stdin")
	fstdout, err2 := os.Create("/tmp/stdout")
	fstderr, err3 := os.Create("/tmp/stderr")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1, err2, err3)
		panic("WOW")
	}

	argv := []string{"hi"}
	procAttr := syscall.ProcAttr{
		Dir:   "/tmp",
		Files: []uintptr{fstdin.Fd(), fstdout.Fd(), fstderr.Fd()},
		Env:   []string{"VAR1=ABC123"},
		Sys: &syscall.SysProcAttr{
			Foreground: false,
		},
	}

	pid, err := syscall.ForkExec(binary, argv, &procAttr)
	fmt.Println("Spawned proc", pid, err)

	time.Sleep(time.Second * 100)
}
