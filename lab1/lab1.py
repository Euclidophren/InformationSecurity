from hashlib import sha256
import subprocess
import hmac
import os

appKey = "some-secret-key"
key = "654e88c6c6909fe89eb8383c8f5e2161b092bc6b194d96a4d65468c68282ca55"

def generateKey(appId):
    err = False
    try:
        id = subprocess.check_output("dmidecode -s system-uuid", shell=True).decode().split(":")[0]
        mac = hmac.new(bytearray(id, encoding='utf-8'), digestmod=sha256)
        mac.update(bytearray(appId,encoding='utf-8'))
        return mac.hexdigest(), err
    except subprocess.SubprocessError:
        err = True
        return None, err

def main():
    generatedKey, err = generateKey(appKey)
    if(err):
        print("Error: Subprocess has non-zero return")
    else:
        if(key == generatedKey):
            print("Key is valid")
        else:
            print("Key is invalid")

if __name__ == "__main__":
    main()



