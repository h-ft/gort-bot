package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type (
	Commands struct {
		CommandList []Command
	}

	Command struct {
		Name    string
		Command func(s *discordgo.Session, m *discordgo.MessageCreate)
	}
)

var Cmds Commands

func Run(name string, s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, command := range Cmds.CommandList {
		if strings.EqualFold(command.Name, name) {
			command.Command(s, m)
			return
		}
	}
}
