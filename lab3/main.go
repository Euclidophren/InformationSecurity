package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

var ivSize = 16

var sBox = []byte{
	0x63, 0x7C, 0x77, 0x7B, 0xF2, 0x6B, 0x6F, 0xC5, 0x30, 0x01, 0x67, 0x2B, 0xFE, 0xD7, 0xAB, 0x76,
	0xCA, 0x82, 0xC9, 0x7D, 0xFA, 0x59, 0x47, 0xF0, 0xAD, 0xD4, 0xA2, 0xAF, 0x9C, 0xA4, 0x72, 0xC0,
	0xB7, 0xFD, 0x93, 0x26, 0x36, 0x3F, 0xF7, 0xCC, 0x34, 0xA5, 0xE5, 0xF1, 0x71, 0xD8, 0x31, 0x15,
	0x04, 0xC7, 0x23, 0xC3, 0x18, 0x96, 0x05, 0x9A, 0x07, 0x12, 0x80, 0xE2, 0xEB, 0x27, 0xB2, 0x75,
	0x09, 0x83, 0x2C, 0x1A, 0x1B, 0x6E, 0x5A, 0xA0, 0x52, 0x3B, 0xD6, 0xB3, 0x29, 0xE3, 0x2F, 0x84,
	0x53, 0xD1, 0x00, 0xED, 0x20, 0xFC, 0xB1, 0x5B, 0x6A, 0xCB, 0xBE, 0x39, 0x4A, 0x4C, 0x58, 0xCF,
	0xD0, 0xEF, 0xAA, 0xFB, 0x43, 0x4D, 0x33, 0x85, 0x45, 0xF9, 0x02, 0x7F, 0x50, 0x3C, 0x9F, 0xA8,
	0x51, 0xA3, 0x40, 0x8F, 0x92, 0x9D, 0x38, 0xF5, 0xBC, 0xB6, 0xDA, 0x21, 0x10, 0xFF, 0xF3, 0xD2,
	0xCD, 0x0C, 0x13, 0xEC, 0x5F, 0x97, 0x44, 0x17, 0xC4, 0xA7, 0x7E, 0x3D, 0x64, 0x5D, 0x19, 0x73,
	0x60, 0x81, 0x4F, 0xDC, 0x22, 0x2A, 0x90, 0x88, 0x46, 0xEE, 0xB8, 0x14, 0xDE, 0x5E, 0x0B, 0xDB,
	0xE0, 0x32, 0x3A, 0x0A, 0x49, 0x06, 0x24, 0x5C, 0xC2, 0xD3, 0xAC, 0x62, 0x91, 0x95, 0xE4, 0x79,
	0xE7, 0xC8, 0x37, 0x6D, 0x8D, 0xD5, 0x4E, 0xA9, 0x6C, 0x56, 0xF4, 0xEA, 0x65, 0x7A, 0xAE, 0x08,
	0xBA, 0x78, 0x25, 0x2E, 0x1C, 0xA6, 0xB4, 0xC6, 0xE8, 0xDD, 0x74, 0x1F, 0x4B, 0xBD, 0x8B, 0x8A,
	0x70, 0x3E, 0xB5, 0x66, 0x48, 0x03, 0xF6, 0x0E, 0x61, 0x35, 0x57, 0xB9, 0x86, 0xC1, 0x1D, 0x9E,
	0xE1, 0xF8, 0x98, 0x11, 0x69, 0xD9, 0x8E, 0x94, 0x9B, 0x1E, 0x87, 0xE9, 0xCE, 0x55, 0x28, 0xDF,
	0x8C, 0xA1, 0x89, 0x0D, 0xBF, 0xE6, 0x42, 0x68, 0x41, 0x99, 0x2D, 0x0F, 0xB0, 0x54, 0xBB, 0x16,
}

