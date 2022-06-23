package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
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

func RegisterCommands() {

	CmdCome := Command{
		Name:    "Come",
		Command: Come,
	}
	Cmds.CommandList = append(Cmds.CommandList, CmdCome)
}

func Come(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Find the channel that the message came from.
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}

	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			_, err := s.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)
			if err != nil {
				logrus.Infoln("Error joining channel:", err)
			}

			return
		}
	}

}
