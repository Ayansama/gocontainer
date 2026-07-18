package namespace

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run() error {
	fmt.Println("[run]forking into new namespace")

	cmd := exec.Command("/proc/self/exe", "child")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
	}

	return cmd.Run()
}

func Child() error {
	fmt.Println("[child] inside isolated namspace")

	if err := syscall.Sethostname([]byte("gocontainer")); err != nil {
		return fmt.Errorf("sethostname: %w", err)
	}

	flags := uintptr(syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV)
	if err := syscall.Mount("proc", "/proc", "proc", flags, ""); err != nil {
		return fmt.Errorf("mount proc: %w", err)
	}
	fmt.Println("[child] /proc mounted")

	env := []string{
		"TERM=xterm-256color",
		"HOME=/root",
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"PS1=[gocontainer]# ",
	}

	return syscall.Exec("/bin/bash", []string{"/bin/bash"}, env)
}
