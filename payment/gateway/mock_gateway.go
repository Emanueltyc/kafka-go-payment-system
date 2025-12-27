package gateway

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"payment/constants"
	"payment/ports"

	"github.com/google/uuid"
)

type MockGateway struct{}

func NewMockGateway() MockGateway {
	return MockGateway{}
}

type webhookPayload struct {
	EventID   string      `json:"event_id"`
	Data      webhookData `json:"data"`
	CreatedAt time.Time   `json:"created_at"`
}

type webhookData struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
	Reason    string `json:"reason,omitempty"`
}

func (m *MockGateway) Charge(ctx context.Context, req *ports.ChargeRequest) (*ports.ChargeResult, error) {
	go func() {
		time.Sleep(500 * time.Millisecond)
		callWebhook(req)
	}()

	return &ports.ChargeResult{
		PaymentID: req.PaymentID,
		Status:    constants.Status["pending"],
	}, nil
}

func callWebhook(req *ports.ChargeRequest) {
	webhookData := &webhookPayload{
		EventID:   uuid.NewString(),
		Data:      webhookData{},
		CreatedAt: time.Now(),
	}

	webhookData.Data.PaymentID = req.PaymentID

	switch req.PaymentToken {
	case "tok_nok":
		webhookData.Data.Status = constants.Status["rejected"]
		webhookData.Data.Reason = "insufficient_funds"
	case "tok_ok":
		webhookData.Data.Status = constants.Status["approved"]
	}

	marshaledBody, _ := json.Marshal(webhookData)
	payload := bytes.NewBuffer(marshaledBody)

	var client http.Client

	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/webhook/v1/payment", payload)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-Mock-Signature", generateSignature(marshaledBody, os.Getenv("WEBHOOK_SECRET")))

	go func() {
		for range(3) {
			client.Do(request)
			time.Sleep(time.Second + 10)
		}
	}()
}

func generateSignature(payload []byte, secret string) string {
	timestamp := time.Now().Unix()

	signedPayload := fmt.Sprintf("%d.%s", timestamp, payload)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signedPayload))
	signature := hex.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("t=%d,v1=%s", timestamp, signature)
}
