module github.com/dayvillefire/dgvoice

go 1.24.2

toolchain go1.24.3

replace github.com/dayvillefire/dgvoice/mpg123 => ./mpg123

require (
	github.com/Carmen-Shannon/gopus v1.0.0
	github.com/bwmarrin/dgvoice v0.0.0-20210225172318-caaac756e02e
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dayvillefire/dgvoice/mpg123 v0.0.0-20251213183936-608202d5f9eb
	github.com/hajimehoshi/ebiten/v2 v2.9.7
)

require (
	github.com/ebitengine/oto/v3 v3.4.0 // indirect
	github.com/ebitengine/purego v0.9.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	golang.org/x/crypto v0.47.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	layeh.com/gopus v0.0.0-20210501142526-1ee02d434e32 // indirect
)
