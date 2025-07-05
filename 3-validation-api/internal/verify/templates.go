package verify

import (
	"bytes"
	"html/template"
)

type EmailData struct {
	VerificationURL string
	Email           string
}

type PageData struct {
	Email   string
	Message string
}

func renderEmailTemplate(data EmailData) (string, error) {
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func renderSuccessPage(email string) (string, error) {
	tmpl, err := template.New("success").Parse(successTemplate)
	if err != nil {
		return "", err
	}
	
	var buf bytes.Buffer
	data := PageData{Email: email}
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	
	return buf.String(), nil
}

func renderErrorPage(message string) (string, error) {
	tmpl, err := template.New("error").Parse(errorTemplate)
	if err != nil {
		return "", err
	}
	
	var buf bytes.Buffer
	data := PageData{Message: message}
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	
	return buf.String(), nil
}

const emailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Подтверждение email</title>
</head>
<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background: #f8f9fa; padding: 30px; border-radius: 10px;">
        <h1 style="color: #333; text-align: center;">Подтвердите ваш email</h1>
        <p style="font-size: 16px; line-height: 1.5;">Здравствуйте!</p>
        <p style="font-size: 16px; line-height: 1.5;">
            Для завершения регистрации нажмите на кнопку ниже для подтверждения вашего email адреса:
        </p>
        <div style="text-align: center; margin: 30px 0;">
            <a href="{{.VerificationURL}}" 
               style="background: #007bff; color: white; padding: 15px 30px; 
                      text-decoration: none; border-radius: 5px; display: inline-block;
                      font-weight: bold; font-size: 16px;">
                ✉️ Подтвердить Email
            </a>
        </div>
        <p style="color: #666; font-size: 14px; text-align: center;">
            Если вы не регистрировались на нашем сайте, проигнорируйте это письмо.
        </p>
        <hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">
        <p style="color: #999; font-size: 12px; text-align: center;">
            Если кнопка не работает, скопируйте эту ссылку в браузер:<br>
            <a href="{{.VerificationURL}}">{{.VerificationURL}}</a>
        </p>
    </div>
</body>
</html>
`

const successTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Email подтвержден</title>
</head>
<body style="font-family: Arial, sans-serif; text-align: center; padding: 50px;">
    <div style="max-width: 500px; margin: 0 auto; background: #f8f9fa; padding: 40px; border-radius: 10px;">
        <h1 style="color: #28a745; margin-bottom: 20px;">✅ Email успешно подтвержден!</h1>
        <p style="font-size: 18px; margin-bottom: 10px;">
            Ваш email адрес <strong>{{.Email}}</strong> был успешно подтвержден.
        </p>
        <p style="color: #666; font-size: 16px;">
            Теперь вы можете пользоваться всеми функциями нашего сервиса.
        </p>
        <div style="margin-top: 30px;">
            <a href="/" style="background: #007bff; color: white; padding: 10px 20px; 
                             text-decoration: none; border-radius: 5px;">
                Вернуться на главную
            </a>
        </div>
    </div>
</body>
</html>
`

const errorTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Ошибка подтверждения</title>
</head>
<body style="font-family: Arial, sans-serif; text-align: center; padding: 50px;">
    <div style="max-width: 500px; margin: 0 auto; background: #f8f9fa; padding: 40px; border-radius: 10px;">
        <h1 style="color: #dc3545; margin-bottom: 20px;">❌ Ошибка подтверждения</h1>
        <p style="font-size: 18px; margin-bottom: 10px;">{{.Message}}</p>
        <p style="color: #666; font-size: 16px;">
            Возможные причины:
        </p>
        <ul style="text-align: left; color: #666;">
            <li>Email уже был подтвержден</li>
            <li>Неверная ссылка</li>
        </ul>
        <div style="margin-top: 30px;">
            <a href="/send" style="background: #007bff; color: white; padding: 10px 20px; 
                               text-decoration: none; border-radius: 5px;">
                Отправить новое письмо
            </a>
        </div>
    </div>
</body>
</html>
`
