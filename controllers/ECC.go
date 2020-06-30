package controllers

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"

	crypt_rand "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"

	"math/rand"
	"os"
	"strings"
	"time"
	//"golang.org/x/crypto/ripemd160"
)

func Initkeys() (error, *GKey) {
	randstring := RandomString(41, "0aA")
	gkey, err := MakeNewKey(randstring)
	if err != nil {
		fmt.Println("err in MakeNewKey...")
		return err, nil
	}
	privateKey = gkey.GetPrivKey()
	pri_string := ByteToString(gkey.GetPrivKey())
	fmt.Println("My privateKey is :", pri_string)
	publicKey = gkey.GetPubKey()
	pub_string := ByteToStringpub(gkey.GetPubKey())
	fmt.Println("My publickKey is :", pub_string)
	return nil, gkey
}
func MakeNewKey(randKey string) (*GKey, error) {
	var err error
	var gkey GKey
	var curve elliptic.Curve

	lenth := len(randKey)
	fmt.Println(lenth)
	if lenth < 224/8+8 {
		err = errors.New("RandKey is too short. It mast be longer than 36 bytes.")
		return &gkey, err
	} else if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth > 256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}

	private, err := ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		log.Panic(err)
	}
	gkey = GKey{private, private.PublicKey}
	return &gkey, nil
}
func RandomString(randLength int, randType string) (result string) {
	var num string = "0123456789"
	var lower string = "abcdefghijklmnopqrstuvwxyz"
	var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}
	var str = b.String()
	var strLen = len(str)
	if strLen == 0 {
		result = ""
		return
	}
	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	fmt.Println(result)
	return
}
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	fmt.Println("3 - Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)")
	bcpy := append([]byte{ver}, b...)
	fmt.Println(ByteToString(bcpy))
	fmt.Println("================")

	/* Create a new SHA256 context */
	sha256H := sha256.New()

	/* SHA256 HASH #1 */
	fmt.Println("4 - Perform SHA-256 hash on the extended PIPEMD-160 result")
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)
	fmt.Println(ByteToString(hash1))
	fmt.Println("================")

	/* SHA256 HASH #2 */
	fmt.Println("5 - Perform SHA-256 hash on the result of the previous SHA-256 hash")
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)
	fmt.Println(ByteToString(hash2))
	fmt.Println("================")

	/* Append first four bytes of hash */
	fmt.Println("6 - Take the first 4 bytes of the second SHA-256 hash. This is the address chechsum")
	fmt.Println(ByteToString(hash2[0:4]))
	fmt.Println("================")

	fmt.Println("7 - Add the 4 checksum bytes from stage 7 at the end of extended PIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.")
	bcpy = append(bcpy, hash2[0:4]...)
	fmt.Println(ByteToString(bcpy))
	fmt.Println("================")

	/* Encode base58 string */
	s = b58encode(bcpy)

	/* For number  of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}
	fmt.Println("8 - Convet the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format")
	fmt.Println(s)
	fmt.Println("================")

	return s
}
func b58encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */
	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	x := new(big.Int).SetBytes(b)
	// Initialize
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x /58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}
	return s
}
func (k GKey) Sign(text []byte) (string, error) {
	r, s, err := ecdsa.Sign(crypt_rand.Reader, k.PrivateKey, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

func Verify(text []byte, signature string, pubKey *ecdsa.PublicKey) (bool, error) {
	rint, sint, err := getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(pubKey, text, &rint, &sint)
	return result, nil
}

func getSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error," + err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}

func GenerateEccKey() error {
	randKey := RandomString(40, "0aA")
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), strings.NewReader(randKey))
	if err != nil {
		return err
	}
	//此处使用509
	private, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	//pem
	block := pem.Block{
		Type:  "esdsa private key",
		Bytes: private,
	}
	file, err := os.Create("./eccprivateKey.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}
	file.Close()
	public := privateKey.PublicKey
	publicKey, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return err
	}
	//pem
	public_block := pem.Block{
		Type:  "ecdsa public key",
		Bytes: publicKey,
	}
	file, err = os.Create("./eccpublicKey.pem")
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(file, &public_block)
	if err != nil {
		return err
	}
	return nil
}

func EccSignature(sourceData []byte, privateKeyFilePath string) ([]byte, []byte) {
	//1，打开私钥文件，读出内容
	file, err := os.Open(privateKeyFilePath)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	hashText := sha1.Sum(sourceData)
	r, s, err := ecdsa.Sign(strings.NewReader("rand"), privateKey, hashText[:])
	if err != nil {
		panic(err)
	}
	rText, err := r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err := s.MarshalText()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return rText, sText
}

func EccVerify(rText, sText, sourceData []byte, publicKeyFilePath string) bool {
	file, err := os.Open(publicKeyFilePath)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	publicStream, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicStream.(*ecdsa.PublicKey)
	hashText := sha1.Sum(sourceData)
	var r, s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)
	res := ecdsa.Verify(publicKey, hashText[:], &r, &s)
	defer file.Close()
	return res
}

// func GetEccprivatekey() []byte {
// 	file, err := os.Open(privateKeyFilePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	info, err := file.Stat()
// 	buf := make([]byte, info.Size())
// 	file.Read(buf)
// 	return buf
// }
// func GetEccpublickey() []byte {
// 	file, err := os.Open(publicKeyFilePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	info, err := file.Stat()
// 	buf := make([]byte, info.Size())
// 	file.Read(buf)
// 	return buf
// }
