package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	baseURL          = "http://pitagoras.mppb.mp.br/PTMP/"
	tipoIndenizacoes = "indenizacoes"
)

var (
	tipos = map[string]int{
		"membrosAtivos": 1,
	}
)

// Crawl retrieves payment files from MPPB.
func Crawl(outputPath string, month, year int) ([]string, error) {
	var files []string
	for typ, url := range links(baseURL, month, year) {
		filePath := fmt.Sprintf("%s/%s-%d-%d.ods", outputPath, typ, month, year)
		f, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		if err := download(url, f); err != nil {
			return nil, err
		}
		files = append(files, filePath)
	}
	return files, nil
}

// Generate endpoints able to download
func links(baseURL string, month, year int) map[string]string {
	links := make(map[string]string)
	links[tipoIndenizacoes] = fmt.Sprintf("%sFolhaVerbaIndenizRemTemporariaOds?mes=%d&exercicio=%d&tipo=", baseURL, month, year)
	for t, id := range tipos {
		links[t] = fmt.Sprintf("%sFolhaPagamentoExercicioMesNewOds?mes=%d&exercicio=%d&tipo=%d", baseURL, month, year, id)
	}
	return links
}

func download(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading file:%q", err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("error copying response content:%q", err)
	}
	return nil
}
