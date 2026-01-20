package example

import (
	"fmt"
	"testing"

	"github.com/enjoy322/wechatpay-b2b/client"
)

// Test
func TestClient(t *testing.T) {
	c, err := client.NewClient(client.Options{
		AccessToken: "your_access_token",
	})
	if err != nil {
		t.Fatal(err)
	}

	uri := "/retail/B2b/getorder"
	appkey := "12345"

	postBody := "{\"mchid\": \"1230000109\",\"out_trade_no\": \"1217752501201407033233368018\"}"
	expectedPaySig := "2cb3e9fda4da9774f0c6f747bbf6c56e9045969d9921db7344fffdc945c1dc23"

	paySig := c.GetPaySig(uri, []byte(postBody), appkey)
	fmt.Printf("paySig: %s\n", paySig)
	if paySig != expectedPaySig {
		t.Fatalf("unexpected paySig: %s", paySig)
	}

}
