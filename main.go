package main

import (
	"fmt"
	"os"
	"strconv"

	dadojusbr "github.com/dadosjusbr/proto"
	"github.com/dadosjusbr/proto/coleta"
	"github.com/dadosjusbr/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	agenciaID  = "mppb"
	repColetor = "https://github.com/dadosjusbr/coletor-mppb"
)

func main() {

	month, err := strconv.Atoi(os.Getenv("MONTH"))
	if err != nil {
		status.ExitFromError(status.NewError(status.InvalidInput, fmt.Errorf("Invalid month (\"%s\"): %w", os.Getenv("MONTH"), err)))
	}
	year, err := strconv.Atoi(os.Getenv("YEAR"))
	if err != nil {
		status.ExitFromError(status.NewError(status.InvalidInput, fmt.Errorf("Invalid year (\"%s\"): %w", os.Getenv("YEAR"), err)))
	}

	outputFolder := os.Getenv("OUTPUT_FOLDER")
	if outputFolder == "" {
		outputFolder = "./output"
	}

	if err := os.Mkdir(outputFolder, os.ModePerm); err != nil && !os.IsExist(err) {
		status.ExitFromError(status.NewError(status.SystemError, fmt.Errorf("Error creating output folder(%s): %w", outputFolder, err)))
	}

	crawlerVersion := os.Getenv("CRAWLER_VERSION")
	if crawlerVersion == "" {
		crawlerVersion = "unspecified"
	}

	files, err := Crawl(outputFolder, month, year)
	if err != nil {
		status.ExitFromError(err)
	}

	chaveColeta := dadojusbr.IDColeta(agenciaID, month, year)

	folha, parseErr := Parse(files, chaveColeta)
	if parseErr != nil {
		status.ExitFromError(parseErr)
	}

	colRes := coleta.Coleta{
		ChaveColeta:        chaveColeta,
		Orgao:              agenciaID,
		Mes:                int32(month),
		Ano:                int32(year),
		TimestampColeta:    timestamppb.Now(),
		RepositorioColetor: repColetor,
		VersaoColetor:      crawlerVersion,
		DirColetor:         agenciaID,
		Arquivos:           files,
	}

	metadados := Metadados(int32(year), int32(month))

	rc := coleta.ResultadoColeta{
		Coleta:    &colRes,
		Folha:     folha,
		Metadados: &metadados,
	}

	b, err := prototext.Marshal(&rc)
	if err != nil {
		status.ExitFromError(status.NewError(status.OutputError, fmt.Errorf("JSON marshaling error: %w", err)))
	}
	fmt.Printf("%s", b)
}
