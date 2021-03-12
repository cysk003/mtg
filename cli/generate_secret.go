package cli

import (
	"fmt"

	"github.com/9seconds/mtg/v2/mtglib"
)

type GenerateSecret struct {
	base

	HostName string `arg optional help:"Hostname to use for domain fronting. Default is '${domain_front}'." name:"hostname" default:"${domain_front}"` // nolint: lll, govet
	Hex      bool   `help:"Print secret in hex encoding."`
}

func (c *GenerateSecret) Run(cli *CLI, _ string) error {
	secret := mtglib.GenerateSecret(cli.GenerateSecret.HostName)

	if cli.GenerateSecret.Hex {
		fmt.Println(secret.Hex()) // nolint: forbidigo
	} else {
		fmt.Println(secret.Base64()) // nolint: forbidigo
	}

	return nil
}