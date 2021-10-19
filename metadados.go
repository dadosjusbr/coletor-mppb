package main

import (
	"github.com/dadosjusbr/proto/coleta"
)

// Estes metadados foram definidos manualmente, pois foi feita uma análise dos dados
// no período de Janeiro de 2018 até Agosto de 2021. A disponibilização dos dados se
// manteve estável durante todo o tempo. Caso ocorra qualquer mudança na maneira que
// o órgão expõe seus dados, alterações devem ser feitas nestes metadados.
func Metadados() (coleta.Metadados) {
	return coleta.Metadados {
	    NaoRequerLogin: true,
	    NaoRequerCaptcha: true,
	    Acesso: coleta.Metadados_AMIGAVEL_PARA_RASPAGEM,
	    Extensao: coleta.Metadados_ODS,
	    EstritamenteTabular:  false,	// Dados limpos, toda linha é uma variável e toda coluna é um valor.
	    FormatoConsistente: true,	// Forma que o órgão expõe os dados, quantidade e ordem das colunas.
	    TemMatricula: true,
	    TemLotacao:  true,
	    TemCargo:  true,
	    ReceitaBase: coleta.Metadados_DETALHADO,
	    Despesas: coleta.Metadados_DETALHADO,
	    OutrasReceitas: coleta.Metadados_DETALHADO,
	}
}