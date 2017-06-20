package handlers

import (
	"goardparser/structs"
	"goardparser/errors"
	"goardparser/validators"
	"goardparser/utils"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"strings"
)

func IndexHandler(writer http.ResponseWriter, r *http.Request)  {
	stuff := "Hello goardparser!"
	utils.JSONResponse(writer, structs.GenericJSON{Stuff: stuff})
}

func ParseDataHandler(writer http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read the request body: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var requestData structs.RequestDataJSON

	if err := json.Unmarshal(body, &requestData); err != nil {
		errors.SendErrorMessage(writer,
			"Could not decode the request body as JSON",
			http.StatusBadRequest)
		return
	}

	if validators.IsValidRequestParams(writer, requestData) {

		channel := make(chan *structs.Board)
		go utils.ParseThread(requestData.Data, channel)

		data :=  <-channel

		if data.Error != nil {
			errors.SendErrorMessage(writer,
				"Thread does not exist",
				http.StatusBadRequest)
			return
		}

		responseJson := &structs.ResponseJSON{}

		for _, post := range data.Threads[0].Posts {

			for _, file := range post.Files{

				if strings.Contains(file.Name, ".webm"){
					responseJson.Files = append(responseJson.Files, file)
				}
			}
		}
		utils.JSONResponse(writer, responseJson)
	}
}
