# One-Time Pad Encryption Tool

This tool provides a simple command-line interface for encrypting and decrypting messages using a one-time pad encryption scheme. It is written in Go and utilizes base64 encoding for safe text transmission.

## Features

- **Generate Random Key**: Creates a random key for encryption.
- **Base64 Encoding/Decoding**: Encodes and decodes strings to and from base64.
- **Encrypt/Decrypt Messages**: Encrypts messages with a one-time pad approach and decrypts them using the provided key.

## Installation

To install and run this tool, you need to have Go installed on your system. 

1. Clone this repository or download the source code.
2. Navigate to the directory containing the source code.
3. Compile the code using Go:

    ```bash
    go build -o executable .
    ```

4. Run the executable:

    ```bash
    ./executable
    ```

## Usage

The tool can be used with the following commands:

1. **Generate a Random Key**

    ```bash
    ./executable generateRandomKey [length]
    ```

    Replace `[length]` with the desired length of the key.

2. **Encode a String to Base64**

    ```bash
    ./executable EncodeTo64 [string]
    ```

    Replace `[string]` with the text you want to encode.

3. **Decode a Base64 String**

    ```bash
    ./executable DecodeFrom64 [base64String]
    ```

    Replace `[base64String]` with the base64 encoded text.

4. **Encrypt a Message**

    ```bash
    ./executable encrypt [message]
    ```

    Replace `[message]` with the text you want to encrypt. The output will be the encrypted message and the key used for encryption.

5. **Decrypt a Message**

    ```bash
    ./executable decrypt [encryptedMessage] [key]
    ```

    Replace `[encryptedMessage]` with your encrypted text and `[key]` with the key received during encryption.

## Note

This encryption tool is designed for educational purposes and should not be used for securing sensitive data.
