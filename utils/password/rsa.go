// @Create   : 2023/4/21 10:57
// @Author   : yaho
// @Remark   : rsa加解密

package password

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"path/filepath"

	"Goldfinger/errors"
)

// ReadRSAPublicKeyFromFile 从文件读取密钥并转换为RSA公钥
func ReadRSAPublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.NewPWDDecodeError("读完RSA公钥文件失败：" + err.Error())
	}

	pubBlock, _ := pem.Decode(pemData)
	if pubBlock == nil {
		return nil, errors.NewPWDDecodeError("解码PEM格式的公钥失败：" + err.Error())
	}

	pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, errors.NewPWDDecodeError("解析RSA公钥失败，请检查您的密钥：" + err.Error())
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.NewPWDDecodeError("将publicKey转换为RSA publicKey类型失败")
	}

	return rsaPubKey, nil
}

// ReadRSAPrivateKeyFromFile 从文件读取密钥并转为RSA私钥
func ReadRSAPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	pubBlock, _ := pem.Decode(pemData)
	if pubBlock == nil {
		return nil, errors.NewPWDDecodeError("解码PEM格式的公钥失败：" + err.Error())
	}

	return x509.ParsePKCS1PrivateKey(pubBlock.Bytes)
}

// RSAPrivateDecode rsa私钥进行解密
func RSAPrivateDecode(pwd string) (string, error) {

	encryptedData, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return "", errors.NewPWDDecodeError("密码解码错误：" + err.Error())
	}

	keyDirPath, err := filepath.Abs("./config/key")
	privatePem, err := ReadRSAPrivateKeyFromFile(keyDirPath + "/private.pem")
	if err != nil {
		return "", errors.NewPWDDecodeError("读取RSA私钥失败：" + err.Error())
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privatePem, encryptedData, nil)
	if err != nil {
		return "", errors.NewPWDDecodeError("密码解密错误：" + err.Error())
	}

	return string(plaintext), nil
}

// RSAPublicEncode rsa公钥进行加密
func RSAPublicEncode(plaintext string) (string, error) {

	keyDirPath, err := filepath.Abs("./config/key")
	publicPem, err := ReadRSAPublicKeyFromFile(keyDirPath + "/public.pem")
	if err != nil {
		return "", errors.NewPWDEncodeError("读取RSA公钥失败：" + err.Error())
	}

	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicPem, []byte(plaintext), nil)
	if err != nil {
		return "", errors.NewPWDEncodeError("RSA公钥加密失败")
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
