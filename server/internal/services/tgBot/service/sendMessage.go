package service

import (
	"context"
	"strings"

	"pkg/errors"
	"pkg/log"

	"server/internal/services/tgBot/model"
)

var Replacer = strings.NewReplacer(
	"_", "\\_",
	"*", "\\*",
	"[", "\\[",
	"]", "\\]",
	"(", "\\(",
	")", "\\)",
	"~", "\\~",
	"`", "\\`",
	">", "\\>",
	"#", "\\#",
	"+", "\\+",
	"-", "\\-",
	"=", "\\=",
	"|", "\\|",
	"{", "\\{",
	"}", "\\}",
	".", "\\.",
	"!", "\\!",
)

// SendMessage отправляет сообщение пользователю в телеграм
func (s *TgBotService) SendMessage(ctx context.Context, req model.SendMessageReq) error {

	if !s.isOn {
		log.Warning(ctx, "Вызвана функция SendMessage. Пуши выключены")
		return nil
	}

	if _, err := s.Bot.Send(s.Chat, req.Message); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
