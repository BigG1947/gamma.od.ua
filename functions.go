package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func UploadImages(file multipart.File) (string, error) {
	tempFile, err := ioutil.TempFile("upload-images", "upload-*.png")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

func UploadPDF(file multipart.File) (string, error) {
	tempFile, err := ioutil.TempFile("upload-pdf", "upload-*.pdf")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

func DeleteImages(src string) error {
	err := os.Remove(src)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func DeletePDF(src string) error {
	err := os.Remove(src)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

//func checkReCaptchaV2(token string, secret string) bool {
//	response, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{
//		"response": {token}, // token from js
//		"secret":   {secret}}) //secret key for you site
//
//	//okay, moving on...
//	if err != nil {
//		log.Printf("%s\n", err)
//		return false
//	}
//
//	defer response.Body.Close()
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Printf("%s\n", err)
//		return false
//	}
//
//	res := struct {
//		Success   bool     `json:"success"`
//		Score     float64  `json:"score"`
//		ErrorCode []string `json:"error-codes"`
//	}{}
//
//	err = json.Unmarshal(body, &res)
//	if err != nil {
//		log.Printf("%s\n", err)
//		return false
//	}
//	return res.Success
//}

func checkReCaptchaV3(token string, secret string) bool {

	response, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{
		"response": {token},   // token from js
		"secret":   {secret}}) //secret key for you site

	if err != nil {
		log.Printf("%s\n", err)
		return false
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s\n", err)
		return false
	}

	//log.Printf("%s\n", string(body))

	res := struct {
		Success   bool     `json:"success"`
		Score     float64  `json:"score"`
		ErrorCode []string `json:"error-codes"`
	}{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Printf("%s\n", err)
		return false
	}
	return res.Success && res.Score >= MAX_VALID_SCORE
}

func getMessageForFeedBack(code int64) string {
	switch code {
	case 1:
		return "Ваше сообщения успешно отправленно! Ожидайте ответ на указанную вами электронную почту."
	case 2:
		return "Вы недавно уже отправляли вопрос. Проверьте пожалуйста вашу электронную почту или попробуйте позже."
	case 3:
		return "Антиспам система выявила с вашей стороны подозрительные действия и заблокировала отправку. Попробуйте позже или свяжитесь с нами вручную."
	default:
		return ""
	}
}
