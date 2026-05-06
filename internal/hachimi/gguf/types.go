package gguf

type Model struct {
	DisplayName string
	Name        string
	Description string
	Groups      []Group
}

type Group struct {
	Parameters string // this parameters
	Quants     []Quant
}

type Quant struct {
	DisplayName string
	Size        string
	Mod         string // quantization mode
	URL         string // hugging face url
	Sha256      string // sha256 verify
}
