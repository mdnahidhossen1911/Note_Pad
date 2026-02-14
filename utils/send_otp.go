package utils

import (
	"crypto/rand"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

func SendOTPToEmail( otp string, toEmail string, UserName string, appPass string, sendermail string) (string, error) {
	toEmail = strings.TrimSpace(toEmail)
	if toEmail == "" {
		return "", fmt.Errorf("email is required")
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := sendermail
	smtpPass := appPass

	from := smtpUser
	subject := "Your OTP Code"
	expiresIn := 2 * time.Minute

	body := buildHTMLBody(otp, UserName, int(expiresIn.Minutes()))

	msg := buildEmail(from, toEmail, subject, body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	addr := smtpHost + ":" + smtpPort

	if err := smtp.SendMail(addr, auth, from, []string{toEmail}, []byte(msg)); err != nil {
		return "", fmt.Errorf("send mail failed: %w", err)
	}

	return otp, nil
}

// ------- helpers (private) -------

func buildHTMLBody(otp string, userName string, expiresInMinutes int) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Your OTP Code</title>
</head>
<body style="margin:0;padding:0;background:#f4f4f4;font-family:Arial,sans-serif;">

  <table width="100%%" cellpadding="0" cellspacing="0" style="background:#f4f4f4;padding:40px 0;">
    <tr>
      <td align="center">
        <table width="480" cellpadding="0" cellspacing="0"
          style="background:#ffffff;border-radius:6px;overflow:hidden;box-shadow:0 2px 8px rgba(0,0,0,0.08);">

          <!-- Header -->
          <tr>
            <td style="background:#66ccff;padding:22px 32px;text-align:center;">
              <h1 style="margin:0;color:#ffffff;font-size:20px;font-weight:700;letter-spacing:0.3px;">
                Your OTP Code
              </h1>
            </td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:32px 36px 24px;">

              <p style="margin:0 0 16px;color:#333333;font-size:14px;">Hello %s,</p>

              <p style="margin:0 0 24px;color:#333333;font-size:14px;line-height:1.6;">
                Your One-Time Password (OTP) for account verification is:
              </p>

              <!-- OTP Box -->
              <table width="100%%" cellpadding="0" cellspacing="0" style="margin-bottom:24px;">
                <tr>
                  <td align="center"
                    style="border:1px solid #dddddd;border-radius:4px;padding:18px 0;background:#fafafa;">
                    <span style="font-size:34px;font-weight:700;color:#66ccff;letter-spacing:6px;font-family:'Courier New',monospace;">
                      %s
                    </span>
                  </td>
                </tr>
              </table>

              <p style="margin:0 0 10px;color:#333333;font-size:14px;line-height:1.6;">
                This OTP is valid for <strong>%d minutes</strong>. Please do not share this code with anyone.
              </p>

              <p style="margin:0 0 10px;color:#333333;font-size:14px;line-height:1.6;">
                If you didn't request this code, please ignore this email.
              </p>

              <p style="margin:0;color:#333333;font-size:14px;line-height:1.6;">
                Thank you for using our service!
              </p>

            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="border-top:1px solid #eeeeee;padding:16px 36px;text-align:center;">
              <p style="margin:0;color:#aaaaaa;font-size:12px;">
                © %d Your Company Name. All rights reserved.
              </p>
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>

</body>
</html>`, userName, otp, expiresInMinutes, time.Now().Year())
}

func buildEmail(from, to, subject, body string) string {
	var sb strings.Builder
	sb.WriteString("From: " + from + "\r\n")
	sb.WriteString("To: " + to + "\r\n")
	sb.WriteString("Subject: " + subject + "\r\n")
	sb.WriteString("MIME-Version: 1.0\r\n")
	sb.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	sb.WriteString("\r\n")
	sb.WriteString(body)
	return sb.String()
}

func GenerateOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid otp length")
	}
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	digits := make([]byte, length)
	for i := 0; i < length; i++ {
		digits[i] = '0' + (b[i] % 10)
	}
	return string(digits), nil
}
