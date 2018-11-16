// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"hidevops.io/hiboot/pkg/app"
	"hidevops.io/hiboot/pkg/app/cli"
	"hidevops.io/hiboot/pkg/utils/crypto/rsa"
)

// define the command
type cryptoCommand struct {
	// embedding cli.BaseCommand in each command
	cli.RootCommand
	Source  string
	Encrypt bool
	Decrypt bool
	Key     string
}

func init() {
	app.Register("cli.rootCommand", newCryptoCommand)
}

// newCryptoCommand crypto command constructor
func newCryptoCommand() *cryptoCommand {
	c := new(cryptoCommand)
	c.Use = "crypto"
	c.Short = "crypto command"
	c.Long = `crypto is a command line tool that encrypt / decrypt text, please run crypto command to encrypt/decrypt 
	`
	c.Example = `
crypto rsa -h
crypto rsa -e -s "text to encrypt"
crypto rsa -d -s "text to decrypt"
`
	p := c.PersistentFlags()
	p.StringVarP(&c.Source, "source", "s", "", "run with option --source=source text to encrypt or encrypt")
	p.StringVarP(&c.Key, "key", "k", "", "run with option --key or -k for rsa key")
	p.BoolVarP(&c.Encrypt, "encrypt", "e", false, "run with option --encrypt or -e for text encryption")
	p.BoolVarP(&c.Decrypt, "decrypt", "d", false, "run with option --decrypt or -d for text encryption")
	return c
}

// Run OnRsa for crypto command rsa
func (c *cryptoCommand) OnRsa(args []string) bool {
	if c.Decrypt {
		res, err := rsa.DecryptBase64([]byte(c.Source), []byte(c.Key))
		if err == nil {
			fmt.Println(string(res))
		}
	} else {
		res, err := rsa.EncryptBase64([]byte(c.Source), []byte(c.Key))
		if err == nil {
			fmt.Println(string(res))
		}
	}
	return true
}
