package main

import (
	"crypto/sha256"
	//	"math"
	//"math/big"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	MineGet = 10
)

type Block struct {
	TimeS   int64
	Tx      []*Contest
	Hash    []byte
	Count   int64
	PreHash []byte
}

func base64Hash(val []byte) string {
	bb := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, bb)
	encoder.Write(val)
	encoder.Close()
	return bb.String()

}

func hashBase64(val []byte) []byte {
	bb := &bytes.Buffer{}
	decoder := base64.NewDecoder(base64.StdEncoding, bb)
	decoder.Read(val)

	return bb.Bytes()
}

func interfaceToBytes(a interface{}) []byte {
	var b bytes.Buffer
	binary.Write(&b, binary.BigEndian, a)

	return b.Bytes()
}

func createFirstBlock() *Block {
	a := sha256.Sum256([]byte("create first world"))

	lastBlock = hex.EncodeToString(a[:])
	fmt.Println("lastBlock is " + lastBlock)
	b := Block{TimeS: time.Now().Unix(), Hash: a[:]}
	b.Tx = []*Contest{createMineContest("ok")}

	b.PreHash = nil

	return &b
}

func createBlock() {
	a := sha256.Sum256([]byte("create first world"))
	b1 := &Block{TimeS: time.Now().Unix(), Hash: a[:]}
	b1.Tx[0] = createMineContest(getUser("a"))
	b1.PreHash = []byte(lastBlock)
	lastBlock = string(base64Hash(b1.Hash))
}

func createMineContest(address string) *Contest {
	var tx Contest
	a := sha256.Sum256([]byte("jiang"))
	tx.Ctid = a[:]
	tx.Input = []Inputtx{Inputtx{Ctid: nil, Index: -1, Laddr: []byte(address), Signname: []byte("")}}
	tx.Output = []Outputtx{Outputtx{Value: MineGet, RAddr: []byte(address)}}
	return &tx

}

func aCountToB() {

}

func createZhuanzhangContest(address1 string, address2 string, val int64) {
	/*var tx Contest
	tx.Ctid = sha256.New("jiang")
	tx.Input = &Inputtx{Ctid: "", Index: 0, Laddr: address, Signname: ""}
	tx.Output = &Outputtx{Value: val, RAddr: address2}
	*/
}

func hashTree(c []*Contest) []byte {
	var a []byte
	for i := 0; i < len(c); i++ {
		b := sha256.Sum256(interfaceToBytes(c[i]))
		a = b[:]
	}
	return a[:]
}

func (b *Block) verityBool() bool {

	bytes.Join([][]byte{HexToString(b.TimeS), hashTree(b.Tx), HexToString(b.Count)}, []byte{})
	return true
}
