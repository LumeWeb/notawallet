# "Definitely Not a Wallet" Key Generator

## Disclaimer

This project is a work of satire and political protest. It is intended to demonstrate the absurdity of attempting to regulate cryptographic algorithms and mathematical operations. The author is exercising their First Amendment rights to free speech and protest against governmental overreach in the realm of cryptography and privacy.

**Legal Notice**: This project is created and distributed for educational and protest purposes only. It is not intended for any illegal activities. The creator(s) of this project are not responsible for any misuse of this software. Use at your own risk.

## About the Project

The "Definitely Not a Wallet" Key Generator is a satirical Go program that benchmarks various cryptographic key generation processes. It humorously frames each key generation as a "violation" to highlight the impracticality and potential consequences of over-regulating cryptographic technologies.

This project is a direct response to the IRS's attempt to regulate non-custodial (also known as "self-hosted") wallets as broker/dealers. This proposed regulation represents a significant overreach and misunderstanding of the nature of cryptographic technology.

## What's Happening: IRS Regulation Attempt (Explained for Non-Techies)

Imagine if the government tried to regulate calculators as if they were banks, just because calculators can be used to count money. That's essentially what's happening here, but with cryptography.

- **What's a non-custodial wallet?** It's like a digital piggy bank that only you have the key to. No bank or company controls it - just you.
- **What's the IRS trying to do?** They want to treat the software that creates these digital piggy banks as if it were a financial institution, like a bank or a stockbroker.
- **Why is this absurd?** Because this software is just math and code. It doesn't hold or transfer money itself. It's a tool, like a calculator or a lockbox.
- **What's the problem?** This regulation could make it very difficult or even illegal to create or use these digital piggy banks, seriously impacting privacy and innovation in the digital world.

This project demonstrates how common cryptographic operations - the same kind used in everyday internet security - could be caught up in this overly broad regulation.

## Features

- Generates and benchmarks:
    - ECDSA Private Keys
    - BIP39 Seeds (crypto wallets)
    - GPG Keys (Go implementation)
    - SSH Keys
    - SSL Certificates
- Calculates hypothetical "processing time" for bureaucratic oversight of key generation

## Installation and Usage

You can install the "Definitely Not a Wallet" Key Generator using Go's package management system. Here are the steps:

1. Ensure you have Go installed on your system (version 1.16 or later is recommended).

2. Open a terminal and run the following command to install the package:

   ```
   go install go.lumeweb.com/notawallet@latest
   ```

3. Once installed, you can run the program by simply typing:

   ```
   notawallet
   ```

   This will execute the key generation benchmarks and display the results.

### Building from Source

If you prefer to build from source or want to contribute to the project:

1. Clone the repository:
   ```
   git clone https://github.com/LumeWeb/notawallet.git
   ```

2. Navigate to the project directory:
   ```
   cd notawallet
   ```

3. Build the project:
   ```
   go build
   ```

4. Run the executable:
   ```
   ./notawallet
   ```

## The Absurdity of Regulating Math

Let me put this in a way that non-devs can understand.

In the 1990s, there was a battle with PGP that informally ended in a stalemate with "code is speech". This is about cryptography (encryption, not blockchain). Then we got SSH, SSL, and crypto wallets.

- PGP: Used for email encryption and a non-crypto form of identity. Has gained popularity as "verified identity" in signing git commits as well. Uses RSA in most cases.
- SSH: Used mostly for Linux servers for remote control, most often RSA or ed25519 keys.
- SSL: Used for banking, e-commerce, etc. Certs started with a private key that, guess what... uses RSA most often.
- Crypto wallets: Uses various cryptography algorithms, but often is deterministic based on a BIP39 seed.

All of these use what's called an RNG, or random number generator. Most simply, it will have a function spit out about 32 bytes of data, and then can use that to generate a public key. In crypto (blockchain), you then hash that for the wallet address. By hashing, I mean the same as password hashes, the ones that website MySQL databases store for your web logins.

So imagine trying to regulate software that does this. Imagine trying to regulate `ssh-keygen` or `GnuPG` because the same private key could technically be used as a wallet private key! That is fucking crazy.

You're trying to regulate math. I repeat, you're TRYING TO REGULATE MATH.

This is antithetical to the web.

## Contributing

This project is meant to spark discussion and raise awareness. Feel free to fork, modify, and share as you see fit, keeping in mind the satirical and educational nature of the project. Pull requests are welcome!

## License

This project is released under the MIT License. See the LICENSE file for details.

Remember: The Empire's regulations are definitely enforceable and not at all absurd. 😉