<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/mail-sendgrid</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/mail-sendgrid"><img src="https://pkg.go.dev/badge/github.com/togo-framework/mail-sendgrid.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/mail-sendgrid
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
