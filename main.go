package main

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (api *API) ResetData(file string) error {
	jsonData, err := json.Marshal([]interface{}{})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/"+file, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) ChangeData(questions []model.Question) error {
	data, err := json.MarshalIndent(questions, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/questions.json", data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) ReadData() ([]model.Question, error) {
	file, err := ioutil.ReadFile("data/questions.json")
	if err != nil {
		return []model.Question{}, err
	}

	var questions []model.Question
	err = json.Unmarshal(file, &questions)
	if err != nil {
		return []model.Question{}, err
	}

	return questions, nil
}

func (api *API) AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
	//kerjain ini dulu ges
	var question model.Question
    err := json.NewDecoder(r.Body).Decode(&question)
    if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Bad Request"))
        return
    }

	fileJson, err := api.ReadData()
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad Request"))
		return
	}

	fileJson = append(fileJson, question)

	err = api.ChangeData(fileJson)
	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResp := model.ErrorResponse{
			Error: "Bad Request",
		}
		jsonErrResp, _ := json.Marshal(errorResp)
		w.Write(jsonErrResp)
		return
	}

    w.WriteHeader(http.StatusCreated)

	successResp := model.SuccessResponse{
		Message: "Question added!",
	}

	jsonSuccessResp, _ := json.Marshal(successResp)
	w.Write(jsonSuccessResp)
}

func (api *API) GetAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	fileJson, err := api.ReadData()
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad Request"))
		return
	}
	w.WriteHeader(http.StatusOK)

	resultFileJson, _ := json.Marshal(fileJson)

	w.Write(resultFileJson)
}

func (api *API) UpdateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	fileJson, err := api.ReadData()
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad Request"))
		return
	}

	qBody := model.Question{}
	err = json.NewDecoder(r.Body).Decode(&qBody)
	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResp := model.ErrorResponse{
			Error: "Bad Request",
		}
		jsonErrResp, _ := json.Marshal(errorResp)
		w.Write(jsonErrResp)
		return
	}

	// update
	isID := false
	for k, q := range fileJson {
		if q.ID == qBody.ID {
			isID = true
			fileJson[k] = qBody
			err = api.ChangeData(fileJson)
			if err!= nil {
				w.WriteHeader(http.StatusBadRequest)
                errorResp := model.ErrorResponse{
                    Error: "Bad Request",
                }
                jsonErrResp, _ := json.Marshal(errorResp)
                w.Write(jsonErrResp)
                return
            }
		}
	}

	if !isID {
		w.WriteHeader(http.StatusNotFound)
		errorResp := model.ErrorResponse{
            Error: "Question not found!",
        }
		jsonErrResp, _ := json.Marshal(errorResp)
		w.Write(jsonErrResp)
		return
	}
	w.WriteHeader(http.StatusOK)
	successResp := model.SuccessResponse{
        Message: "Question updated!",
    }
	jsonSuccessResp, _ := json.Marshal(successResp)
	w.Write(jsonSuccessResp)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/question/add", http.HandlerFunc(api.AddQuestionHandler))
	mux.Handle("/question/get-all", http.HandlerFunc(api.GetAllQuestionsHandler))
	mux.Handle("/question/update", http.HandlerFunc(api.UpdateQuestionHandler))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
