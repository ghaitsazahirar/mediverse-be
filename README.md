# 🩺 Mediverse Backend

Mediverse adalah platform reservasi medis online yang menghubungkan pasien dengan dokter dan fasilitas kesehatan melalui satu aplikasi terpadu.

Backend ini dibangun dengan bahasa **Go** menggunakan framework **Raiden**, serta memanfaatkan **Supabase** untuk layanan backend-as-a-service seperti autentikasi dan database.

---

## 🛠️ Teknologi yang Digunakan

- Go (Golang)
- Raiden Framework
- Supabase (PostgreSQL, Auth, Storage)
- GitHub (untuk versioning & kolaborasi)

---

## 📁 Struktur Folder

```bash
mediverse-be/
├── configs/            # File konfigurasi (app.yaml, env, dsb)
├── internal/
│   ├── domain/         # Definisi domain model dan interface
│   ├── handler/        # HTTP handler (Controller)
│   ├── service/        # Logika bisnis (Usecase)
│   └── repository/     # Akses ke database (Repository)
├── pkg/                # Package helper, utils, middleware
├── migrations/         # File migrasi database
├── go.mod              # Konfigurasi module Go
└── main.go             # Entry point aplikasi

## Cara Menjalankan Project
git clone https://github.com/ghaitsazahirar/mediverse-be.git
cd mediverse-be

go mod tidy                # Inisialisasi dependency Go
raiden configure           # Konfigurasi Supabase & Raiden
raiden run                 # Menjalankan server lokal
