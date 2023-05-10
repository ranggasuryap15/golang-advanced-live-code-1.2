package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"strconv"

	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Quiz App", Ordered, func() {
	var server main.API
	BeforeEach(func() {
		server = main.NewAPI()
		server.ResetData("questions.json")
	})

	Describe("GET /question/add", func() {
		AfterEach(func() {
			server.ResetData("questions.json")
		})
		When("request is valid", func() {
			It("returns status code 201, response body 'Question added!' and add data to file 'data/question.json'", func() {
				newQuestion := model.Question{
					ID:       "q1",
					Question: "What is the capital city of Indonesia?",
					Options: []string{
						"Jakarta",
						"Bandung",
						"Surabaya",
					},
					Answer: "Jakarta",
				}
				reqBody, err := json.Marshal(newQuestion)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/question/add", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(http.StatusCreated))
				Expect(w.Body.String()).To(MatchRegexp("Question added!"))

				datas, err := server.ReadData()
				Expect(err).To(BeNil())
				Expect(datas).To(HaveLen(1))
				Expect(datas[0].ID).To(Equal("q1"))
				Expect(datas[0].Question).To(Equal("What is the capital city of Indonesia?"))
				Expect(datas[0].Options).To(Equal([]string{"Jakarta", "Bandung", "Surabaya"}))
				Expect(datas[0].Answer).To(Equal("Jakarta"))
			})
		})

		When("request has invalid JSON body", func() {
			It("returns status code 400", func() {
				reqBody := []byte(`{"invalid_json": }`)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/question/add", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
			})

			It("returns response body 'Bad Request'", func() {
				reqBody := []byte(`{"invalid_json": }`)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/question/add", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)

				Expect(w.Body.String()).To(MatchRegexp("Bad Request"))
			})
		})
	})

	Describe("GET /question/get-all", func() {
		BeforeEach(func() {
			for i := 0; i < 3; i++ {
				newQuestion := model.Question{
					ID:       "q" + strconv.Itoa(i+1),
					Question: "What is the capital city of Indonesia?",
					Options: []string{
						"Jakarta",
						"Bandung",
						"Surabaya",
					},
					Answer: "Jakarta",
				}
				reqBody, err := json.Marshal(newQuestion)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/question/add", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)
			}
		})

		AfterEach(func() {
			server.ResetData("questions.json")
		})

		When("request is valid", func() {
			It("returns status code 200", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/question/get-all", nil)
				server.Handler().ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})

			It("returns all questions in response body", func() {
				expectedQuestions := []model.Question{
					{
						ID:       "q1",
						Question: "What is the capital city of Indonesia?",
						Options: []string{
							"Jakarta",
							"Bandung",
							"Surabaya",
						},
						Answer: "Jakarta",
					},
					{
						ID:       "q2",
						Question: "What is the capital city of Indonesia?",
						Options: []string{
							"Jakarta",
							"Bandung",
							"Surabaya",
						},
						Answer: "Jakarta",
					},
					{
						ID:       "q3",
						Question: "What is the capital city of Indonesia?",
						Options: []string{
							"Jakarta",
							"Bandung",
							"Surabaya",
						},
						Answer: "Jakarta",
					},
				}

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/question/get-all", nil)
				server.Handler().ServeHTTP(w, r)

				var questions []model.Question
				err := json.Unmarshal(w.Body.Bytes(), &questions)
				Expect(err).To(BeNil())
				Expect(questions).To(Equal(expectedQuestions))
			})
		})
	})

	Describe("PUT /question/update", func() {
		BeforeEach(func() {
			newQuestion := model.Question{
				ID:       "q1",
				Question: "What is the capital city of Indonesia?",
				Options: []string{
					"Jakarta",
					"Bandung",
					"Surabaya",
				},
				Answer: "Jakarta",
			}
			reqBody, err := json.Marshal(newQuestion)
			Expect(err).To(BeNil())

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/question/add", bytes.NewBuffer(reqBody))
			server.Handler().ServeHTTP(w, r)
		})

		AfterEach(func() {
			server.ResetData("questions.json")
		})

		When("request is valid", func() {
			It("returns status code 200, response body 'Question updated!' and update file 'data/questions.json'", func() {
				updatedQuestion := model.Question{
					ID:       "q1",
					Question: "What is the largest city in Indonesia?",
					Options: []string{
						"Jakarta",
						"Surabaya",
						"Bandung",
					},
					Answer: "Jakarta",
				}
				reqBody, err := json.Marshal(updatedQuestion)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPut, "/question/update", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(MatchRegexp("{\"message\":\"Question updated!\"}"))

				datas, err := server.ReadData()
				Expect(err).To(BeNil())
				Expect(datas).To(HaveLen(1))
				Expect(datas[0].ID).To(Equal("q1"))
				Expect(datas[0].Question).To(Equal("What is the largest city in Indonesia?"))
				Expect(datas[0].Options).To(Equal([]string{"Jakarta", "Surabaya", "Bandung"}))
				Expect(datas[0].Answer).To(Equal("Jakarta"))
			})
		})

		When("request has invalid JSON body", func() {
			It("returns status code 400", func() {
				reqBody := []byte(`{"invalid_json": }`)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPut, "/question/update", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
			})

			It("returns response body 'Bad Request'", func() {
				reqBody := []byte(`{"invalid_json": }`)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPut, "/question/update", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)

				Expect(w.Body.String()).To(MatchRegexp("Bad Request"))
			})
		})

		When("question with given ID does not exist", func() {
			It("returns status code 404", func() {
				updatedQuestion := model.Question{
					ID:       "nonexistent_id",
					Question: "What is the largest city in Indonesia?",
					Options: []string{
						"Jakarta",
						"Surabaya",
						"Bandung",
					},
					Answer: "Jakarta",
				}
				reqBody, err := json.Marshal(updatedQuestion)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPut, "/question/update", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})

			It("returns response body 'Question not found'", func() {
				updatedQuestion := model.Question{
					ID:       "nonexistent_id",
					Question: "What is the largest city in Indonesia?",
					Options: []string{
						"Jakarta",
						"Surabaya",
						"Bandung",
					},
					Answer: "Jakarta",
				}
				reqBody, err := json.Marshal(updatedQuestion)
				Expect(err).To(BeNil())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPut, "/question/update", bytes.NewBuffer(reqBody))
				server.Handler().ServeHTTP(w, r)

				Expect(w.Body.String()).To(MatchRegexp("{\"error\":\"Question not found!\"}"))
			})
		})
	})

})
