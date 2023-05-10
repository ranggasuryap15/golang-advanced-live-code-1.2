# Pemrograman Backend Lanjutan

## Quiz App 3 - Live Code

### Implementation technique

Siswa akan melaksanakan sesi live code di 25 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 25 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli submit`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Quiz App adalah sebuah aplikasi web yang digunakan untuk mengelola kumpulan pertanyaan (questions) pada sebuah quiz. Aplikasi ini memiliki beberapa fitur seperti:

- Menambahkan data pertanyaan baru ke dalam data
- Menampilkan seluruh pertanyaan yang ada pada data
- Memperbaharui sebuah pertanyaan berdasarkan ID pada data

Data dari aplikasi tersebut disimpan pada sebuah file dengan format JSON yaitu file `data/questions.json`

### Constraints

- Aplikasi ini hanya dapat digunakan secara lokal (localhost) dan hanya mendukung protokol HTTP.
- Aplikasi ini juga hanya memiliki tiga endpoint, yaitu:
  - `/question/add`
  - `/question/get-all`
  - dan `/question/update`.
- Gunakan struct yang sudah di sediakan pada file model `model/model.go` untuk memetakan data request atau response yaitu:
  - `Question`: untuk memetakan data soal menjadi JSON.
  - `ErrorResponse`: untuk memetakan error response menjadi JSON.
  - `SuccessResponse`: untuk memetakan sukses response menjadi JSON.
- Kamu bisa memanfaatkan fungsi yang telah disediakan untuk memanipulasi data file `data/question.json` yaitu:
  - `ChangeData(questions []model.Question) error` digunakan untuk mengubah semua data
  - `ReadData() ([]model.Question, error)` digunakan untuk membaca semua data

Aplikasi ini memiliki **3** fungsi utama yang harus dilengkapi, yaitu:

- `AddQuestionHandler` dengan end point `/question/add` digunakan untuk menambahkan pertanyaan baru ke dalam daftar pertanyaan. Pertanyaan yang ditambahkan harus dalam format JSON dan di dalam request body. Contoh:

  ```http
  POST /question/add HTTP/1.1
  Host: localhost:8080
  Content-Type: application/json

  {
    "id": "q1",
    "question": "Apa ibu kota Indonesia?",
    "options": ["Jakarta", "Bandung", "Surabaya", "Medan"],
    "answer": "Jakarta"
  }
  ```
  
  - Jika berhasil menambahkan pertanyaan, maka:
    - Berikan status code **201** (Created) dan response message JSON `{"message":"Question added!"}`
  - Jika format JSON yang diberikan pada request body tidak sesuai dengan struct `model.Question` atau ada kesalahan dalam decoding JSON, maka:
    - Berikan status code  **400** (Bad Request) dan response message JSON `{"error":"Bad Request"}`
  - Jika terjadi kesalahan saat membaca atau menulis file, maka:
    - Berikan status code  **500** (Internal Server Error)
- `GetAllQuestionsHandler` dengan end point `question/get-all` digunakan untuk menampilkan semua pertanyaan yang ada dalam daftar pertanyaan dan mengembalikannya dalam format JSON pada response body, dengan ketentuan:
  - Berikan status code **200** (OK) jika berhasil mengambil semua pertanyaan dan mengembalikan pertanyaan dalam format JSON pada response body. Contoh:

    ```http
    GET /question/get-all HTTP/1.1
    Host: localhost:8080

    [
      {
        "id": "q1",
        "question": "Apa ibu kota Indonesia?",
        "options": ["Jakarta", "Bandung", "Surabaya", "Medan"],
        "answer": "Jakarta"
      },
      {
        "id": "q2",
        "question": "Siapakah presiden pertama Indonesia?",
        "options": ["Soekarno", "Jokowi", "SBY", "Habibie"],
        "answer": "Soekarno"
      }
    ]
    ```

  - Jika terjadi kesalahan saat membaca atau menulis file, maka:
    - Berikan status code  **500** (Internal Server Error)
- `UpdateQuestionHandler` dengan end point `question/update` digunakan untuk mengubah atau memperbarui pertanyaan yang sudah ada di dalam file JSON dengan menggunakan metode HTTP `PUT`. Fungsi ini mengambil ID dari request body dengan format:

  ```json
  PUT /question/update HTTP/1.1
  Host: localhost:8080
  Content-Type: application/json

  {
    "id": "q1",
    "question": "Apa ibu kota Indonesia?",
    "options": ["Jakarta", "Bandung", "Surabaya", "Medan"],
    "answer": "Jakarta"
  }
  ```

  - Jika pertanyaan dengan ID yang diberikan ditemukan
    - Berikan status code **200** (OK) dan response message JSON `{"message":"Question updated!"}`

  - Jika pertanyaan dengan ID yang diberikan tidak ditemukan, maka:
    - Kembalikan status code **404** (Not Found) dan mengembalikan pesan `{"error":"Question not found!"}`
  - Jika terjadi kesalahan saat membaca atau menulis file, maka:
    - Berikan status code  **500** (Internal Server Error)

### Test Case Examples

#### Test Case 1

**Input**:

```http
POST /question/add HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "id": "q1",
    "question": "Apa ibu kota Indonesia?",
    "options": ["Jakarta", "Bandung", "Surabaya", "Medan"],
    "answer": "Jakarta"
}
```

**Expected Output / Behavior:**

- Jika request berhasil, server akan mengembalikan status code `201 Created` dan response body dalam format JSON seperti contoh di bawah ini:

  ```json
  {
      "message": "Question added!"
  }
  ```

- Jika terjadi kesalahan saat memproses request, server akan mengembalikan status code sesuai jenis kesalahan yang terjadi (misalnya `400 Bad Request` jika body request tidak sesuai format yang diharapkan) dan response body dalam format JSON seperti contoh di bawah ini:

  ```json
  {
      "error": "Error message"
  }
  ```

#### Test Case 2

**Input**:

```http
GET /question/get-all HTTP/1.1
Host: localhost:8080
```

**Expected Output / Behavior:**

```http
HTTP/1.1 200 OK
Content-Type: application/json

[
    {
        "id": "q1",
        "question": "Apa ibu kota Indonesia?",
        "options": ["Jakarta", "Bandung", "Surabaya", "Medan"],
        "answer": "Jakarta"
    },
    {
        "id": "q2",
        "question": "Siapakah presiden pertama Indonesia?",
        "options": ["Soekarno", "Jokowi", "SBY", "Habibie"],
        "answer": "Soekarno"
    }
]
```

#### Test Case 3

**Input**:

```http
PUT /question/update HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "id": "q1",
    "question": "Apa ibu kota Malaysia?",
    "options": ["Hongkong", "Kuala Lumpur", "Jakarta", "Cianjur"],
    "answer": "Kuala Lumpur"
}
```

**Expected Output / Behavior:**

- Response jika pertanyaan ditemukan:

  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json

  {
      "message": "Question updated!"
  }
  ```

- Response jika pertanyaan tidak ditemukan:

  ```http
  HTTP/1.1 404 Not Found
  Content-Type: application/json

  {
      "error": "Question not found!"
  }
  ```
