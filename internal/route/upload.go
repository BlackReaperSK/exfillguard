package route

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

// Token do seu bot do Discord
const discordBotToken = "--"

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
			return
		}

		// Gerar um ID único para o arquivo
		fileID := fmt.Sprintf("%d", time.Now().UnixNano())

		// Criar um canal no Discord com o ID do arquivo
		channelID, err := createDiscordChannel(fileID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Discord channel"})
			return
		}

		err = sendFileToDiscord(file, fileID, channelID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send file to Discord"})
			return
		}

		// Responder com o ID do arquivo
		c.JSON(http.StatusOK, gin.H{"fileID": fileID})
	}
}

func createDiscordChannel(fileID string) (string, error) {
	dg, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		return "", err
	}
	defer dg.Close()
	// ID DO SERVIDOR
	guildID := "--"

	channel, err := dg.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name: fmt.Sprintf("file-upload-%s", fileID),
		Type: discordgo.ChannelTypeGuildText,
	})
	if err != nil {
		return "", err
	}

	return channel.ID, nil
}

func sendFileToDiscord(file io.Reader, fileID, channelID string) error {
	dg, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		return err
	}
	defer dg.Close()

	// Ler o conteúdo do arquivo em um buffer de bytes
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return err
	}

	_, err = dg.ChannelFileSend(channelID, fmt.Sprintf("%s-file", fileID), buffer)
	if err != nil {
		return err
	}

	return nil
}
