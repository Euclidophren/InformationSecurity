def fillRotor(rotor, file):
	with open(file) as f:
		lines = f.readlines()
		for i in range(len(lines)):
			line = lines[i].rstrip('\n')
			(key, val) = line.split(' ')
			rotor[key] = val[0]
		f.close()

def fillRotors():
	rotors = []
	for i in range(4):
		rotor = {}
		rot = "rot" + str(i)
		fillRotor(rotor, rot)
		rotors.append(rotor)
	return rotors

def fillReflector():
	reflector = {}
	fillRotor(reflector, "reflector")
	return reflector

