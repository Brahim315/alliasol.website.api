package handler

import (
	"AlliasolDevis/utils"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"net/http"
)

func DevisHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	info := new(utils.UserInfo)

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	/*from := "from@example.com"
	password := "e3a1e1689e15db"
	to := []string{"gnaoui-brahim@hotmail.fr"}
	smtpHost := "smtp.mailtrap.io"
	smtpPort := "587"

	message := info.Firstname + " " + info.Lastname + " - " + info.City + " - " +
		info.Email + " - " + info.PhoneNumber + " - " + info.Project

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}*/
	err = SendMail(info.Firstname, info.Lastname, info.City, info.Email, info.PhoneNumber, info.Project)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent succesfully")
	w.WriteHeader(204)
}

func SendMail(name, lastname, city, email, phoneNumber, project string) error {

	subject := "Demande de devis"
	dialer := gomail.NewDialer("smtp.mailtrap.io", 587, "7ca6d9ab0815a2", "e3a1e1689e15db")

	message := gomail.NewMessage()
	message.SetAddressHeader("From", "no-reply@alliasol.com", "Alliasol")
	message.SetAddressHeader("To", "devis@alliasol.com", "Devis")
	message.SetHeader("Subject", subject)

	message.SetBody("text/plain", name+" "+lastname+" "+city+" "+email+" "+phoneNumber+" "+project)

	return dialer.DialAndSend(message)
}
