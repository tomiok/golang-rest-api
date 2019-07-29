package model

import "errors"

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

func SearchBySerial(serialNumber int64) (Computer, error) {
	for i, element := range Computers {
		if element.SerialNumber == serialNumber {
			return Computers[i], nil
		}
	}
	return Computer{}, errors.New("not in this list")
}

func UpdateRam(ram string, serial int64) *Computer {

	for i, comps := range Computers {
		if comps.SerialNumber == serial {

			computer := Computers[i]
			computer.Ram = ram
			p := append(Computers[:i], computer)
			return &p[i]
		}
	}

	return nil
}