var invSBox = []byte{
	0x52, 0x09, 0x6A, 0xD5, 0x30, 0x36, 0xA5, 0x38, 0xBF, 0x40, 0xA3, 0x9E, 0x81, 0xF3, 0xD7, 0xFB,
	0x7C, 0xE3, 0x39, 0x82, 0x9B, 0x2F, 0xFF, 0x87, 0x34, 0x8E, 0x43, 0x44, 0xC4, 0xDE, 0xE9, 0xCB,
	0x54, 0x7B, 0x94, 0x32, 0xA6, 0xC2, 0x23, 0x3D, 0xEE, 0x4C, 0x95, 0x0B, 0x42, 0xFA, 0xC3, 0x4E,
	0x08, 0x2E, 0xA1, 0x66, 0x28, 0xD9, 0x24, 0xB2, 0x76, 0x5B, 0xA2, 0x49, 0x6D, 0x8B, 0xD1, 0x25,
	0x72, 0xF8, 0xF6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xD4, 0xA4, 0x5C, 0xCC, 0x5D, 0x65, 0xB6, 0x92,
	0x6C, 0x70, 0x48, 0x50, 0xFD, 0xED, 0xB9, 0xDA, 0x5E, 0x15, 0x46, 0x57, 0xA7, 0x8D, 0x9D, 0x84,
	0x90, 0xD8, 0xAB, 0x00, 0x8C, 0xBC, 0xD3, 0x0A, 0xF7, 0xE4, 0x58, 0x05, 0xB8, 0xB3, 0x45, 0x06,
	0xD0, 0x2C, 0x1E, 0x8F, 0xCA, 0x3F, 0x0F, 0x02, 0xC1, 0xAF, 0xBD, 0x03, 0x01, 0x13, 0x8A, 0x6B,
	0x3A, 0x91, 0x11, 0x41, 0x4F, 0x67, 0xDC, 0xEA, 0x97, 0xF2, 0xCF, 0xCE, 0xF0, 0xB4, 0xE6, 0x73,
	0x96, 0xAC, 0x74, 0x22, 0xE7, 0xAD, 0x35, 0x85, 0xE2, 0xF9, 0x37, 0xE8, 0x1C, 0x75, 0xDF, 0x6E,
	0x47, 0xF1, 0x1A, 0x71, 0x1D, 0x29, 0xC5, 0x89, 0x6F, 0xB7, 0x62, 0x0E, 0xAA, 0x18, 0xBE, 0x1B,
	0xFC, 0x56, 0x3E, 0x4B, 0xC6, 0xD2, 0x79, 0x20, 0x9A, 0xDB, 0xC0, 0xFE, 0x78, 0xCD, 0x5A, 0xF4,
	0x1F, 0xDD, 0xA8, 0x33, 0x88, 0x07, 0xC7, 0x31, 0xB1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xEC, 0x5F,
	0x60, 0x51, 0x7F, 0xA9, 0x19, 0xB5, 0x4A, 0x0D, 0x2D, 0xE5, 0x7A, 0x9F, 0x93, 0xC9, 0x9C, 0xEF,
	0xA0, 0xE0, 0x3B, 0x4D, 0xAE, 0x2A, 0xF5, 0xB0, 0xC8, 0xEB, 0xBB, 0x3C, 0x83, 0x53, 0x99, 0x61,
	0x17, 0x2B, 0x04, 0x7E, 0xBA, 0x77, 0xD6, 0x26, 0xE1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0C, 0x7D,
}

func subBytes(s [][]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s[i][j] = sBox[s[i][j]]
		}
	}
}

func invSubBytes(s [][]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s[i][j] = invSBox[s[i][j]]
		}
	}
}

func shiftRows(s [][]byte) {
	s[0][1], s[1][1], s[2][1], s[3][1] = s[1][1], s[2][1], s[3][1], s[0][1]
	s[0][2], s[1][2], s[2][2], s[3][2] = s[2][2], s[3][2], s[0][2], s[1][2]
	s[0][3], s[1][3], s[2][3], s[3][3] = s[3][3], s[0][3], s[1][3], s[2][3]
}

func invShiftRows(s [][]byte) {
	s[0][1], s[1][1], s[2][1], s[3][1] = s[3][1], s[0][1], s[1][1], s[2][1]
	s[0][2], s[1][2], s[2][2], s[3][2] = s[2][2], s[3][2], s[0][2], s[1][2]
	s[0][3], s[1][3], s[2][3], s[3][3] = s[1][3], s[2][3], s[3][3], s[0][3]
}

func addRoundKey(s [][]byte, k [][]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s[i][j] ^= k[i][j]
		}
	}
}

func xtime(a byte) byte {
	if a&0x80 != 0 {
		return ((a << 1) ^ 0x1B) & 0xFF
	}
	return a << 1
}

func mixSingleColumn(a []byte) {
	t := byte(a[0] ^ a[1] ^ a[2] ^ a[3])
	u := byte(a[0])
	a[0] ^= t ^ xtime(a[0]^a[1])
	a[1] ^= t ^ xtime(a[1]^a[2])
	a[2] ^= t ^ xtime(a[2]^a[3])
	a[3] ^= t ^ xtime(a[3]^u)
}

func mixColumns(s [][]byte) {
	for i := 0; i < 4; i++ {
		mixSingleColumn(s[i])
	}
}

func invMixColumns(s [][]byte) {
	for i := 0; i < 4; i++ {
		u := xtime(xtime(s[i][0] ^ s[i][2]))
		v := xtime(xtime(s[i][1] ^ s[i][3]))
		s[i][0] ^= u
		s[i][1] ^= v
		s[i][2] ^= u
		s[i][3] ^= v
	}

	mixColumns(s)
}

func bytes2matrix(text []byte) [][]byte {
	res := [][]byte{}
	for i := 0; i < 4; i++ {
		res = append(res, []byte{})
		res[i] = append(res[i], text[4*i:4*i+4]...)
	}
	return res
}

