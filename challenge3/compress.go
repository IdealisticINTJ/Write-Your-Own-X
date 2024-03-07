
package huffman

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// Compress takes a string and compresses it using Huffman coding
func Compress(content string) ([]byte, error) {
	inputContent := strings.Split(content, "")

	freqTable := getFreq(inputContent)

	root := createRootNode(freqTable)

	prefixCodes := generatePrefixCodes(root)

	bitString, paddingAdded := stringToBitString(inputContent, prefixCodes)

	compressed, err := bitStringToByteArray(bitString)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to convert bitstring to bytes: %w", err)
	}

	header := encodeHeader(*root, paddingAdded)

	res := append(header, compressed...)

	return res, nil
}

// Decompress takes a byte array and decompresses it using Huffman coding
func Decompress(content []byte) (string, error) {
	pad, decodedTree, byteString, err := decodeContent(content)
	if err != nil {
		return "", fmt.Errorf("failed to decode content: %w", err)
	}

	decodedUncompressed := byteArrayToBitString(byteString)
	decoded := decode(decodedUncompressed, decodedTree, int(pad))

	return decoded, nil
}

func generatePrefixCodes(item *Item) map[string]string {
	prefixCodes := map[string]string{}
	var prefix string
	getPrefixMap(item, &prefix, prefixCodes)
	return prefixCodes
}

func createRootNode(f FrequencyTable) *Item {
	pq := NewPriorityQueue(f)
	root := pq.generateNodeTree()
	return &root
}

func encodeHeader(root Item, padding int) []byte {
	serializedTree := Serialize(&root)

	treeLength := uint64(len(serializedTree))

	header := []byte{}
	header = binary.LittleEndian.AppendUint64(header, treeLength)

	padHeader := binary.LittleEndian.AppendUint32([]byte{}, uint32(padding))
	header = append(header, padHeader...)

	header = append(header, serializedTree...)

	return header
}

func decodeContent(content []byte) (uint32, *Item, []byte, error) {
	var treeLength uint64
	var paddingLength uint32
	err := binary.Read(bytes.NewReader(content[:8]), binary.LittleEndian, &treeLength)
	if err != nil {
		return 0, nil, []byte{}, fmt.Errorf("failed to read binary: %w", err)
	}

	err = binary.Read(bytes.NewReader(content[8:12]), binary.LittleEndian, &paddingLength)
	if err != nil {
		return 0, nil, []byte{}, fmt.Errorf("failed to read binary: %w", err)
	}

	ftStartIdx := 12
	ftEndIdx := treeLength

	treeText := content[ftStartIdx : ftEndIdx+uint64(ftStartIdx)]

	tree := Deserialize(string(treeText))

	return paddingLength, tree, content[ftEndIdx+uint64(ftStartIdx):], nil

}

func decode(s string, node *Item, padding int) string {
	s = s[:len(s)-padding]
	var result string

	currentNode := node
	if currentNode.LeftNode == nil || currentNode.RightNode == nil {
		return ""
	}
	for _, v := range s {
		if v == '0' {
			currentNode = currentNode.LeftNode
		} else if v == '1' {
			currentNode = currentNode.RightNode
		}

		if currentNode == nil {
			continue
		}
		if currentNode.isLeafNode() {
			result += currentNode.Value
			currentNode = node
		}
	}
	return result
}

func stringToBitString(in []string, prefixCodes map[string]string) (string, int) {
	bitString := ""

	for _, c := range in {
		bitString += prefixCodes[string(c)]
	}

	paddingAdded := 0
	for len(bitString)%8 != 0 {
		bitString += "0"
		paddingAdded++
	}

	return bitString, paddingAdded
}

func bitStringToByteArray(bitString string) ([]byte, error) {
	lenBits := len(bitString) / 8
	if len(bitString)%8 != 0 {
		return []byte{}, fmt.Errorf("provided string: %s is not valid 8 length", bitString)
	}

	out := make([]byte, lenBits)

	for i := 0; i < lenBits; i++ {
		start := i * 8
		end := start + 8

		byteValue := bitString[start:end]

		int, err := strconv.ParseUint(byteValue, 2, 8)
		if err != nil {
			return []byte{}, fmt.Errorf("failed to parse to uint: %w", err)
		}

		out[i] = byte(int)
	}
	return out, nil
}

func byteArrayToBitString(b []byte) string {
	var stringByte string
	for _, char := range b {
		stringByte += fmt.Sprintf("%08b", char)
	}
	return stringByte
}

func getPrefixMap(i *Item, prefix *string, prexfixCodes map[string]string) {
	if i == nil {
		return
	} else {
		zero := fmt.Sprintf("%s0", *prefix)
		one := fmt.Sprintf("%s1", *prefix)
		getPrefixMap(i.LeftNode, &zero, prexfixCodes)
		getPrefixMap(i.RightNode, &one, prexfixCodes)

		if i.isLeafNode() {
			prexfixCodes[i.Value] = *prefix
		}
	}
}

func getFreq(s []string) FrequencyTable {
	m := make(map[string]int, len(s))
	for _, c := range s {
		m[c]++
	}
	return m
}
