package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// NOTE: All of the below fields are required for this example to work correctly.
	var (
		Token     = flag.String("t", "", "Token or email address of Discord account.")
		GuildID   = flag.String("g", "", "Guild ID")
		ChannelID = flag.String("c", "", "Channel ID")
		Folder    = flag.String("f", "", "Folder of files to play.")
		err       error
	)
	flag.Parse()

	// Connect to Discord
	discord, err := discordgo.New(*Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open Websocket
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Connect to voice channel.
	// NOTE: Setting mute to false, deaf to true.
	dgv, err := discord.ChannelVoiceJoin(*GuildID, *ChannelID, false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start loop and attempt to play all files in the given folder
	fmt.Println("Reading Folder: ", *Folder)
	files, _ := os.ReadDir(*Folder)
	for _, f := range files {
		fmt.Println("PlayAudioFile:", f.Name())
		discord.UpdateStatus(0, f.Name())

		dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", *Folder, f.Name()), make(chan bool))
	}

	// Close connections
	dgv.Close()
	discord.Close()

	return
}
