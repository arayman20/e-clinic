package domain

import "fmt"

type PesertaBPJSDomain struct {
	PesertaBPJSRepo PesertaBPJSInterface
}

func (dom *PesertaBPJSDomain) RegisterPesertaBPJS(data PesertaBPJS) (code int, message string) {
	dataPeserta, err := dom.PesertaBPJSRepo.GetPesertaByNoBPJS(data.NoBPJS)
	if dataPeserta != nil && err == nil {
		return 400, fmt.Sprintf("No BPJS %s sudah terdaftar", data.NoBPJS)
	}

	if err != nil {
		return 500, err.Error()
	}

	errInsert := dom.PesertaBPJSRepo.RegisterPesertaBPJS(data)
	if errInsert != nil {
		return 500, err.Error()
	}

	return 201, fmt.Sprintf("Success register peserta BPJS a/n %s dan No BPJS %s", data.Name, data.NoBPJS)
}

