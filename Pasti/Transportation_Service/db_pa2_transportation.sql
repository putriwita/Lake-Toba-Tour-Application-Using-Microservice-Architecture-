-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 05 Jun 2023 pada 06.18
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
-- Database: `db_pa2_transportation`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `transportations`
--

CREATE TABLE `transportations` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `price` int(10) UNSIGNED DEFAULT NULL,
  `route` varchar(255) DEFAULT NULL,
  `duration` int(10) UNSIGNED DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transportations`
--

INSERT INTO `transportations` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `image`, `type`, `price`, `route`, `duration`, `location`) VALUES
(1, '2023-06-05 10:30:05', '2023-06-05 10:31:07', NULL, 'Oyo kisaran', 'test', 'ok', 15000000, 'Medan', 6, 'jakarta');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `transportations`
--
ALTER TABLE `transportations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transportations_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `transportations`
--
ALTER TABLE `transportations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
