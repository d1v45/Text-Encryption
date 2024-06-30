######################################################################################################
# Title: TEXT ENCRYPTION                                                                             #
# Author: DIVAS A S                                                                                  #
# Github : https://github.com/d1v45                                                                  #
######################################################################################################

import rsa
import base64
import os

def banner():
    print('''
\t"    █▀▀█    ┏━━┓░░░░░░░░░┏┓┏┓░┏┳┓      ░░
\t"    █░      ┃━━┫┏━┓┏┳┓┏┳┓┣┫┃┗┓┃┃┃    █▀▀█
\t"    █▄▄█    ┣━━┃┃━┫┃┃┃┃┏┛┃┃┃┏┫┣┓┃      ░█
\t"    ░░      ┗━━┛┗━┛┗━┛┗┛░┗┛┗━┛┗━┛    █▄▄█
''')
    

if __name__ == '__main__':
    banner()
    example_msg = input("ENTER THE MESSAGE : ")
    public_key, private_key = rsa.newkeys(2048)

    current_directory = os.getcwd()
    public_key_path = os.path.join(current_directory, "public.pem")
    private_key_path = os.path.join(current_directory, "private.pem")
    encrypted_message_path = os.path.join(current_directory, "enc_msg.txt")

    with open(public_key_path, "wb") as f:
        f.write(public_key.save_pkcs1("PEM"))

    with open(private_key_path, "wb") as f:
        f.write(private_key.save_pkcs1("PEM"))

    with open(public_key_path, "rb") as f:
        public_key = rsa.PublicKey.load_pkcs1(f.read())

    with open(private_key_path, "rb") as f:
        private_key = rsa.PrivateKey.load_pkcs1(f.read())

    def encode_base64(input_string: str) -> str:
        bytes_input = input_string.encode('utf-8')
        base64_bytes = base64.b64encode(bytes_input)
        return base64_bytes.decode('utf-8')

    def decode_base64(base64_string: str) -> str:
        base64_bytes = base64_string.encode('utf-8')
        decoded_bytes = base64.b64decode(base64_bytes)
        return decoded_bytes.decode('utf-8')

    base64_encoded_message = encode_base64(example_msg)
    print("Base64 Encoded message:", base64_encoded_message)

    encrypted_message = rsa.encrypt(base64_encoded_message.encode(), public_key)
    print("Encrypted text is:", encrypted_message)

    with open(encrypted_message_path, "wb") as f:
        f.write(encrypted_message)

    with open(encrypted_message_path, "rb") as f:
        encrypted_message = f.read()

    decrypted_base64_message = rsa.decrypt(encrypted_message, private_key)
    print("Decrypted Base64 message:", decrypted_base64_message.decode())

    final_message = decode_base64(decrypted_base64_message.decode())
    print("Final plaintext message:", final_message)
