package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// generateId() generates random ID for each peer of size 32 bytes
func generateID() string {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	return hex.EncodeToString(buf)
}

// hashKey() hashes the key name of the file
func hashKey(key string) string {
	hash := md5.Sum([]byte(key))
	return hex.EncodeToString(hash[:])
}

// newEncryptionKey() generates unique private key that will be used to encrypt file content data
func newEncryptionKey() []byte {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	return buf
}

// copyStream() takes a stream of data from src and encryptes/decryptes it with stream and writes to dst sequencially and return the total of bytes of written data in bytes
func copyStream(stream cipher.Stream, blockSize int, src io.Reader, dst io.Writer) (int, error) {
	var (
		buf = make([]byte, 32*1024) // max size chunk
		nw  = blockSize             // already appended iv in the file
	)

	for {
		n, err := src.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n]) // encyptes/decryptes using XOR property
			nn, err := dst.Write(buf[:n])
			if err != nil {
				return 0, err
			}
			nw += nn
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, nil
		}
	}
	return nw, nil
}

// copyDecrypt() extracts iv from the file and copies the decrypted from src to dst via key
func copyDecrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	// Read the IV from the passed io.Reader whose size will be the block.BlockSize()
	iv := make([]byte, block.BlockSize()) // 16 bytes
	if _, err := src.Read(iv); err != nil {
		return 0, err
	}

	stream := cipher.NewCTR(block, iv)
	return copyStream(stream, block.BlockSize(), src, dst)
}

// copyEncrypt() generates random iv from key for encryption and copies the encrypted data from src to dst
func copyEncrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	iv := make([]byte, block.BlockSize()) // 16 bytes
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return 0, err
	}

	// prepend iv to the file
	if _, err := dst.Write(iv); err != nil {
		return 0, err
	}

	stream := cipher.NewCTR(block, iv)
	return copyStream(stream, block.BlockSize(), src, dst)
}
