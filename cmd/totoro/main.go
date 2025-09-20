package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/totoro-paradise/goapp/internal/data"
	"github.com/totoro-paradise/goapp/internal/types"
	"github.com/totoro-paradise/goapp/internal/utils"
)

func main() {
	target := time.Now().Add(150 * time.Millisecond)
	utils.WaitUntilRun(target, nil)
	fmt.Println("开始运行...")

	cipherText, err := encryptSamplePayload()
	if err != nil {
		log.Fatalf("encrypt sample payload: %v", err)
	}

	payload, err := utils.DecryptRequestContent(cipherText)
	if err != nil {
		log.Fatalf("decrypt payload: %v", err)
	}
	payloadJSON, _ := json.Marshal(payload)
	fmt.Printf("解密结果: %s\n", payloadJSON)

	route, err := utils.GenerateRoute(0.2, sampleTask())
	if err != nil {
		log.Fatalf("generate route: %v", err)
	}
	fmt.Printf("生成路线: 长度%.2f公里, 点位数量%d\n", route.Distance, len(route.MockRoute))
}

func encryptSamplePayload() (string, error) {
	block, _ := pem.Decode([]byte(data.PublicKey))
	if block == nil {
		return "", fmt.Errorf("invalid public key pem")
	}
	publicKeyAny, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("parse public key: %w", err)
	}
	publicKey, ok := publicKeyAny.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("unexpected public key type: %T", publicKeyAny)
	}

	payload := map[string]any{
		"message": "hello from totoro",
		"count":   3,
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, raw)
	if err != nil {
		return "", fmt.Errorf("encrypt payload: %w", err)
	}

	return base64.StdEncoding.EncodeToString(cipher), nil
}

func sampleTask() types.RunPoint {
	return types.RunPoint{
		TaskID:    "demo",
		PointID:   "start",
		PointName: "Demo Route",
		Longitude: "116.397128",
		Latitude:  "39.916527",
		PointList: []types.Point{
			{Longitude: "116.397128", Latitude: "39.916527"},
			{Longitude: "116.397500", Latitude: "39.917000"},
			{Longitude: "116.398000", Latitude: "39.916800"},
			{Longitude: "116.397600", Latitude: "39.916200"},
		},
	}
}
