package openrtb_test

import (
	"encoding/json"
	"testing"

	"pkg/openrtb"
)

func TestSource_Unmarshal(t *testing.T) {
	expected := openrtb.Source{
		FinalSaleDecision: 1,
		TransactionID:     "transaction-id",
		PaymentChain:      "payment-chain",
		Ext:               json.RawMessage([]byte("{}")),
	}

	assertEqualJSON(t, "source", &expected)
}
