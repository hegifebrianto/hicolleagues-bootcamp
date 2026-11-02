-- Jalankan ini dulu di MySQL (lewat mysql CLI, phpMyAdmin, atau
-- TablePlus/DBeaver) sebelum menjalankan server.

CREATE DATABASE IF NOT EXISTS tugasdb;
USE tugasdb;

CREATE TABLE IF NOT EXISTS buku (
    id INT AUTO_INCREMENT PRIMARY KEY,
    judul TEXT NOT NULL,
    penulis TEXT NOT NULL,
    tahun INT NOT NULL
);

-- Data contoh (opsional, supaya GET /api/buku langsung ada isinya)
INSERT INTO buku (judul, penulis, tahun) VALUES
    ('Laskar Pelangi', 'Andrea Hirata', 2005),
    ('Bumi Manusia', 'Pramoedya Ananta Toer', 1980),
    ('Filosofi Teras', 'Henry Manampiring', 2019);
