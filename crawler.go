package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dadosjusbr/status"
)

const (
	baseURL          = "https://transparencia.mppb.mp.br/PTMP/"
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
	links[tipoIndenizacoes] = fmt.Sprintf("%sFolhaVerbaIndenizRemTemporariaOds2022?&exe=%d&mes=%d&html=true/javascript:downloadArquivo();", baseURL, year, month)
	for t, id := range tipos {
		links[t] = fmt.Sprintf("%sFolhaPagamentoExercicioMesNewOds2022?&exe=%d&mes=%d&tipo=%d&html=true/javascript:downloadArquivo();", baseURL, year, month, id)
	}
	return links
}

func download(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return status.NewError(status.ConnectionError, fmt.Errorf("error downloading file:%q", err))
	}
	if resp.StatusCode != 200 {
		return status.NewError(status.DataUnavailable, fmt.Errorf("Sem dados"))
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return status.NewError(status.SystemError, fmt.Errorf("error copying response content:%q", err))
	}

	return nil
}
