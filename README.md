# Soter

Sign and verify files using public key cryptography.

More specifically, it uses Ed25519 to sign and verify files. It uses ChaCha20 with Argon2id to encrypt keyfiles.

```
USAGE:
   soter [global options] command [command options] [arguments...]
VERSION:
   0.1 (alpha)
COMMANDS:
     generate-keypair, gk   Generate a keypair
     backup-keypair, bk     Backup a keypair
     restore-keypair, rk    Restore a keypair
     sign, s                Sign a file
     verify, v              Verify a file
     print-public-key, ppk  Print a public key
     help, h                Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --help, -h               show help (default: false)
   --init-completion value  generate completion code. Value must be 'bash' or 'zsh'
   --version, -v            print the version (default: false)
```
