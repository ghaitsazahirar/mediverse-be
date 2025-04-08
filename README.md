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

Oke sip! Ini aku buatin bagian **Git Flow** untuk dimasukin ke README `mediverse-be`, berdasarkan alur kamu waktu awal push ke GitHub dan pakai flow seperti di diagram tadi:

---

### 🌀 Git Flow

Proyek ini mengikuti _Git Flow_ sebagai strategi branching untuk pengembangan yang terstruktur dan kolaboratif. Berikut alur awal penggunaan Git Flow pada proyek ini:

1. **Inisialisasi Git Repository**
   ```
   git init
   ```

2. **Menambahkan Remote Origin (GitHub)**
   ```
   git remote add origin https://github.com/ghaitsazahirar/mediverse-be.git
   ```

3. **Membuat dan Checkout ke Branch `develop`**
   ```
   git checkout -b develop
   ```

4. **Menambahkan Semua File dan Commit Awal**
   ```
   git add .
   git commit -m "Initial commit"
   ```

5. **Push ke Branch `main` di GitHub**
   ```
   git push -u origin main
   ```

---

### 📌 Struktur Branch

Git Flow pada proyek ini mengikuti struktur branch berikut:

- `main` (atau `master`) → versi produksi.
- `develop` → pengembangan aktif.
- `feature/*` → penambahan fitur baru.
- `release/*` → persiapan rilis ke `main`.
- `hotfix/*` → perbaikan cepat pada `main`.

> Branch `feature`, `release`, dan `hotfix` akan selalu di-*merge* ke `develop` (dan `main` jika perlu) sesuai dengan diagram Git Flow berikut:

