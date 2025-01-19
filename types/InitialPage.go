package types

type InitialPageResponse struct {
    TotalProcessos	 string  `json:"total_processos"`
    TotalEmAnalise   string  `json:"total_em_analise"`
    TotalAtrasados   string  `json:"total_atrasados"`
    TotalEmPauta10   string  `json:"total_10_dias"`
    TotalEmPauta5    string  `json:"total_5_dias"`
    TotalEmPauta1    string  `json:"total_1_dia"`  
}
