module logger-implementations

go 1.19

require (
	github.com/rs/zerolog v1.29.0
	github.com/sirupsen/logrus v1.9.0
	github.com/tenenwurcel/logger v0.0.0-20230202144500-0f9c2f2e9f77
)

require (
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace (
	github.com/tenenwurcel/logger => ../logger
)
