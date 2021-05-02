package payments

import (
	"bytes"
	"encoding/json"
	"log"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func Encrypt(data interface{}, key string) string {
	passphrase := []byte(key)

	message, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(nil)
	w, err := armor.Encode(buffer, openpgp.PublicKeyType, make(map[string]string))
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := openpgp.SymmetricallyEncrypt(w, passphrase, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = plaintext.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	plaintext.Close()
	w.Close()
	return buffer.String()
}
