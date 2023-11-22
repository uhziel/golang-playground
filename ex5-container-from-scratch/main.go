package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func run(args []string) {
	log.Println("running", args, "as", os.Getpid())
	newArgs := []string{}
	newArgs = append(newArgs, "child")
	newArgs = append(newArgs, args...)
	cmd := exec.Command("/proc/self/exe", newArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		log.Fatalf("run: cmd.Run(): %v", err)
	}
}

func child(name string, args []string) {
	log.Println("running", name, args, "as", os.Getpid())

	syscall.Sethostname([]byte("zlcontainer"))
	syscall.Chroot("./bundle-alpine/rootfs")
	syscall.Chdir("/")
	syscall.Mount("proc", "/proc", "proc", 0, "")

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("child: cmd.Run(): %v", err)
	}
	syscall.Unmount("/proc", 0)
}

func main() {
	switch os.Args[1] {
	case "run":
		run(os.Args[2:])
	case "child":
		child(os.Args[2], os.Args[3:])
	default:
		log.Println("invalid.")
	}
}
