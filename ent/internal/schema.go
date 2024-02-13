// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"team-manager/ent/schema","Package":"team-manager/ent","Schemas":[{"name":"Quote","config":{"Table":""},"fields":[{"name":"id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"quote","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"author","type":{"Type":19,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"2024-02-10T19:36:55.358554+06:00","default_kind":25,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"User","config":{"Table":""},"fields":[{"name":"id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"tg_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"surname","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"phone","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"2024-02-10T19:36:55.359701+06:00","default_kind":25,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}}]}],"Features":["intercept","schema/snapshot","sql/upsert","sql/modifier","entql"]}`