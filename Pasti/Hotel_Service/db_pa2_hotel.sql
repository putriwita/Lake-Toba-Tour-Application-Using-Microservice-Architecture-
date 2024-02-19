-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 05 Jun 2023 pada 06.17
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_pa2_hotel`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `hotels`
--

CREATE TABLE `hotels` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nama_hotel` varchar(255) DEFAULT NULL,
  `lokasi` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `harga` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `hotels`
--

INSERT INTO `hotels` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama_hotel`, `lokasi`, `description`, `harga`) VALUES
(1, '2023-06-04 23:03:09', '2023-06-04 23:03:09', NULL, '', '', '', 0),
(2, '2023-06-04 23:05:28', '2023-06-04 23:05:28', NULL, '', '', '', 0),
(3, '2023-06-04 23:06:13', '2023-06-04 23:06:13', NULL, '', '', '', 0),
(4, '2023-06-04 23:07:28', '2023-06-04 23:07:28', NULL, '', '', '', 0),
(5, '2023-06-04 23:08:16', '2023-06-04 23:08:16', '2023-06-05 08:01:15', '', '', '', 0),
(6, '2023-06-04 23:08:22', '2023-06-04 23:08:22', '2023-06-04 23:59:42', 'Labersa', 'Balige', 'Nyaman', 150000),
(7, '2023-06-04 23:28:33', '2023-06-04 23:28:33', '2023-06-04 23:54:56', 'asdasd', 'Balige', 'Nyaman', 150000),
(8, '2023-06-04 23:58:32', '2023-06-04 23:58:32', '2023-06-04 23:59:21', 'Oyo', 'Balige', 'Nyaman', 150000),
(9, '2023-06-05 07:59:55', '2023-06-05 08:00:40', '2023-06-05 08:00:53', 'Oyo Medan', 'Medan', 'Kotor', 15000000);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `hotels`
--
ALTER TABLE `hotels`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_hotels_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `hotels`
--
ALTER TABLE `hotels`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
