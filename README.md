🔐 AES-GCM Encryption Sample in Python

This code demonstrates how to generate an "authorization" field for an API request using AES-GCM encryption in Python.

Features:

AES-GCM encryption with a 256-bit key
12-byte random nonce (IV)
Hex-encoded output for easy use in JSON

Dependencies:

cryptography

Install:

pip install cryptography
Example Use Case:

For an API that expects a field like:

{
  "authorization": "ciphertext-<nonce>-<tag>"
}
This script generates that exact structure.
