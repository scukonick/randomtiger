module github.com/scukonick/randomtiger

go 1.14

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/jackc/pgx/v4 v4.9.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/scukonick/giphy v0.0.1
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)

replace github.com/peterhellberg/giphy => ../giphy
