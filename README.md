# Chief

## Visão Geral

Este projeto é uma ferramenta interna desenvolvida em Golang para gerenciar migrations de banco de dados. Ele permite criar, aplicar e reverter migrations de forma simples e eficiente, garantindo consistência nos ambientes de desenvolvimento e produção.

## Roadmap (Prioridades)

- [x] Criação de migrations
- [z] Execução de migrations
- [ ] Rollback de migrations

## Uso

### Criar uma nova migration

```sh
go run main.go -create <nome_da_migration>
```

### Aplicar migrations pendentes

```sh
go run main.go -migrate
```

### Reverter a última migration

```sh
go run main.go -rollback
```

### Listar migrations aplicadas

```sh
go run main.go -history
```
