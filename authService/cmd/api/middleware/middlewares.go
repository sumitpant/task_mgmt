package middleware

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

/*
* Middle ware function that attaches header to request,
   returns a http Handler which in turns returns a handler
   which redirect to other handler in main file
*/

func AddHeaders(headers http.Header) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(headers) > 0 {
				for key, value := range headers {
					w.Header()[key] = value
				}
			}
			next.ServeHTTP(w, r)
		})

	}

}

func padPlainText(plainText []byte,blockSize int) []byte {
	padding:=blockSize-(len(plainText)%blockSize);
	padText:=bytes.Repeat([]byte{byte(padding)},padding)
	return append(plainText,padText...);

}

func Encrypt(text string) (string, error){
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte(text);
	plaintext=padPlainText(plaintext,aes.BlockSize);

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	fmt.Println(aes.BlockSize);
	fmt.Println(len(plaintext))
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "",err;
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "",err;
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	return string(ciphertext),nil;
}



