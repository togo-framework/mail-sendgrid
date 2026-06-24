// Package sendgrid is a SendGrid (sendgrid.com) driver for togo mail. It shows the
// driver-plugin pattern: a plugin that depends on the mail plugin and registers a
// driver with it. Blank-import + set MAIL_DRIVER=sendgrid and SENDGRID_API_KEY.
package sendgrid

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/togo-framework/mail"
	"github.com/togo-framework/togo"
)

const endpoint = "https://api.sendgrid.com/v3/mail/send"

func init() {
	mail.RegisterDriver("sendgrid", func(k *togo.Kernel) (mail.Mailer, error) {
		key := os.Getenv("SENDGRID_API_KEY")
		if key == "" {
			return nil, errors.New("mail-sendgrid: SENDGRID_API_KEY not set")
		}
		return &mailer{key: key, client: &http.Client{Timeout: 15 * time.Second}}, nil
	})
}

type mailer struct {
	key    string
	client *http.Client
}

func emails(addrs []string) []map[string]string {
	out := make([]map[string]string, 0, len(addrs))
	for _, a := range addrs {
		out = append(out, map[string]string{"email": a})
	}
	return out
}

func (m *mailer) Send(ctx context.Context, msg mail.Message) error {
	personalization := map[string]any{"to": emails(msg.To)}
	if len(msg.Cc) > 0 {
		personalization["cc"] = emails(msg.Cc)
	}
	if len(msg.Bcc) > 0 {
		personalization["bcc"] = emails(msg.Bcc)
	}
	var content []map[string]string
	if msg.Text != "" {
		content = append(content, map[string]string{"type": "text/plain", "value": msg.Text})
	}
	if msg.HTML != "" {
		content = append(content, map[string]string{"type": "text/html", "value": msg.HTML})
	}
	body := map[string]any{
		"personalizations": []map[string]any{personalization},
		"from":             map[string]string{"email": msg.From},
		"subject":          msg.Subject,
		"content":          content,
	}
	buf, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.key)
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("mail-sendgrid: status %d: %s", resp.StatusCode, string(b))
	}
	return nil
}
