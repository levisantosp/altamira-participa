module github.com/levisantosp/altamira-participa

go 1.27.0

ignore ./pg

require (
	entgo.io/ent v0.14.6
	github.com/Oudwins/zog v0.22.2
	github.com/danielgtaylor/huma/v2 v2.39.1
	github.com/go-chi/chi/v5 v5.3.2
	github.com/go-chi/cors v1.2.2
	github.com/jackc/pgx/v5 v5.10.0
	github.com/joho/godotenv v1.5.1
	github.com/redis/go-redis/v9 v9.22.0
)

require (
	ariga.io/atlas v0.36.2-0.20250730182955-2c6300d0a3e1 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/kingpin/v2 v2.4.0 // indirect
	github.com/alecthomas/units v0.0.0-20240927000941-0f3dac36c52b // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dave/dst v0.27.3 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golangci/golines v0.15.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.18.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/ldez/structtags v0.6.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/rogpeppe/go-internal v1.15.0 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/mod v0.40.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/telemetry v0.0.0-20260811182544-a038080d80e5 // indirect
	golang.org/x/term v0.45.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	golang.org/x/tools v0.49.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	mvdan.cc/gofumpt v0.11.0 // indirect
)

tool (
	github.com/golangci/golines
	golang.org/x/tools/cmd/goimports
	mvdan.cc/gofumpt
)
