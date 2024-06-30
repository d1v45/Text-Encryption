# Text Encryption

A Python script for encrypting and decrypting text messages using RSA encryption and Base64 encoding. The script generates public and private keys, encodes a message in Base64, encrypts it with RSA, and then decrypts it back to the original message.

## Features

- RSA key pair generation
- Base64 encoding and decoding
- RSA encryption and decryption
- Saving and loading keys and encrypted messages from files

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/d1v45/Text-Encryption.git
   ```

2. Navigate to the project directory:
   ```bash
   cd Text-Encryption
   ```

3. Run the script:
   ```bash
   python Text_Encryption.py
   ```

4. Follow the prompt to enter the message you want to encrypt.

## Example

```
ENTER THE MESSAGE: Hello, World!
Base64 Encoded message: SGVsbG8sIFdvcmxkIQ==
Encrypted text is: b'\x8b\xd4\xeb...'
Decrypted Base64 message: SGVsbG8sIFdvcmxkIQ==
Final plaintext message: Hello, World!
```

## File Paths

- Public key path: `./public.pem`
- Private key path: `./private.pem`
- Encrypted message path: `./enc_msg.txt`

## Requirements

- Python 3.x
- rsa module
- base64 module

Install the required modules using pip:
```bash
pip install rsa
```
