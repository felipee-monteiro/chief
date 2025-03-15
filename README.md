# Chief

## Visão Geral
Este projeto é uma ferramenta interna desenvolvida em Golang para gerenciar migrations de banco de dados. Ele permite criar, aplicar e reverter migrations de forma simples e eficiente, garantindo consistência nos ambientes de desenvolvimento e produção.

## Funcionalidades
- Criar novas migrations automaticamente
- Aplicar migrations pendentes ao banco de dados
- Reverter migrations aplicadas
- Listar o histórico de migrations executadas
- Suporte a diferentes bancos de dados

## Requisitos
- Go 1.24+
- Banco de dados compatível (SQLServer2016+)

## Instalação
1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```
2. Instale as dependências:
   ```sh
   go mod tidy
   ```
3. Configure as variáveis de ambiente para conexão com o banco de dados:
   ```sh
   export DB_DRIVER=postgres
   export DB_DSN="user=usuario password=senha dbname=banco sslmode=disable"
   ```

## Uso
### Criar uma nova migration
```sh
go run main.go create <nome_da_migration>
```

### Aplicar migrations pendentes
```sh
go run main.go migrate
```

### Reverter a última migration
```sh
go run main.go rollback
```

### Listar migrations aplicadas
```sh
go run main.go history
```

## Estrutura do Projeto
```
/seu-projeto
├── migrations/       # Arquivos de migrations
├── internal/         # Código interno da ferramenta
│   ├── database/     # Gerenciamento de conexão e execução de migrations
│   ├── cli/          # Interface de linha de comando
├── main.go           # Ponto de entrada do projeto
├── go.mod            # Gerenciamento de dependências
└── README.md         # Documentação do projeto
```

## Contato
Para mais informações, entre em contato com a equipe de desenvolvimento.

