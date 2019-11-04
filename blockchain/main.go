package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

var lastBlock = ""
var blacklist = make(map[string]*Block, 100)
var userlist = make(map[string]*User, 10)
var userMoney = make(map[string]int64, 10)
var powerLeft = 16
var powerValue big.Int

func init() {

	powerValue := big.NewInt(1)
	powerValue.Lsh(powerValue, uint(256-powerLeft))

}

func showAllBlock() {
	if lastBlock == "" {
		return
	} else {
		p1 := blacklist[lastBlock]
		if p1 != nil {
			fmt.Println(hex.EncodeToString(p1.Hash))
			for ; p1.PreHash != nil; p1 = blacklist[string(p1.PreHash)] {
				fmt.Println(string(p1.Hash))
			}
		}
	}

}

func showCount(c []Contest) int {
	if lastBlock == "" {
		return 0
	} else {
		p1 := blacklist[lastBlock]
		fmt.Println(hex.EncodeToString(p1.Hash))
		for ; p1.PreHash != nil; p1 = blacklist[string(p1.PreHash)] {

			for i := 0; i < len(p1.Tx); i++ {
				for _, v := range p1.Tx[i].Output {
					userMoney[string(v.RAddr)] += v.Value
				}
			}
		}
	}
	for k, v := range userMoney {
		fmt.Println(k, v)
	}
	return 0

}

type User struct {
	Name    string
	Address string
}

func createUsers() {

	userlist["A"] = &User{"A", "address1"}
	userlist["B"] = &User{"B", "address2"}

}

func getUser(name string) string {

	if userlist[name] != nil {
		return userlist[name].Address
	}
	return ""
}

func hashBijiao(v1 []byte) bool {
	var b big.Int
	c := b.SetBytes(v1)
	if c.Cmp(&powerValue) == -1 {
		return true
	} else {
		return false
	}
}

func aCountToB() {

}

func main() {
	createUsers()
	fmt.Println("ok")
	block := createFirstBlock()
	fmt.Println("ok1111")
	blacklist[hex.EncodeToString(block.Hash)] = block

	fmt.Println("szblockchain first show block")
	showAllBlock()
	fmt.Println("\n")
	fmt.Println("szblockchain aCountToB\n")
	aCountToB()
	fmt.Println("\n")
	fmt.Println("szblockchain zhuanzhang show block\n")
	showAllBlock()
	fmt.Println("\n")
	fmt.Println("szblockchain end")
}
