import sys
from random import randrange
import struct
from math import sqrt

def check_prime(number):
    for i in range(1, int(sqrt(number))):
        if(number % i == 0):
            return False
    return True

def xgcd(a, b):
    x0, x1, y0, y1 = 0, 1, 1, 0
    while a != 0:
        q, b, a = b // a, a, b % a
        y0, y1 = y1, y0 - q * y1
        x0, x1 = x1, x0 - q * x1
    return b, x0, y0

def mulinv(a, b):
    g, x, _ = xgcd(a, b)
    if g == 1:
        return x % b

def calculateModulusAndPhi(p, q):
    n = p * q
    phi = (p - 1) * (q - 1)
    return n, phi

def calculateEAndD(phi):
    while True:
        e = randrange(2, phi)
        modulus, x, _ = xgcd(e, phi)
        if(modulus == 1):
            d = mulinv(e, phi) 
            return e, d

def encrypt(filename_read, filename_write, e, n):
       with open(filename_read, "rb") as fr, open(filename_write, "w") as fw:
               data = fr.read()
               for item in data:
                   new_item = pow(item, e, n)
                   fw.write(str(new_item) + "\n")

def decrypt (filename_read, filename_write, d, n):   
    with open(filename_read, "r") as fr, open(filename_write, "wb") as fw:

        line = fr.readline()
        while line:
            num = int(line)
            byte = pow(num, d, n)
            fw.write(struct.pack('B', byte)) 
            line = fr.readline()

if __name__ == '__main__':
        filename = sys.argv[1]

        if len(sys.argv) == 2:
            p = 199
            q = 179
        else:
            p = int(sys.argv[2])
            q = int(sys.argv[3])
        with open(filename, 'rb') as file1:
            data = file1.read()
            n, phi = calculateModulusAndPhi(p, q)
            e, d = calculateEAndD(phi)
            print("\tP:", p, "\n\tQ:", q, "\n\tE:", e, "\n\tN:", n, "\n\tD:", d)
            print("Encrypting...")
            encrypt(filename, filename.split(".")[0] + ".encoded", e, n)
            print("Decrypting...")
            decrypt(filename.split('.')[0] + ".encoded", filename.split('.')[0] + ".decoded", d, n)
