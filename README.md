# Esport Gacor: Aplikasi Manajemen Turnamen E-Sport

Selamat datang di **Esport Gacor**, sebuah aplikasi konsol (CLI) yang dirancang untuk mengelola tim, pemain, dan pertandingan dalam sebuah turnamen e-sport. Aplikasi ini memungkinkan pengguna untuk dengan mudah menambah, mengubah, dan menghapus data tim serta mencatat dan memperbarui hasil pertandingan untuk menghasilkan klasemen secara otomatis.

## 🚀 Fitur Utama

-   **Manajemen Tim**:
    -   Menambah tim baru beserta 5 pemainnya.
    -   Mengubah detail tim, termasuk nama tim, nama pelatih, dan data pemain.
    -   Menghapus tim dari daftar.
-   **Manajemen Pertandingan**:
    -   Mencatat hasil pertandingan baru antar tim.
    -   Mengedit skor pertandingan yang sudah ada.
    -   Menampilkan daftar semua pertandingan yang telah berlangsung.
-   **Klasemen Dinamis**:
    -   Menampilkan klasemen (peringkat) tim berdasarkan poin yang didapat.
    -   Poin otomatis diperbarui setiap kali hasil pertandingan ditambahkan atau diubah.
    -   Klasemen diurutkan dari poin tertinggi ke terendah.
-   **Tampilan Informasi**:
    -   Menampilkan daftar lengkap semua tim beserta pemain dan detailnya.
    -   Tampilan menu yang terstruktur dan mudah dinavigasi.

## 🛠️ Struktur Proyek

Proyek ini disusun dengan modularitas untuk kemudahan pengembangan dan pemeliharaan.

```
esportgacor/
├── main.go                  # Titik masuk utama aplikasi
├── go.mod                   # File modul Go
├── go.sum                   # Checksum dependensi
├── database/                # Inisialisasi & manajemen data
│   ├── constants.go
│   └── database.go
├── handlers/                # Logika menu & input pengguna
│   ├── match.go
│   ├── menu.go
│   └── team.go
├── models/                  # Definisi struktur data
│   ├── database.go
│   ├── match.go
│   └── team.go
└── utils/                   # Fungsi bantu (sorting, searching, dll.)
    ├── classement.go
    ├── match.go
    ├── menus.go
    ├── search.go
    ├── sort.go
    └── team.go
```

## ⚙️ Cara Menjalankan Aplikasi

1.  **Prasyarat**: Pastikan Anda telah menginstal [Go (Golang)](https://go.dev/doc/install) pada Device Anda.

2.  **Clone Repositori** (Jika sudah ada di Git):
    ```bash
    git clone <URL_REPOSITORI_ANDA>
    cd esportgacor
    ```

3.  **Jalankan Aplikasi**:
    Buka terminal atau command prompt, navigasikan ke direktori root proyek (`esportgacor/`), dan jalankan perintah berikut:
    ```bash
    go run main.go
    ```

4.  **Nikmati Aplikasinya!**
    Aplikasi akan menampilkan menu utama di konsol Anda. Ikuti petunjuk yang ada untuk mulai mengelola tim dan pertandingan.
