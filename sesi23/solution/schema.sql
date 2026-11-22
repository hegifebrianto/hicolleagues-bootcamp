-- mysql -u root -p < schema.sql
CREATE DATABASE IF NOT EXISTS bukudb CHARACTER SET utf8mb4;
USE bukudb;

CREATE TABLE IF NOT EXISTS buku (
    id      INT AUTO_INCREMENT PRIMARY KEY,
    judul   VARCHAR(200) NOT NULL,
    penulis VARCHAR(120) NOT NULL,
    tahun   INT          NOT NULL
);

INSERT INTO buku (judul, penulis, tahun)
SELECT * FROM (
    SELECT 'Laskar Pelangi' AS judul, 'Andrea Hirata' AS penulis, 2005 AS tahun
    UNION ALL SELECT 'Bumi Manusia', 'Pramoedya Ananta Toer', 1980
    UNION ALL SELECT 'Filosofi Teras', 'Henry Manampiring', 2019
) AS contoh
WHERE NOT EXISTS (SELECT 1 FROM buku);
