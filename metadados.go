package main

import (
	"github.com/dadosjusbr/proto/coleta"
)

func Metadados() (coleta.Metadados) {
	var metadados coleta.Metadados

	metadados.NaoRequerLogin = true
	metadados.NaoRequerCaptcha = true
	metadados.Acesso = coleta.Metadados_FormaDeAcesso(1)
	metadados.Extensao = coleta.Metadados_Extensao(1)
	metadados.EstritamenteTabular = false
	metadados.FormatoConsistente = true
	metadados.TemMatricula = true
	metadados.TemLotacao = true
	metadados.TemCargo = true
	metadados.ReceitaBase = coleta.Metadados_OpcoesDetalhamento(2)
	metadados.Despesas = coleta.Metadados_OpcoesDetalhamento(2)
	metadados.OutrasReceitas = coleta.Metadados_OpcoesDetalhamento(2)

	return metadados
}