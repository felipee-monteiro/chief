package routes 

import (
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
    	"strings"

	"api-test/types"
	"api-test/db"
    	"api-test/utils"
)

func InitialPage(db_conn *sql.DB) {
        
	http.HandleFunc("/api/initial-page", func(w http.ResponseWriter, r *http.Request) {
		
		var (
			total_processos string
			total_analise   string
			total_atrasados string
			total_10dias    string
			total_5dias     string
			total_1dia      string
		)

		query := r.URL.Query()
		user_id := query.Get("usuario_id")

		if len(strings.TrimSpace(user_id)) != 4 {
			w.WriteHeader(400)

			error_validation := types.ValidationErrorResponse{"ID de Usuario invalido"}

			message, err := json.Marshal(error_validation)

			if err != nil {
				log.Fatal("Erro ao escrever erro de validacao : ", err)
			}

			_, e := w.Write(messagbindingf e != nil {
				log.Fatal("Erro ao escrever resposta de validacao da API")
			}

                        return
		}

        if !utils.IsNumeric(user_id) {
            w.WriteHeader(400)

            message := types.ValidationErrorResponse{"ID do usuario precisa ser numerico"}

            value, err := json.Marshal(message)
            
            if err != nil {
                log.Fatal("Erro ao escrever resposta de validacao")
                return
            }

            _, e := w.Write(value)

            if e != nil {
                log.Fatal("Erro ao enviar resposta ao cliente")   
            }

            return
        }

		rows := db.Query(`
			 SELECT
			COUNT(*) AS total_processos,
			ISNULL(
				SUM(
                		CASE
						WHEN (pc.dias_na_pauta - pc.dias_em_pausa) <= 20 THEN 1
						ELSE 0
						END
						),
						0
						) AS total_em_analise,
						ISNULL(
							SUM(
						CASE
						WHEN (pc.dias_na_pauta - pc.dias_em_pausa) > 20 THEN 1
						ELSE 0
						END
						),
						0
						) AS total_atrasados,
						ISNULL(
					SUM(
						CASE
						WHEN ((20 - (pc.dias_na_pauta - pc.dias_em_pausa)) <= 10) AND ((20 - (pc.dias_na_pauta - pc.dias_em_pausa)) > 5) THEN 1
						ELSE 0
						END
						),
						0
						) AS total_em_pauta_10,
						ISNULL(
							SUM(
								CASE
							WHEN ((20 - (pc.dias_na_pauta - pc.dias_em_pausa)) <= 5) AND ((20 - (pc.dias_na_pauta - pc.dias_em_pausa)) > 1) THEN 1
							ELSE 0
						END
						),
						0
						) AS total_em_pauta_5,
						ISNULL(
							SUM(
								CASE
								WHEN ((20 - (pc.dias_na_pauta - pc.dias_em_pausa)) <= 1) AND ((20 - (pc.dias_na_pauta - pc.dias_em_pausa)) > 0) THEN 1
							ELSE 0
						END
						),
						0
						) AS total_em_pauta_1
						FROM
				VIEW_PROCESSOS_CALCULADOS pc
				WHERE pc.usuario_id = '` + user_id +`';
		`, db_conn)

		for rows.Next() {
			if err := rows.Scan(
				&total_processos, 
				&total_analise, 
				&total_atrasados, 
				&total_10dias, 
				&total_5dias, 
				&total_1dia,
			); err != nil {
				log.Fatal(err)
			}
		}
	
		response := types.InitialPageResponse{total_processos, total_analise, total_atrasados, total_10dias, total_5dias, total_1dia}
		
		result, err := json.Marshal(response)
	
		if err != nil {
			log.Fatal("Algo de errado não está certo: ", err)
		}
		
		_, e := w.Write(result)
		
		if e != nil {
			log.Fatal("Algo de errado não está certo ao escrever a resposta: ",)
		} 
	})
}
