package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func Signstr(message, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	//return fmt.Sprintf("%x", base64.StdEncoding.EncodeToString((h.Sum(nil))))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println(Signstr([]byte("GETcvm.api.qcloud.com/v2/index.php?Action=DescribeInstances&InstanceIds.0=ins-09dx96dg&Nonce=11886&Region=ap-guangzhou&SecretId=AKIDz8krbsJ5yKBZQpn74WFkmLPx3gnPhESA&SignatureMethod=HmacSHA256&Timestamp=1465185768"), []byte("Gu5t9xGARNpq86cd98joQYCN3Cozk1qA")))
	//fmt.Println(Signstr("GETcvm.api.qcloud.com/v2/index.php?Action=DescribeInstances&InstanceIds.0=ins-09dx96dg&Nonce=11886&Region=ap-guangzhou&SecretId=AKIDz8krbsJ5yKBZQpn74WFkmLPx3gnPhESA&SignatureMethod=HmacSHA256&Timestamp=1465185768", "Gu5t9xGARNpq86cd98joQYCN3Cozk1qA"))
}
