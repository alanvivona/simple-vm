package link

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type header struct {
	Signature    [2]byte  // [2]byte{0x59,0x59}
	Sha512sum    [64]byte // Fingerprint of publick key used for encryption
	Size         uint64
	CodeMark     string  // "code"
	EndDelimiter [2]byte // [2]byte{0x60,0x60}
}

func (h header) serialize() []byte {
	serial := []byte{}
	serial = append(serial, h.Signature[:]...)
	serial = append(serial, h.Sha512sum[:]...)
	size := make([]byte, 8)
	binary.BigEndian.PutUint64(size, h.Size)
	serial = append(serial, size...)
	serial = append(serial, h.CodeMark...)
	return append(serial, h.EndDelimiter[:]...)
}

type fileFormat struct {
	sectionHeader header
	sectionCode   []byte
}

func (ff fileFormat) serialize() []byte {
	return append(ff.sectionHeader.serialize(), ff.sectionCode...)
}

func getDefaultHeader() header {
	return header{
		Signature:    [2]byte{0x59, 0x59},
		CodeMark:     "code",
		EndDelimiter: [2]byte{0x5a, 0x5a},
	}
}

func Link(in []byte) ([]byte, error) {
	h := getDefaultHeader()

	h.Sha512sum = sha512.Sum512(in)
	logrus.Info("Code checksum:\n%s", checksum2String(h.Sha512sum[:]))

	bytecodeSize := uint64(len(in))
	if bytecodeSize < 1 {
		return nil, errors.New("Provided binary section is empty")
	}
	h.Size = bytecodeSize
	logrus.Infof("Code size: %d (0x%x)", h.Size, h.Size)

	return fileFormat{sectionHeader: h, sectionCode: in}.serialize(), nil
}

func ExtractExecutable(in []byte) ([]byte, error) {
	h := getDefaultHeader()

	// signature
	start := 0
	end := start + len(h.Signature)
	inSignature := in[start:end]
	if !bytes.Equal(h.Signature[:], inSignature) {
		return []byte{}, fmt.Errorf("Broken header: Signature %s does not match %s", inSignature, h.Signature)
	}
	logrus.Info("Signature: OK")

	// get sum, can't check yet
	start = end
	end = start + len(h.Sha512sum)
	copy(h.Sha512sum[:], in[start:end])

	// get size, can't check yet
	start = end
	end = start + 8
	inSize := in[start:end]
	h.Size = binary.BigEndian.Uint64(inSize)
	if h.Size < 1 {
		return []byte{}, fmt.Errorf("Broken header: Empty size field")
	}

	// code mark
	start = end
	end = start + len(h.CodeMark)
	inCodeMark := in[start:end]
	if !bytes.Equal([]byte(h.CodeMark), inCodeMark) {
		return []byte{}, fmt.Errorf("Broken header: No code mark found. Expected '%s'. Found '%s'", h.CodeMark, inCodeMark)
	}
	logrus.Info("Code mark: OK")

	// header end delimiter
	start = end
	end = start + len(h.EndDelimiter)
	inEndDelimiter := in[start:end]
	if !bytes.Equal(h.EndDelimiter[:], inEndDelimiter) {
		return []byte{}, fmt.Errorf("Broken header: Not end delimiter found. Expected '%s'. Found '%s'", h.EndDelimiter, inEndDelimiter)
	}
	logrus.Info("Header end delimiter: OK")

	start = end
	end = start + int(h.Size)
	codeSection := in[start:]
	codeSectionSize := len(codeSection)
	if int(h.Size) != codeSectionSize {
		return []byte{}, fmt.Errorf("Broken header: Size field value %d (0x%x) doesn't match actual code section size %d (0x%x)", h.Size, h.Size, codeSectionSize, codeSectionSize)
	}
	logrus.Info("Code section size: OK")

	// get code and compare with checksum
	actualSum := sha512.Sum512(codeSection)
	if !bytes.Equal(h.Sha512sum[:], actualSum[:]) {
		expectedSumStr := checksum2String(h.Sha512sum[:])
		actualSumStr := checksum2String(actualSum[:])
		return []byte{}, fmt.Errorf(`
			Broken header: Checksum does not match
			Expected:
			%s
			Got:
			%s
		`, expectedSumStr, actualSumStr)
	}
	logrus.Info("Checksum: OK")

	return codeSection, nil
}

func checksum2String(checksum []byte) string {
	hexStr := make([]byte, hex.EncodedLen(len(checksum)))
	hex.Encode(hexStr, checksum[:])
	return string(hexStr)
}
