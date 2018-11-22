package crypto

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

//the sigScript from TxID : fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4
var validSig = []byte{
	0x30,
	0x46, //sigLenth
	0x02,
	0x21,
	0x00, 0xc3, 0x52, 0xd3, 0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4,
	0xa6, 0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94, 0x70, 0xab,
	0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58, 0xe4, 0xeb, 0x5d, 0xce, 0x82,
	0x02,
	0x21,
	0x00, 0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62, 0x81, 0x9f,
	0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c, 0xf7, 0xb5, 0xee, 0x1a, 0xf1,
	0xeb, 0xcc, 0x60, 0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
}

func TestIsValidSignatureEncoding(t *testing.T) {
	ret := IsValidSignatureEncoding(validSig)
	assert.Equal(t, true, ret)

	sig1 := []byte{0x30, 0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca}
	ret = IsValidSignatureEncoding(sig1)
	assert.Equal(t, false, ret)

	sig2 := []byte{0x30, 0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00, 0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62, 0xbc, 0x1f,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48, 0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60, 0xee, 0x1a,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c, 0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60, 0xc3, 0xaf,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3,
	}
	ret = IsValidSignatureEncoding(sig2)
	assert.Equal(t, false, ret)

	sig3 := []byte{0x31,
		0x46, //sigLenth
		0x02, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig3)
	assert.Equal(t, false, ret)

	sig4 := []byte{0x30,
		0x47, //sigLenth
		0x02, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig4)
	assert.Equal(t, false, ret)

	sig5 := []byte{0x30,
		0x46, //sigLenth
		0x02, 0x52, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig5)
	assert.Equal(t, false, ret)

	sig6 := []byte{0x30,
		0x46, //sigLenth
		0x02, 0x22, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig6)
	assert.Equal(t, false, ret)

	sig7 := []byte{0x30,
		0x46, //sigLenth
		0x01, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig7)
	assert.Equal(t, false, ret)

	sig8 := []byte{0x30,
		0x3f, //sigLenth
		0x02, 0x00, 0x00, 0x3b, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9,
	}
	ret = IsValidSignatureEncoding(sig8)
	assert.Equal(t, false, ret)

	sig9 := []byte{0x30,
		0x46, //sigLenth
		0x02, 0x21, 0x80, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig9)
	assert.Equal(t, false, ret)

	sig10 := []byte{
		0x30,
		0x46, //sigLenth
		0x02, 0x21, 0x00, 0x00, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig10)
	assert.Equal(t, false, ret)

	sig11 := []byte{
		0x30,
		0x46, //sigLenth
		0x02, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x01, 0x21, 0x00,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig11)
	assert.Equal(t, false, ret)

	sig12 := []byte{
		0x30,
		0x25, //sigLenth
		0x02, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x00, 0x00,
		//sigType
	}
	ret = IsValidSignatureEncoding(sig12)
	assert.Equal(t, false, ret)

	sig13 := []byte{
		0x30,
		0x46, //sigLenth
		0x02, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x80,
		0x84, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig13)
	assert.Equal(t, false, ret)

	sig14 := []byte{
		0x30,
		0x46, //sigLenth
		0x02, 0x21, 0x00, 0xc3, 0x52, 0xd3,
		0xdd, 0x99, 0x3a, 0x98, 0x1b, 0xeb, 0xa4, 0xa6,
		0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca, 0x94,
		0x70, 0xab, 0xfc, 0xd5, 0x7d, 0xa9, 0x3b, 0x58,
		0xe4, 0xeb, 0x5d, 0xce, 0x82, 0x02, 0x21, 0x00,
		0x00, 0x07, 0x92, 0xbc, 0x1f, 0x45, 0x60, 0x62,
		0x81, 0x9f, 0x15, 0xd3, 0x3e, 0xe7, 0x05, 0x5c,
		0xf7, 0xb5, 0xee, 0x1a, 0xf1, 0xeb, 0xcc, 0x60,
		0x28, 0xd9, 0xcd, 0xb1, 0xc3, 0xaf, 0x77, 0x48,
		0x01, //sigType
	}
	ret = IsValidSignatureEncoding(sig14)
	assert.Equal(t, false, ret)

	ret = IsDefineHashtypeSignature(validSig)
	assert.Equal(t, false, ret)

	sighashType := []byte{}
	ret = IsDefineHashtypeSignature(sighashType)
	assert.Equal(t, false, ret)

	ret = IsDefineHashtypeSignature(sig12)
	assert.Equal(t, false, ret)
}

func TestParseSignature(t *testing.T) {
	InitSecp256()
	sig := validSig
	signature, err := ParseDERSignature(sig)
	if err != nil {
		t.Error(err)
	}
	sigByte := signature.Serialize()
	if !bytes.Equal(sigByte, sig) {
		t.Errorf("the new serialize signature %v should be equal origin sig %v: ", sigByte, sig)
	}
	ret := signature.EcdsaNormalize()
	assert.Equal(t, ret, true)

	secp256k1Context = nil
	_, err = ParseDERSignature(sig)
	if err == nil {
		t.Errorf("secp256k1Context is nil")
	}
}

func TestIsLowDERSignature(t *testing.T) {
	ret := CheckLowS(validSig)
	assert.Equal(t, ret, false)

	InitSecp256()
	ret = CheckLowS(validSig)
	assert.Equal(t, ret, false)

	sig1 := []byte{0x30, 0x3a, 0xd1, 0x5c, 0x20, 0x92, 0x75, 0xca}
	ret = CheckLowS(sig1)
	assert.Equal(t, ret, false)

	sig2 := []byte{
		0x30,
		0x45,
		0x02,
		0x21,
		0x00, 0xe4, 0xc8, 0x2b, 0x4e, 0xd6, 0xc6, 0x25, 0xa5, 0x4c, 0xd9,
		0x35, 0xdd, 0xc0, 0xf5, 0x99, 0x3a, 0x09, 0x78, 0x00, 0x30, 0x08,
		0x07, 0xae, 0x7f, 0xe3, 0xe2, 0x0c, 0x5a, 0xbc, 0xe9, 0xbf, 0xcc,
		0x02,
		0x20,
		0x2d, 0xc3, 0x50, 0x9b, 0x5a, 0x26, 0xa7, 0xc7, 0xb8, 0xcc, 0x55,
		0xd7, 0x57, 0xb1, 0xe0, 0x67, 0xec, 0xd3, 0xcb, 0x1f, 0x63, 0xe5,
		0xb2, 0x8d, 0x33, 0x71, 0xe2, 0xa6, 0x7b, 0xf5, 0xbe, 0x29,
	}

	ret = CheckLowS(sig2)
	assert.Equal(t, true, ret)
}
