# mail-sendgrid

> **SendGrid** driver for [togo](https://to-go.dev) **mail**.

Sends transactional email through the SendGrid v3 API.

## Install

```bash
togo install togo-framework/mail-sendgrid
```

Then set the driver + key:

```ini
MAIL_DRIVER=sendgrid
SENDGRID_API_KEY=SG.xxxxx
```

Blank-importing this plugin registers the `sendgrid` mail driver with the kernel; `togo serve` picks it up. Implements `mail.Mailer` (`Send(ctx, mail.Message)`).

MIT
