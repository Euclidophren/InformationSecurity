from Crypto.PublicKey import RSA
from Crypto.Hash import SHA256
from Crypto.Signature import pkcs1_15
import os

def readFromFile(filename):
    fileToRead = open(filename, "rb")
    data = fileToRead.read()
    fileToRead.close()
    return data

def writeToFile(data, filename):
    fileToWrite = open(filename, "wb")
    fileToWrite.write(data)
    fileToWrite.close()

def main():
    source = input("Enter file name: ")
    data = readFromFile(source)

    privateKey = RSA.generate(2048)
    publicKey = privateKey.publickey()

    hashed = SHA256.new(data)
    signature = pkcs1_15.new(privateKey).sign(hashed)

    writeToFile(signature, source + ".pem")

    signature = readFromFile(source + ".pem")
    data = readFromFile(source)

    hashed = SHA256.new(data)
    try:
        pkcs1_15.new(publicKey).verify(hashed, signature)
        print("VALID signature")
    except (ValueError, TypeError):
        print("INVALID signature")

if __name__ == "__main__":
    main()