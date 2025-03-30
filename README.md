# Chief

## Visão Geral

Este projeto é uma ferramenta desenvolvida em Golang para gerenciar migrations de banco de dados. Ele permite criar, aplicar e reverter migrations de forma simples e eficiente, garantindo consistência nos ambientes de desenvolvimento e produção.

## Roadmap (Prioridades)

- [x] Criação de migrations
- [x] Execução de migrations
- [ ] Rollback de migrations

## Uso

### Criar uma nova migration

```sh
chief -create --name <nome_da_migration>
```

### Aplicar migrations pendentes

```sh
chief -migrate
```

### Reverter a última migration

```sh
chief -rollback
```

### Listar migrations aplicadas

```sh
chief -history
```
