# REST API Rental menggunakan Golang

| Nama | Zuyinatin Khofifah |
| ---- | ------------------ |

## Deskripsi

Program ini adalah implementasi REST API untuk sistem penyewaan properti, yang dibangun menggunakan bahasa pemrograman Golang. Program ini adalah proyek unjuk keterampilan sebagai bagian dari pelatihan Alterra Academy untuk berkarir sebagai Backend Engineer dengan Golang.

Program ini menggunakan beberapa pustaka dan alat-alat berikut:

- **Gorm**: Sebuah ORM (Object Relational Mapping) untuk berinteraksi dengan database.
- **Viper**: Pustaka konfigurasi yang kuat untuk mengelola konfigurasi aplikasi.
- **Gorilla/Mux**: Router HTTP yang kuat untuk menangani permintaan HTTP.
- **Logrus**: Pustaka logging yang fleksibel dan kuat.
- **Air**: Alat live reload untuk pengembangan aplikasi Go.

## Persyaratan

Sebelum Anda memulai, pastikan Anda telah menginstal Golang di komputer Anda. Pastikan juga bahwa Anda memiliki database yang diaktifkan dan konfigurasi koneksi database telah diatur.
Install Golang https://go.dev/doc/install

- Go (minimal versi 1.14)
- [GORM](https://gorm.io/) - Toolkit ORM untuk Golang
- [Viper](https://github.com/spf13/viper) - Library konfigurasi untuk Golang
- [Gorilla Mux](https://github.com/gorilla/mux) - Router HTTP dan URL matcher untuk Golang
- [Logrus](https://github.com/sirupsen/logrus) - Pustaka logging yang kuat untuk Golang
- [Air](https://github.com/cosmtrek/air) - Alat live-reload untuk aplikasi Golang

## Instalasi

1. Klon repositori ini ke komputer Anda:
   ```bash
   git clone https://github.com/username/go-rest-api-rental.git
   ```
2. Masuk ke direktori proyek:
   ```bash
   cd go-rest-api-rental
   ```
3. Install dependensi dengan Go Modules:
   ```bash
   go mod tidy
   ```
4. Buat konfigurasi (.env) sesuai kebutuhan Anda.
5. Jalankan program menggunakan Air untuk live-reload:
   ```bash
   air
   ```

Program akan berjalan dan API akan dapat diakses di 'http://localhost:8080'.

## Penggunaan

Berikut merupakan penggunaan endpoint API:

#### Properties

- Mengambil Daftar Properti: `GET /properties`
- Mengambil Detail Properti: `GET /properties/{id}/detail`
- Menambahkan Data Properti: `POST /properties`
- Mengubah Data Properti: `PUT /properties/{id}/update`
- Menambahkan Data Properti: `DELETE /properties/{id}/delete`

#### Tenant

- Mengambil Daftar Tenant: `GET /tenants`
- Mengambil Detail Tenant: `GET /tenants/{id}/detail`
- Menambahkan Data Tenant: `POST /tenants`
- Mengubah Data Tenant: `PUT /tenants/{id}/update`
- Menambahkan Data Tenant: `DELETE /tenants/{id}/delete`

## Credits

This project was developed by Zuyinatin Khofifah.
Special thanks to Alterra Academy for providing the training and guidance.
