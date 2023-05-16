# Menggunakan base image golang:latest
FROM golang:latest

# Menyalin seluruh kode sumber ke dalam image
COPY . .

# Mengatur direktori kerja
WORKDIR /Back-end-API-GoLang-case-Todo-list

# Membangun aplikasi
RUN go mod download
RUN go build -o main .

# Menjalankan aplikasi ketika container dijalankan
CMD ["./main"]




