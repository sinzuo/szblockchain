package main

type Inputtx struct {
	Ctid     []byte
	Index    int
	Signname []byte
	Laddr    []byte
}

type Outputtx struct {
	RAddr []byte
	Value int64
}

type Contest struct {
	Ctid   []byte
	Input  []Inputtx
	Output []Outputtx
}
