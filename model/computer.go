package model

type Computer struct {
	SerialNumber int64  `json:"serialNumber"`
	Model        string `json:"model"`
	Ram          string `json:"ram"`
	Processor    string `json:"processor"`
}

type pcs []Computer

var Computers = pcs{
	{
		SerialNumber: 15556,
		Model:        "Apple",
		Ram:          "8GB",
		Processor:    "i7",
	},

	{
		SerialNumber: 559845,
		Model:        "Lenovo",
		Ram:          "4GB",
		Processor:    "i5",
	},
}

func GetAllComputers() []Computer {
	return Computers
}
