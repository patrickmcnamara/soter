# Soter

Sign and verify files using public key cryptography.

More specifically, it uses Ed25519 to sign and verify files. It uses ChaCha20 with Argon2id to encrypt keyfiles.

## Usage

Use `soter generate-keypair` or `soter gk` to generate a keypair and store it in a keyfile. This will also print out your public key. Distribute it far and wide, if you want. Your private key will be encrypted in your keyfile.

Now that you have a keypair, you can sign files using `soter sign --file FILENAME` or `soter s -f FILENAME`. This will create a new signature file which is the filename prepended with ".s.soter".

If you want to verify a file, you need three things: the file, the signature and the public key of the signer. Use `soter verify --file FILENAME --public-key PUBLICKEY` or `soter v -f FILENAME -pk PUBLICKEY`. This will print out whether or not it was successfully verified.

You are able to customise things a bit more than this. Be sure to use the help flags to delve deeper into each command.
