package verify

import (
	"fmt"
	"net/http"
	"net/smtp"
	"net/textproto"
	"purpleschool/3-validation-api/configs"
	"purpleschool/3-validation-api/pkg/response"

	"github.com/jordan-wright/email"
)


type VerifyHandler struct{
	*configs.Config
}


type VerifyHandlerDeps struct {
	*configs.Config
}

type sendEmailArg struct {
	emailTo []string
	emailFrom string
	password string
	address string
}

func NewVerifyHandler(route *http.ServeMux, deps VerifyHandlerDeps){
	handler := &VerifyHandler{
		deps.Config,
	}
	route.HandleFunc("GET /verify/{hash}", handler.Verify())
	route.HandleFunc("POST /send", handler.Send())
}


func (v *VerifyHandler) Verify() http.HandlerFunc{
	return func (w http.ResponseWriter, req *http.Request){
		
	}
}

func sendEmail(params sendEmailArg) error{
	e := &email.Email {
		To: params.emailTo,
		From: fmt.Sprintf("Jordan Wright <%s>", params.emailFrom),
		Subject: "Для пуски милашки",
		Text: []byte("Тут какой то текст!"),
		HTML: []byte("<h1>Пускин атакушкин, зевашкин</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	if err := e.Send(fmt.Sprintf("%s:587", params.address), smtp.PlainAuth("", params.emailFrom, params.password, params.address)); err != nil {
		return err
	}
	return nil
}

func (v *VerifyHandler) Send() http.HandlerFunc{
	return func (w http.ResponseWriter, req *http.Request){
		sendParams := sendEmailArg{
			emailTo: []string{"kornilovsergey55@gmail.com"},
			emailFrom: v.ConnectEmail.Email,
			password: v.ConnectEmail.Password,
			address: v.ConnectEmail.Address,
		}
		err := sendEmail(sendParams)

		if err != nil {
			fmt.Println("Error send email: ", err.Error())
			res := SendResponse{
				Message: "falled",
			}
			response.Json(w, res, 500)
			return
		}
		res := SendResponse{
			Message: "success",
		}
		response.Json(w, res, 201)

	}
}

