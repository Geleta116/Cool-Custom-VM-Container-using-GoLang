package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
	"golang.org/x/sys/unix"
)

const (
    CLONE_NEWUTS  = 0x04000000
    CLONE_NEWPID  = 0x20000000
    CLONE_NEWNS   = 0x00020000
)

func main() {
    switch os.Args[1] {
    case "run":
        parent()
    case "child":
        child()
    default:
        panic("not working")}}
    
func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...))
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: unix.CLONE_NEWUTS | unix.CLONE_NEWPID | unix.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}}
		

func child() {
    must(unix.Mount("rootfs", "rootfs", "", unix.MS_BIND, ""))
    must(os.MkdirAll("rootfs/oldrootfs", 0700))
    must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
    must(os.Chdir("/"))

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}