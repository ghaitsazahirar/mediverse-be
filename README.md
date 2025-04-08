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

---

### 🌀 Git Flow

Proyek ini mengikuti _Git Flow_ sebagai strategi branching untuk pengembangan yang terstruktur dan kolaboratif. Berikut adalah alur awal penggunaan Git Flow pada proyek ini:

1. **Inisialisasi Git Repository**
   ```bash
   git init
   ```

2. **Menambahkan Remote Origin (GitHub)**
   ```bash
   git remote add origin https://github.com/ghaitsazahirar/mediverse-be.git
   ```

3. **Membuat dan Checkout ke Branch `develop`**
   ```bash
   git checkout -b develop
   ```

4. **Menambahkan Semua File dan Commit Awal**
   ```bash
   git add .
   git commit -m "Initial commit"
   ```

5. **Push ke Branch `develop` di GitHub**
   ```bash
   git push -u origin develop
   ```

---

### 📌 Struktur Branch

Struktur branching yang digunakan dalam Git Flow pada proyek ini:

- `main` → versi produksi (_production-ready_).
- `develop` → pengembangan aktif.
- `feature/*` → penambahan fitur baru.
- `release/*` → persiapan rilis ke `main`.
- `hotfix/*` → perbaikan cepat pada `main`.

> Semua branch `feature`, `release`, dan `hotfix` akan selalu di-*merge* ke `develop`, dan jika sudah siap rilis, akan di-*merge* ke `main` sesuai diagram Git Flow.

---
