# Student Management REST API

API ini dibangun menggunakan Go dan framework Fiber untuk mengelola data entitas mahasiswa, menggunakan arsitektur Repository Pattern dan PostgreSQL sebagai basis data.

## Skema Tabel Basis Data

Aplikasi ini membutuhkan tabel `students` dengan struktur sebagai berikut:

| Kolom | Tipe Data | Batasan (Constraint) | Keterangan |
| :--- | :--- | :--- | :--- |
| `id` | SERIAL | PRIMARY KEY | Dihasilkan otomatis oleh DB |
| `username` | VARCHAR(50) | NOT NULL, UNIQUE | Digunakan untuk identitas/NIM |
| `email` | VARCHAR(255)| NOT NULL | Unik |
| `password` | VARCHAR(255)| NOT NULL | - |
| `is_active` | BOOLEAN | NOT NULL, DEFAULT TRUE | Status keaktifan |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW()| Waktu pencatatan |

## Cara Menyiapkan Basis Data

1. Pastikan PostgreSQL sudah terinstal dan berjalan di komputer.
2. Buat database baru melalui pgAdmin atau terminal psql (misal: `praktikum_backend`).
3. Jalankan *query* SQL yang terdapat di dalam file `migrations/001_create_students.sql` pada database tersebut untuk membuat tabel dan indeks yang dibutuhkan.

## Variabel Environment

Aplikasi ini membutuhkan beberapa variabel *environment* agar dapat berjalan. Buat file bernama `.env` di *root* direktori proyek, lalu isi dengan format berikut (sesuaikan nilai dengan pengaturan PostgreSQL di komputermu):

```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=masukkan_password_database_disini
DB_NAME=praktikum_backend