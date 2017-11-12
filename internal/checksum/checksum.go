package checksum

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// Calculate calculates SHA256 checksum from filename and content
func Calculate(filename, content io.Reader) (string, error) {
	hasher := sha256.New()

	buf := make([]byte, 8)

	for {
		n, err := filename.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hasher.Write(buf[:n])
	}

	for {
		n, err := content.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hasher.Write(buf[:n])
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
