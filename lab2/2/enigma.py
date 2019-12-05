import rotors as rt
from collections import OrderedDict
from itertools import islice, cycle
import base64

BASE = 26
rotors = rt.fillRotors()
reflector = rt.fillReflector()
count = 0

def rotateRotor(rotor, shift):
	shift %= len(rotor)
	return OrderedDict(
		(k, v)
		for k, v in zip(rotor.keys(), islice(cycle(rotor.values()), shift, None))
	)

def applyRotor(rotor, symbol):
	return rotor[symbol]

def reverseApplyRotor(rotor, symbol):
	return getKey(rotor, symbol)

def getKey(dictionary, value):
	items = dictionary.items()
	for item in items:
		if item[1] == value:
			key = item[0]
	return  key

def goThroughRotors(rotors, inputStr):
	res = ""
	global count
	global BASE
	for sym in inputStr:
		for i in range(len(rotors)):
			sym = applyRotor(rotors[i], sym)
		sym = applyRotor(reflector, sym)
		for i in range(len(rotors)):
			sym = reverseApplyRotor(rotors[len(rotors) - i - 1], sym)
		res += sym		
		for i in range(len(rotors)):
			if count >= BASE*(i + 1) - BASE and count < BASE*(i + 1):
				rotors[i] = rotateRotor(rotors[i], 1)
				break
		count += 1
	return res

ch = 1
while ch != "0":
	ch = input("0 - stop program\n1 - encode\n2 - decode\n    Input: ")
	if ch == "1":
		filename = input("	File name: ")
		with open(filename, 'rb') as file:			
			output = input("	Output file name: ")
			data = file.read()
			encodedBytes = base64.b64encode(data)
			inputString = str(encodedBytes, "utf-8")
			res = goThroughRotors(rotors, inputString)
			rotors = rt.fillRotors()
			count = 0
			file.close()			
			outputString = res
			with open(output, 'w') as file:
				file.write(outputString)
			file.close()
	elif ch == "2":
		filename = input("	File name: ")
		with open(filename) as file:			
			output = input("	Output file name: ")
			data = file.read()
			inputString = data.rstrip('\n')
			res = goThroughRotors(rotors, inputString)
			rotors = rt.fillRotors()
			count = 0
			file.close()
			decodedBytes = base64.b64decode(res)	
			with open(output, 'wb') as file:
				file.write(decodedBytes)
			file.close()


