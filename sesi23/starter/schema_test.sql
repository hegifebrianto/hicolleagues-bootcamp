-- Database TERPISAH khusus integration test (jangan pakai bukudb yang berisi data asli).
--   mysql -u root -p < schema_test.sql
CREATE DATABASE IF NOT EXISTS bukudb_test CHARACTER SET utf8mb4;
USE bukudb_test;

CREATE TABLE IF NOT EXISTS buku (
    id      INT AUTO_INCREMENT PRIMARY KEY,
    judul   VARCHAR(200) NOT NULL,
    penulis VARCHAR(120) NOT NULL,
    tahun   INT          NOT NULL
);
