package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature intercept,schema/snapshot,sql/upsert,sql/modifier,entql ./schema --idtype int64
