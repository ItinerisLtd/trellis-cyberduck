module github.com/itinerisltd/trellis-cyberduck

go 1.15

require (
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/mitchellh/cli v1.1.2
	github.com/spf13/cobra v1.1.1
	trellis-cli v0.0.0-00010101000000-000000000000
)

replace trellis-cli => github.com/roots/trellis-cli v0.9.3-0.20201208100110-a9b128aa89e9
