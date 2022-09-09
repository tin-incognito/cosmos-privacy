package coin

import (
	"fmt"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/repos/operation"
	"testing"

	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
)

func TestOTAReceiver_FromAddress(t *testing.T) {
	privateKey := "112t8rnXSdPp6Ygsv5VUycUe4E4SHMn1eQFTFgni7HyExwh81CYecR3Uc2dvWrmSP3j8cgXYjxTwWGp442skQMH84BoFqEa1VmF5tFpbfmzA"

	keyWallet, err := wallet.Base58CheckDeserialize(privateKey)
	if err != nil {
		panic(err)
	}
	keySet := key.KeySet{}
	err = keySet.InitFromPrivateKeyByte(keyWallet.KeySet.PrivateKey)
	if err != nil {
		panic(err)
	}

	type fields struct {
		PublicKey *operation.Point
		TxRandom  TxRandom
	}
	type args struct {
		addr key.PaymentAddress
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Generate OTA",
			fields: fields{},
			args: args{
				addr: keySet.PaymentAddress,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recv := &OTAReceiver{
				PublicKey: tt.fields.PublicKey,
				TxRandom:  tt.fields.TxRandom,
			}
			if err := recv.FromAddress(tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("OTAReceiver.FromAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
			temp, err := recv.String()
			if err != nil {
				panic(err)
			}
			fmt.Println("OTAReceiver:", temp)
		})
	}
}
