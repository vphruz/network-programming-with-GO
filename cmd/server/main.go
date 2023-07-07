package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(handleCommands)
	log.Println("firing up the server...")
	log.Fatal(ssh.ListenAndServe("127.0.0.1:4444", nil,
		// had to change from the default port of 22 to 4444 because of my pc.
		ssh.HostKeyFile(filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")),
		ssh.PasswordAuth(
			ssh.PasswordHandler(func(ctx ssh.Context, password string) bool {
				return password == os.Getenv("SSH_PASSWORD")
			}),
		),
	))
}

func handleCommands(s ssh.Session) {
	switch s.RawCommand() {
	case "show version":
		ver := `Hostname:lab_srx
		Model: srx210he
		JUNOS Software Release[12.1X44-045.2]`

		io.WriteString(s, ver)
	}
}