func matrix2bytes(matrix [][]byte) []byte {
	res := []byte{}
	for i := 0; i < 4; i++ {
		res = append(res, matrix[i]...)
	}
	return res
}

func xorBytes(a []byte, b []byte) []byte {
	res := []byte{}
	for i := 0; i < len(a); i++ {
		res = append(res, a[i]^b[i])
	}
	return res
}

func pad(plaintext []byte) []byte {
	if len(plaintext)%16 == 0 {
		return plaintext
	}

	paddingLen := 16 - (len(plaintext) % 16)
	padding := []byte{}
	for i := 0; i < paddingLen; i++ {
		padding = append(padding, byte(paddingLen))
	}
	plaintext = append(plaintext, padding...)
	return plaintext
}

func unpad(plaintext []byte) []byte {
	if len(plaintext)%16 == 0 {
		return plaintext
	}

	paddingLen := int(plaintext[len(plaintext)-1])
	message := plaintext[:len(plaintext)-1-paddingLen]
	return message
}

// AES implement 128 bit aes algorithm
type AES struct {
	nRounds int
	key     [][]byte
}

// AESInit inits AES (wow)
func AESInit(key []byte) AES {
	aes := AES{
		nRounds: 10,
		key:     [][]byte{},
	}
	for i := 0; i < len(key)/4; i++ {
		aes.key = append(aes.key, key[4*i:4*i+4])
	}
	return aes
}

func (aes *AES) encryptBlock(plaintext []byte) []byte {
	plainState := bytes2matrix(plaintext)

	addRoundKey(plainState, aes.key)

	for i := 0; i < aes.nRounds; i++ {
		subBytes(plainState)
		shiftRows(plainState)
		mixColumns(plainState)
		addRoundKey(plainState, aes.key)
	}

	subBytes(plainState)
	shiftRows(plainState)
	addRoundKey(plainState, aes.key)

	return matrix2bytes(plainState)
}

func (aes *AES) decryptBlock(ciphertext []byte) []byte {
	cipherState := bytes2matrix(ciphertext)

	addRoundKey(cipherState, aes.key)
	invShiftRows(cipherState)
	invSubBytes(cipherState)

	for i := 0; i < aes.nRounds; i++ {
		addRoundKey(cipherState, aes.key)
		invMixColumns(cipherState)
		invShiftRows(cipherState)
		invSubBytes(cipherState)
	}

	addRoundKey(cipherState, aes.key)

	return matrix2bytes(cipherState)
}

func (aes *AES) encryptCBC(plaintext []byte, iv []byte) []byte {
	plaintext = pad(plaintext)

	blocks := []byte{}
	previous := iv
	for i := 0; i < len(plaintext); i += 16 {
		plaintextBlock := plaintext[i : i+16]
		block := aes.encryptBlock(xorBytes(plaintextBlock, previous))
		blocks = append(blocks, block...)
		previous = block
	}

	return blocks
}

func (aes *AES) decryptCBC(ciphertext []byte, iv []byte) []byte {
	blocks := []byte{}
	previous := iv
	for i := 0; i < len(ciphertext); i += 16 {
		ciphertextBlock := ciphertext[i : i+16]
		blocks = append(blocks, xorBytes(previous, aes.decryptBlock(ciphertextBlock))...)
		previous = ciphertextBlock
	}

	return unpad(blocks)
}

func encrypt(key []byte, plaintext []byte) []byte {
	iv := make([]byte, ivSize)
	rand.Read(iv)
	aes := AESInit(key)
	ciphertext := aes.encryptCBC(plaintext, iv)

	return append(iv, ciphertext...)
}

func decrypt(key []byte, ciphertext []byte) []byte {
	iv, ciphertext := ciphertext[:ivSize], ciphertext[ivSize:]
	aes := AESInit(key)
	return aes.decryptCBC(ciphertext, iv)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 0644)
}

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

func main() {
	key := []byte("dyftvghdsjofgios")
	// data := []byte("ssdfsfsfsfsfjshsbfjsfsdklfjnsgsdfsdfsjkfsdjf")
	// encrypted := encrypt(key, data)
	// decrypted := decrypt(key, encrypted)
	// fmt.Println(string(data))
	// fmt.Println(string(encrypted))
	// fmt.Println(string(decrypted))
	fmt.Print("Enter file name: ")
	filename := readline()

	data, err := readFromFile(filename)
	if err != nil {
		panic(err)
	}

	encrypted := encrypt(key, data)
	// fmt.Println(string(encrypted))
	writeToFile(string(encrypted), filename+".encoded")
	encrypted, err = readFromFile(filename + ".encoded")
	decrypted := decrypt(key, encrypted)
	// fmt.Println(string(decrypted))
	writeToFile(string(decrypted), filename+".decoded")
}
