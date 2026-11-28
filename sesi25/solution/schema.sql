-- Jalankan sekali:  mysql -u root -p < schema.sql
CREATE DATABASE IF NOT EXISTS bukudb;
USE bukudb;

CREATE TABLE IF NOT EXISTS buku (
  id      BIGINT AUTO_INCREMENT PRIMARY KEY,
  judul   VARCHAR(200) NOT NULL,
  penulis VARCHAR(100) NOT NULL DEFAULT '',
  tahun   INT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id            BIGINT AUTO_INCREMENT PRIMARY KEY,
  nama          VARCHAR(100) NOT NULL,
  email         VARCHAR(150) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  role          VARCHAR(20)  NOT NULL DEFAULT 'user'
);

INSERT INTO buku (judul, penulis, tahun) VALUES
  ('Laskar Pelangi', 'Andrea Hirata', 2005),
  ('Bumi Manusia', 'Pramoedya Ananta Toer', 1980);

-- Akun demo, password keduanya: rahasia
-- (hash dibuat dengan: go run ./cmd/buathash rahasia)
INSERT INTO users (nama, email, password_hash, role) VALUES
  ('Admin', 'admin@mail.com', 'pbkdf2$100000$F5Xh8aKbBqWUcxwNTQSEkw$NdBXhDWjyfIzU7Y2JR5xV9x5bERBe2bTx/ZFIOoKa5o', 'admin'),
  ('Budi',  'budi@mail.com',  'pbkdf2$100000$EhI8tU3LQe0ChXgKoRCkcA$537q10OeKchorDffAeJyiyEDmRML1UuOTM6dJeB+GOE', 'user');
