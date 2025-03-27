package Config

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func HandleVoiceMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	audioPath, err := downloadTelegramAudio(bot, message.Voice.FileID)
	if err != nil {
		log.Println("Error descargando audio:", err)
		return
	}
	defer os.Remove(audioPath)

	transcription, err := transcribeAudio(audioPath)
	if err != nil {
		log.Println("Error transcribiendo audio:", err)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, transcription)

	_, err = bot.Send(msg)
	if err != nil {
		return
	}
}

func convertToMP3(inputPath, outputPath string) error {
	err := ffmpeg.Input(inputPath).
		Output(outputPath, ffmpeg.KwArgs{"acodec": "libmp3lame"}).
		OverWriteOutput().
		Run()
	return err
}

func downloadTelegramAudio(bot *tgbotapi.BotAPI, fileID string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return "", err
	}

	url := "https://api.telegram.org/file/bot" + bot.Token + "/" + file.FilePath
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	audioOga := "audio.oga"
	out, err := os.Create(audioOga)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	audioMP3 := "audio.mp3"
	err = convertToMP3(audioOga, audioMP3)
	if err != nil {
		return "", err
	}

	os.Remove(audioOga)

	return audioMP3, nil
}

func transcribeAudio(audioFile string) (string, error) {
	file, err := os.Open(audioFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", audioFile)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	_ = writer.WriteField("model", "gpt-4o-transcribe")
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("apiKeyOpenAi"))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("🛠️ Respuesta de OpenAI:", string(responseBody)) // <--- Agregado para depuración

	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return "", err
	}

	if text, ok := result["text"].(string); ok {
		return text, nil
	}

	return "", fmt.Errorf("error en la transcripción: %s", string(responseBody)) // <-- Mejor mensaje de error
}
