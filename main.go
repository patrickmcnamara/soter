package main

import (
	"fmt"
	"os"
	"time"

	"github.com/patrickmcnamara/soter/crypto"
	"github.com/patrickmcnamara/soter/encoding"
	"github.com/patrickmcnamara/soter/files"
	"github.com/patrickmcnamara/soter/flags"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:        "Soter",
		Usage:       "Sign and verify files",
		Version:     "0.1 (alpha)",
		Description: "Sign and verify files using public key cryptography.",
		Authors:     []*cli.Author{{Name: "Patrick McNamara", Email: "hello@patrickmcnamara.xyz"}},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "generate-keypair",
			Aliases: []string{"gk"},
			Usage:   "Generate a keypair",

			Flags: []cli.Flag{
				flags.KeyfileFlag,
			},

			Action: func(c *cli.Context) error {
				keyfile := c.String("keyfile")

				fmt.Print("Soter is generating a new keypair...\n\n")
				time.Sleep(time.Second)
				publicKey, privateKey := crypto.GenerateKeypair()
				publicKeyEncoded := encoding.PublicKeyToEncoded(publicKey)
				files.WriteKeyfile(keyfile, privateKey)
				fmt.Printf("Public key is \"%s\".\n", publicKeyEncoded)
				fmt.Printf("Private key is in keyfile \"%s\".\n", keyfile)
				return nil
			},
		},
		{
			Name:    "backup-keyfile",
			Aliases: []string{"bk"},
			Usage:   "Backup a keyfile",

			Flags: []cli.Flag{
				flags.BackupKeyfileFlag,
				flags.KeyfileFlag,
			},

			Action: func(c *cli.Context) error {
				backupKeyfile := c.String("file")
				keyfile := c.String("keyfile")

				fmt.Print("Soter is backupping a keyfile...\n\n")
				time.Sleep(time.Second)
				files.BackupKeyfile(backupKeyfile, keyfile)
				fmt.Printf("Backupped keyfile \"%s\" to keyfile \"%s\".\n", keyfile, backupKeyfile)
				return nil
			},
		},
		{
			Name:    "restore-keyfile",
			Aliases: []string{"rk"},
			Usage:   "Restore a keyfile",

			Flags: []cli.Flag{
				flags.BackupKeyfileFlag,
				flags.KeyfileFlag,
			},

			Action: func(c *cli.Context) error {
				backupKeyfile := c.String("file")
				keyfile := c.String("keyfile")

				fmt.Print("Soter is restoring a keyfile...\n\n")
				time.Sleep(time.Second)
				files.RestoreKeyfile(backupKeyfile, keyfile)
				fmt.Printf("Restored keyfile \"%s\" to keyfile \"%s\".\n", backupKeyfile, keyfile)
				return nil
			},
		},
		{
			Name:    "change-keyfile-password",
			Aliases: []string{"ckp"},
			Usage:   "Change password to keyfile",

			Flags: []cli.Flag{
				flags.KeyfileFlag,
			},

			Action: func(c *cli.Context) error {
				keyfile := c.String("keyfile")

				fmt.Print("Soter is changing password to a keyfile...\n\n")
				time.Sleep(time.Second)
				files.ChangeKeyfilePassword(keyfile)
				fmt.Printf("Changed password for keyfile \"%s\".\n", keyfile)
				return nil
			},
		},
		{
			Name:    "sign",
			Aliases: []string{"s"},
			Usage:   "Sign a file using a keyfile",

			Flags: []cli.Flag{
				flags.SignFileFlag,
				flags.KeyfileFlag,
			},

			Action: func(c *cli.Context) error {
				keyfile := c.String("keyfile")
				file := c.String("file")

				fmt.Print("Soter is signing a file with a private key...\n\n")
				time.Sleep(time.Second)
				privateKey := files.ReadKeyfile(keyfile)
				files.SignFile(file, privateKey)
				fmt.Printf("Signed file \"%s\" with private key in keyfile \"%s\".\n", file, keyfile)
				return nil
			},
		},
		{
			Name:    "verify",
			Aliases: []string{"v"},
			Usage:   "Verify a file with a public key",

			Flags: []cli.Flag{
				flags.VerifyFileFlag,
				flags.PublicKeyFlag,
			},

			Action: func(c *cli.Context) error {
				file := c.String("file")
				publicKeyEncoded := c.String("public-key")

				fmt.Print("Soter is verifying a file with a public key...\n\n")
				time.Sleep(time.Second)
				publicKey := encoding.EncodedToPublicKey(publicKeyEncoded)
				verification := files.VerifyFile(file, publicKey)
				fmt.Printf("Verified file \"%s\" with public key \"%s\" as %t.\n", file, publicKeyEncoded, verification)
				return nil
			},
		},
		{
			Name:    "print-public-key",
			Aliases: []string{"ppk"},
			Usage:   "Print a public key from a keyfile",

			Flags: []cli.Flag{
				flags.KeyfileFlag,
			},

			Action: func(c *cli.Context) error {
				keyfile := c.String("keyfile")

				fmt.Print("Soter is printing a public key from a keyfile...\n\n")
				time.Sleep(time.Second)
				privateKey := files.ReadKeyfile(keyfile)
				publicKey := crypto.GetPublicKey(privateKey)
				publicKeyEncoded := encoding.PublicKeyToEncoded(publicKey)
				fmt.Printf("Public key \"%s\" from keyfile \"%s\".\n", publicKeyEncoded, keyfile)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
