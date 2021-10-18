package main

import (
	"github.com/dadosjusbr/proto/coleta"
)

func Metadados() (coleta.Metadados) {
	return coleta.Metadados {
	    NaoRequerLogin: true,
	    NaoRequerCaptcha: true,
	     Acesso: coleta.Metadados_AMIGAVEL_PARA_RASPAGEM,
	     Extensao: coleta.Metadados_ODS,
	     EstritamenteTabular:  false,
	    FormatoConsistente: true,
	    TemMatricula: true,
	    TemLotacao:  true,
	    TemCargo:  true,
	    ReceitaBase: coleta.Metadados_DETALHADO,
	    Despesas: coleta.Metadados_DETALHADO,
	    OutrasReceitas: coleta.Metadados_DETALHADO,
	}
}