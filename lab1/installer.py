from hashlib import sha256
import subprocess
import hmac
import os

appKey = "some-secret-key"

def protect(appId, id):
    mac = hmac.new(bytearray(id, encoding='utf-8'), digestmod=sha256)
    mac.update(bytearray(appId, encoding='utf-8'))
    return mac.hexdigest()

def main():
    try:
        machineId = subprocess.check_output("dmidecode -s system-uuid", shell=True).decode().split(":")[0]
        toSave = ""
        machineFile = open("lab1.py", "r")

        for line in machineFile:
            if(line.startswith("key")):
                toSave += "key = \"" + str(protect(appKey, machineId)) + "\"\n"
            else:
                toSave += line
        machineFile.close()
        machineFile = open("lab1.py", "w")
        machineFile.write(toSave)
        machineFile.close()
    except subprocess.SubprocessError:
        print("Process output is non-zero")

if __name__ == "__main__":
    main()