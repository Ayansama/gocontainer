package pty

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/creack/pty"
)

type Session struct {
	Master *os.File
	cmd    *exec.Cmd
}

func Start() (*Session, error) {
	cmd := exec.Command("/proc/self/exe", "child")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
		Setsid: true,
	}

	master, err := pty.Start(cmd)
	if err != nil {
		return nil, fmt.Errorf("pty.Start: %w", err)
	}

	fmt.Println("[pty] shell started behind PTY master")
	return &Session{Master: master, cmd: cmd}, nil
}

func (s *Session) Resize(rows, cols uint16) error {
	return pty.Setsize(s.Master, &pty.Winsize{Rows: rows, Cols: cols})
}

func (s *Session) Wait() error {
	defer s.Master.Close()
	return s.cmd.Wait()
}
