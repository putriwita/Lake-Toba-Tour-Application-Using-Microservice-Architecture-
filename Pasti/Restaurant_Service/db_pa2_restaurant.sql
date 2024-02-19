-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 05 Jun 2023 pada 06.19
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
-- Database: `db_pa2_restaurant`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `restaurants`
--

CREATE TABLE `restaurants` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nama_restaurant` varchar(255) DEFAULT NULL,
  `lokasi` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `harga` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `restaurants`
--

INSERT INTO `restaurants` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama_restaurant`, `lokasi`, `description`, `harga`) VALUES
(1, '2023-06-04 22:00:12', '2023-06-04 22:00:12', '2023-06-04 22:12:12', '', '', '', 20000),
(2, '2023-06-04 22:02:26', '2023-06-04 22:14:25', NULL, 'Bistro', 'Balige', 'Enak', 20000),
(3, '2023-06-04 22:03:00', '2023-06-04 22:03:00', '2023-06-04 23:26:36', '', 'Balige', 'Enak', 20000),
(4, '2023-06-04 22:03:50', '2023-06-04 22:03:50', '2023-06-04 23:26:01', '', 'Balige', 'Enak', 20000),
(5, '2023-06-05 07:53:50', '2023-06-05 07:53:50', NULL, '', 'Balige', 'Enak', 20000),
(6, '2023-06-05 07:54:56', '2023-06-05 07:54:56', NULL, '', 'Balige', 'Enak', 20000),
(7, '2023-06-05 07:55:50', '2023-06-05 07:55:50', NULL, 'haha', 'Balige', 'Enak', 20000),
(8, '2023-06-05 07:56:15', '2023-06-05 07:57:19', NULL, 'Bistro', 'Medan', 'Enak', 20000);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `restaurants`
--
ALTER TABLE `restaurants`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_restaurants_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `restaurants`
--
ALTER TABLE `restaurants`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
