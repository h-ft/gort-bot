package discord

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"main/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"main.go/config"
)

func Init(cfg *config.Config) error {
	token := cfg.DiscordToken

	logrus.Infoln("Initializing bot ...")
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		logrus.Errorf("[discord.Init] Error initializing bot: %v", err)
		return err
	}

	if err = bot.Open(); err != nil {
		logrus.Errorf("[discord.bot.Open] Error opening connection: %v", err)
		return err
	}

	bot.AddHandler(messageCreate)

	// Set bot's status to active
	bot.UpdateGameStatus(0, "")

	logrus.Infoln("Bot is now running, press ctrl + c to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()

	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	logrus.Infoln("[" + m.Author.Username + "] " + m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "?") {
		command := strings.Split(m.Content[1:len(m.Content)], " ")
		name := strings.ToLower(command[0])
		commands.RunCommand(name, s, m)
		return
	}
}
