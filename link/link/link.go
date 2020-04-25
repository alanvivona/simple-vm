package link

import (
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/sirupsen/logrus"
)

type header struct {
	Magic            [2]byte  // [2]byte{0x59,0x59}
	Sha512sum        [64]byte // Fingerprint of publick key used for encryption
	Size             uint64
	CodeMark         string  // "code"
	PaddingDelimiter [2]byte // [2]byte{0x60,0x60}
}

func (h header) serialize() []byte {
	serial := []byte{}
	serial = append(serial, h.Magic[:]...)
	serial = append(serial, h.Sha512sum[:]...)
	size := make([]byte, 8)
	binary.BigEndian.PutUint64(size, h.Size)
	serial = append(serial, size...)
	serial = append(serial, h.CodeMark...)
	return append(serial, h.PaddingDelimiter[:]...)
}

type fileFormat struct {
	sectionHeader header
	sectionCode   []byte
}

func (ff fileFormat) serialize() []byte {
	return append(ff.sectionHeader.serialize(), ff.sectionCode...)
}

func Link(in []byte) ([]byte, error) {
	header := header{}

	header.Magic = [2]byte{0x59, 0x59}
	header.Sha512sum = sha512.Sum512(in)

	hexStr := make([]byte, hex.EncodedLen(len(header.Sha512sum)))
	hex.Encode(hexStr, header.Sha512sum[:])
	logrus.Infof("Code checksum:\n%s", hexStr)

	bytecodeSize := uint64(len(in))
	if bytecodeSize < 1 {
		return nil, errors.New("Provided binary section is empty")
	}
	header.Size = bytecodeSize
	logrus.Infof("Code size: %d (0x%x)", header.Size, header.Size)

	header.CodeMark = "code"

	header.PaddingDelimiter = [2]byte{0x5a, 0x5a}

	return fileFormat{sectionHeader: header, sectionCode: in}.serialize(), nil
}
