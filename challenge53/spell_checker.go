package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"os"
)

type BloomFilter struct {
	bitArray []bool
	numBits  uint32
	numHash  uint16
}

func NewBloomFilter(numBits uint32, numHash uint16) *BloomFilter {
	return &BloomFilter{
		bitArray: make([]bool, numBits),
		numBits:  numBits,
		numHash:  numHash,
	}
}

func (bf *BloomFilter) Insert(item string) {
	for i := uint16(0); i < bf.numHash; i++ {
		index := bf.hash(item, i)
		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) Query(item string) bool {
	for i := uint16(0); i < bf.numHash; i++ {
		index := bf.hash(item, i)
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) hash(item string, index uint16) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(item))
	return hash.Sum32() % uint32(bf.numBits)
}

func main() {
	numBits := uint32(1000000)
	numHash := uint16(3)

	bf := NewBloomFilter(numBits, numHash)

	file, err := os.Open("dict.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		bf.Insert(word)
	}

	outFile, err := os.Create("words.bf")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	header := []interface{}{"CCBF", uint16(1), numHash, numBits}
	for _, value := range header {
		err := binary.Write(outFile, binary.BigEndian, value)
		if err != nil {
			fmt.Println("Error writing header to file:", err)
			return
		}
	}

	for _, bit := range bf.bitArray {
		var value uint8
		if bit {
			value = 1
		}
		err := binary.Write(outFile, binary.BigEndian, value)
		if err != nil {
			fmt.Println("Error writing bit array to file:", err)
			return
		}
	}

	file, err = os.Open("words.bf")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the header from the file
	var identifier [4]byte
	var version uint16
	var numHashLoaded uint16
	var numBitsLoaded uint32

	err = binary.Read(file, binary.BigEndian, &identifier)
	if err != nil {
		fmt.Println("Error reading identifier from file:", err)
		return
	}

	if string(identifier[:]) != "CCBF" {
		fmt.Println("Invalid file format")
		return
	}

	err = binary.Read(file, binary.BigEndian, &version)
	if err != nil {
		fmt.Println("Error reading version from file:", err)
		return
	}

	err = binary.Read(file, binary.BigEndian, &numHashLoaded)
	if err != nil {
		fmt.Println("Error reading numHash from file:", err)
		return
	}

	err = binary.Read(file, binary.BigEndian, &numBitsLoaded)
	if err != nil {
		fmt.Println("Error reading numBits from file:", err)
		return
	}

	// Check if loaded Bloom filter parameters match expected values
	if numHashLoaded != numHash || numBitsLoaded != numBits {
		fmt.Println("Mismatched Bloom filter parameters")
		return
	}

	// Read the bit array from the file and populate the Bloom filter
	var bit byte
	for i := uint32(0); i < numBits; i++ {
		err := binary.Read(file, binary.BigEndian, &bit)
		if err != nil {
			fmt.Println("Error reading bit array from file:", err)
			return
		}
		bf.bitArray[i] = bit == 1
	}

	// Step 5: Test words against Bloom filter
	testWords := []string{"hi", "hello", "word", "concurrency", "coding", "challenges", "imadethis", "up"}
	fmt.Println("These words are spelt wrong:")
	for _, word := range testWords {
		if !bf.Query(word) {
			fmt.Println(" ", word)
		}
	}
}
