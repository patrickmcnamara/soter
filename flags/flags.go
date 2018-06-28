package flags

import (
	"os/user"

	"gopkg.in/urfave/cli.v2"
)

var currentUser, _ = user.Current()
var homeDirectory = currentUser.HomeDir

var keyfile string
var file string
var publicKeyEncoded string

var KeyfileFlag = &cli.StringFlag{
	Name:        "keyfile",
	Aliases:     []string{"k"},
	Usage:       "use `file` as keyfile",
	Value:       homeDirectory + "/.k.soter",
	Destination: &keyfile,
}

var BackupKeyfileFlag = &cli.StringFlag{
	Name:        "file",
	Aliases:     []string{"f"},
	Usage:       "use `file` as backup keyfile",
	Value:       homeDirectory + "/.k.soter.BAK",
	Destination: &file,
}

var SignFileFlag = &cli.StringFlag{
	Name:        "file",
	Aliases:     []string{"f"},
	Usage:       "use `file` for signing",
	Destination: &file,
}

var VerifyFileFlag = &cli.StringFlag{
	Name:        "file",
	Aliases:     []string{"f"},
	Usage:       "use `file` for verification",
	Destination: &file,
}

var PublicKeyFlag = &cli.StringFlag{
	Name:        "public-key",
	Aliases:     []string{"pk"},
	Usage:       "use `string` as public key",
	Destination: &publicKeyEncoded,
}
