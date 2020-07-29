package merkletree

import (
	"testing"
)

func TestNewTxTree(t *testing.T) {
	myTx := "{\"from\":\"andrej\",\"to\":\"caesar\",\"value\":102}"

	txs := []Content{
		TestSHA256Content{
			x: "{\"from\":\"andrej\",\"to\":\"babayaga\",\"value\":101}",
		},
		TestSHA256Content{
			x: myTx,
		},
		TestSHA256Content{
			x: "{\"from\":\"andrej\",\"to\":\"daenerys\",\"value\":103}",
		},
	}

	tree, err := NewTree(txs)
	if err != nil {
		t.Fatal(err)
	}

	isTxInBlock, err := tree.VerifyContent(TestSHA256Content{x: myTx})
	if err != nil {
		t.Fatal(err)
	}

	if !isTxInBlock {
		t.Fatal("TX should be stored in the tree")
	}
}
