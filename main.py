from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend
import os

hex_key = "9865a8d227283b9243e52d0f119838e3c40f268f6b0c36a7c494b3e0afc9b65e"
key = bytes.fromhex(hex_key)

transactionId = b"the-trx-id" #transaction id

nonce = os.urandom(12)
encryptor = Cipher(
    algorithms.AES(key),
    modes.GCM(nonce),
    backend=default_backend()
).encryptor()

ciphertext = encryptor.update(transactionId) + encryptor.finalize()
tag = encryptor.tag

print("Ciphertext (hex):", ciphertext.hex())
print("Nonce (hex):", nonce.hex())
print("Tag (hex):", tag.hex())
