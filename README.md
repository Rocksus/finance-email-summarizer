# Fundtract
Self hosted solution on aggregating your online transactions through various sources.



# Code styles

## Database Migration

Use `go-migrate` for creating new migrations. Here is an example command

```bash
migrate create -ext sql -dir db/migrations -seq <operation>_<table_name>_<detail:optional>
```

- Operation describes what the migration file will do, `create`, `drop`, or `update`. Use `insert`, or `delete` if the migration script is for adding/removing data (e.g. adding default secret values)
- Table name describes the name of the table
- Detail is for additional details