package main

import (
	"github.com/dadosjusbr/proto/coleta"
)

// Estes metadados foram definidos manualmente, pois foi feita uma análise dos dados
// no período de Janeiro de 2018 até Agosto de 2021. A disponibilização dos dados se
// manteve estável durante todo o tempo. Caso ocorra qualquer mudança na maneira que
// o órgão expõe seus dados, alterações devem ser feitas nestes metadados.
func Metadados(year int32, month int32) coleta.Metadados {
	var metadado coleta.Metadados
	metadado.Acesso = coleta.Metadados_ACESSO_DIRETO
	metadado.Extensao = coleta.Metadados_ODS
	metadado.EstritamenteTabular = false // Dados limpos, toda linha é uma variável e toda coluna é um valor.
	if year == 2022 && month == 6 {
		metadado.FormatoConsistente = false // Forma que o órgão expõe os dados, quantidade e ordem das colunas.
	} else {
		metadado.FormatoConsistente = true
	}
	metadado.TemMatricula = true
	metadado.TemLotacao = true
	metadado.TemCargo = true
	metadado.ReceitaBase = coleta.Metadados_DETALHADO
	metadado.Despesas = coleta.Metadados_DETALHADO
	metadado.OutrasReceitas = coleta.Metadados_DETALHADO
	return metadado
}
