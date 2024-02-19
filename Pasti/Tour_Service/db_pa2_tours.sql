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
-- Database: `db_pa2_tours`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `tours`
--

CREATE TABLE `tours` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `detail` varchar(255) DEFAULT NULL,
  `harga` int(10) UNSIGNED DEFAULT NULL,
  `jumlahorang` int(10) UNSIGNED DEFAULT NULL,
  `diskon` int(10) UNSIGNED DEFAULT NULL,
  `hargatotal` int(10) UNSIGNED DEFAULT NULL,
  `jumlah_hari` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tours`
--

INSERT INTO `tours` (`id`, `created_at`, `updated_at`, `deleted_at`, `image`, `nama`, `detail`, `harga`, `jumlahorang`, `diskon`, `hargatotal`, `jumlah_hari`) VALUES
(1, '2023-06-05 08:38:12', '2023-06-05 08:39:35', NULL, 'Malbas', 'Balige', 'Enak', 20000, 10, 10, 12000000, 20);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `tours`
--
ALTER TABLE `tours`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_tours_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `tours`
--
ALTER TABLE `tours`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
