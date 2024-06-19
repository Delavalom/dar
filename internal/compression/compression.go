package compression

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

type Compressor interface {
	WriteFile(content []byte, key string)
	ReadAndWriteDecompressedFile(key, path string)
}

type ZLibCompression struct{}

func New() Compressor {
	return &ZLibCompression{}
}

// WriteFile compresses the content and writes it to .dar/ objects store
func (z *ZLibCompression) WriteFile(content []byte, key string) {
	// creating a new file in the .dar/objects directory with the hash as the name
	file, err := os.Create(fmt.Sprintf(".dar/objects/%s", key))
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer

	w := zlib.NewWriter(&buf)
	if _, err := w.Write(content); err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}

	if _, err := io.Copy(file, &buf); err != nil {
		panic(err)
	}
}

// ReadAndWriteFile reads the compressed file and writes it to the repository file system
func (z *ZLibCompression) ReadAndWriteDecompressedFile(key, path string) {
	compressedFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer compressedFile.Close()

	reader, err := zlib.NewReader(compressedFile)
	if err != nil {
		panic("Reading contentFile content: " + err.Error())
	}
	defer reader.Close()

	file, err := os.Create(key)
	if err != nil {
		panic("Creating file path: " + err.Error())
	}
	defer file.Close()

	if _, err = io.Copy(file, reader); err != nil {
		panic("copy from compressed file to repository file system: " + err.Error())
	}
}
