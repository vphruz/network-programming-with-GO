package devcon

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type sshClient struct {
	target string
	cfg    *ssh.ClientConfig
}

func NewClient(target string) *sshClient {
	return &sshClient{
		target: target,
		cfg: &ssh.ClientConfig{
			User: os.Getenv("SSH_USER"),
			Auth: []ssh.AuthMethod{
				ssh.Password(os.Getenv("SSH_PASSWORD")),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Second * 5,
		},
	}
}

func (c *sshClient) Run(cmd string) (string, error) {
	// had to change from the default port of 22 to 4444 because of my pc.
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:4444", c.target), c.cfg)
	if err != nil {
		return "", errors.Wrap(err, "dial failed")
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
