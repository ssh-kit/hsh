package exec

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/crypto/ssh"
)

func OpenSSHCmd(urlStr string) (*SSHCmd, error) {
	c := &SSHCmd{}
	return c, c.Open(urlStr)
}

type SSHCmd struct {
	*ssh.Client
}

func (c *SSHCmd) Open(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("parse url: %v", err)
	}
	if u.Port() == "" {
		u.Host += ":22"
	}

	cfg := &ssh.ClientConfig{
		User:            u.User.Username(),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if passwd, ok := u.User.Password(); ok {
		cfg.Auth = []ssh.AuthMethod{
			ssh.Password(passwd),
		}
	}
	client, err := ssh.Dial("tcp", u.Host, cfg)
	if err != nil {
		return fmt.Errorf("dial ssh: %v", err)
	}
	c.Client = client

	return nil
}

func (c *SSHCmd) Run(ctx context.Context, cmd *Cmd) (*Result, error) {
	sess, err := c.Client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("new session: %v", err)
	}
	defer sess.Close()

	var stdout, stderr bytes.Buffer
	res := new(Result)
	cc := []string{
		cmd.Path,
	}
	if len(cmd.Args) > 0 {
		cc = append(cc, cmd.Args...)
	}
	sess.Stdout = &stdout
	sess.Stderr = &stderr
	if err := sess.Run(strings.Join(cc, " ")); err != nil {
		switch err := err.(type) {
		case *ssh.ExitError:
			res.Code = err.ExitStatus()
		default:
			res.Code = CodeIOError
		}
		if res.Code < 0 || res.Code > 255 {
			res.Code = CodeOutRange
		}
	}
	res.Stdout = stdout.String()
	res.Stderr = stderr.String()

	return res, nil
}

func (c *SSHCmd) Close() error {
	if c.Client != nil {
		return c.Client.Close()
	}
	return nil
}
