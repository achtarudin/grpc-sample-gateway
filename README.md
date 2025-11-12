# 🌐 gRPC Sample Gateway

**HTTP/REST Gateway dengan Swagger UI untuk gRPC Microservices**

[![Go](https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![gRPC Gateway](https://img.shields.io/badge/gRPC_Gateway-v2.27.3-green?style=flat)](https://github.com/grpc-ecosystem/grpc-gateway)
[![Gorilla Mux](https://img.shields.io/badge/Gorilla_Mux-v1.8.1-blue?style=flat)](https://github.com/gorilla/mux)
[![Swagger](https://img.shields.io/badge/Swagger-UI-85EA2D?style=flat&logo=swagger)](https://swagger.io/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://docker.com/)

## 🌐 Live Demo & Documentation

**📚 Live API Documentation:** https://grpc-gateway.cutbray.tech/doc/  
**🚀 gRPC Server Backend:** https://grpc-server.cutbray.tech  
**📄 OpenAPI Spec:** https://grpc-gateway.cutbray.tech/swagger.json

> **Catatan:** Akses dokumentasi interaktif dan testing interface langsung melalui link di atas!

## 📋 Deskripsi

Gateway ini adalah **jembatan** antara aplikasi web biasa dan server gRPC. Tugasnya adalah menerjemahkan permintaan HTTP biasa menjadi panggilan gRPC, sehingga developer web bisa menggunakan API gRPC dengan mudah melalui HTTP/REST seperti biasa.

**Untuk Portofolio:** Project ini adalah hasil belajar dan eksperimen dalam:
- 🌉 **API Gateway** - Mencoba menghubungkan HTTP dengan gRPC
- 📚 **Dokumentasi API** - Belajar membuat dokumentasi dengan Swagger
- 🔄 **Integrasi Layanan** - Mengintegrasikan beberapa service dalam satu gateway  
- 🐳 **Deployment** - Belajar setup Docker dan konfigurasi production
- 🏗️ **Arsitektur Microservices** - Mencoba membangun sistem terdistribusi
- ⚡ **Fitur Production** - Implementasi shutdown yang aman dan konfigurasi environment

## ✨ Fitur Berdasarkan Struktur Project

### 🔧 **Fitur Utama Gateway**
- 🌐 **Penerjemah HTTP ke gRPC** - Mengubah request HTTP biasa jadi panggilan gRPC
- 🎯 **Router HTTP** - Mengatur rute HTTP dengan fleksibel
- 📚 **Dokumentasi Swagger** - Halaman dokumentasi interaktif di `/doc/`
- 📄 **File API Spec** - Spesifikasi lengkap API di `/swagger.json`
- ⚙️ **Konfigurasi Mudah** - Setting lewat environment variables
- 🛡️ **Shutdown Aman** - Matikan service dengan aman dalam 5 detik

### 🚀 **Service yang Terdaftar**
Gateway ini menghubungkan ke service berikut:
- ✅ **Hello Service** - Service sapaan dengan 4 cara komunikasi berbeda
- ✅ **Resiliency Service** - Service untuk testing ketahanan sistem

### 🏗️ **Architecture Components**

## 🏗️ Struktur Project (Actual)

```
grpc-sample-gateway/
├── cmd/
│   └── server/                     # Main gateway application
│       └── main.go                # Server bootstrap dengan Gorilla Mux
├── internal/
│   ├── adapter/
│   │   ├── gateway/              # gRPC Gateway configuration
│   │   │   ├── gateway.go        # Service registration (Hello + Resiliency)
│   │   │   └── gateway_config.go # Gateway configuration struct
│   │   ├── http/handler/         # HTTP handlers
│   │   │   ├── grpc_gateway_handler.go  # gRPC Gateway routing
│   │   │   └── swagger_handler.go       # Swagger UI + OpenAPI serving
│   │   └── logging/              # Structured logging
│   ├── helper/                   # Environment utilities
│   └── port/                     # Interface definitions
├── docker/
│   ├── go.dev.Dockerfile        # Development container
│   └── go.prod.Dockerfile       # Production container dengan Traefik labels
├── docker-compose.yml           # Development (port 5000:5045)
├── docker-compose.prod.yml      # Production dengan Traefik routing
├── Makefile                     # Build automation
├── .env.example                 # Environment template
└── go.mod                       # Dependencies (Go 1.24.4)
```

### 🔄 **Pengaturan Konfigurasi**

Setting yang bisa diatur lewat environment variables:
- **`GRPC_REMOTE_SERVER`** (default: `localhost:7000`) - Alamat server gRPC
- **`GRPC_TLS`** (default: `false`) - Aktifkan enkripsi TLS atau tidak
- **`GATEWAY_PORT`** (default: `8081`) - Port HTTP untuk gateway

### 🌐 **Endpoint yang Tersedia**

**Halaman Utama Gateway:**
- `GET /doc/` - Halaman dokumentasi Swagger yang interaktif
- `GET /swagger.json` - File spesifikasi API lengkap

**Endpoint REST Otomatis** (dibuat dari gRPC):
- `POST /hello/v1/say-hello` - Service sapaan biasa
- `POST /hello/v1/say-many-hellos` - Service sapaan berulang dari server
- `POST /hello/v1/say-hello-to-everyone` - Service sapaan untuk banyak orang  
- `POST /hello/v1/say-hello-continuous` - Service sapaan dua arah real-time
- Endpoint resiliency service (untuk testing sistem)

## 🚀 Quick Start

### Yang Dibutuhkan
- Go versi 1.24.4 ke atas
- Server gRPC harus jalan (grpc-sample-server) atau ubah setting `GRPC_REMOTE_SERVER`
- gow (tool untuk auto-reload saat development)

### 💻 Setup Development

```bash
# Download project
git clone https://github.com/achtarudin/grpc-sample-gateway.git  
cd grpc-sample-gateway

# Install tool dan dependensi
make install-tools    # Install gow untuk auto-reload
make install-deps     # Install semua library Go yang dibutuhkan

# Jalankan server development (auto-reload)
make dev-server       # Gateway akan restart otomatis kalau ada perubahan code

# Build untuk production
make build-server     # Buat file executable di ./bin/grpc-sample-gateway
make prod-server      # Build + jalankan langsung
```

### ⚙️ Setting Environment

Contoh file `.env`:
```env
# Konfigurasi Gateway
GATEWAY_PORT=8081
GRPC_REMOTE_SERVER=localhost:7000  # atau grpc-sample-server-dev:9000 untuk Docker
GRPC_TLS=false

# Setting development (dari .env.example):
GRPC_REMOTE_SERVER="grpc-sample-server-dev:9000"
```

## 🐳 Deployment dengan Docker

### Mode Development

```bash
# Buat network Docker (kalau belum ada)
docker network create grpc_sample_network

# Jalankan container development (port 5000:5045)
docker-compose up --build dev
```

**Setup Development:**
- **Container:** `grpc-sample-gateway-dev`
- **Network:** `grpc_sample_network` (sama dengan grpc-server)
- **Port:** `5000:5045` (akses lewat localhost:5000)
- **Auto-reload:** Code otomatis update tanpa restart container

### Mode Production

```bash
# Deploy production dengan Traefik (reverse proxy)
docker-compose -f docker-compose.prod.yml up --build -d
```

**Setup Production:**
- **Image:** `grpc-sample-gateway:latest` (bisa diubah lewat environment)
- **Network:** `traefik-network` untuk reverse proxy
- **Port:** `6000:5045` (internal, diakses lewat domain)
- **Domain:** `grpc-gateway.cutbray.tech` (otomatis lewat Traefik)
- **Config:** File `.env.prod` di-mount sebagai readonly

## Run with Docker Compose (dev)

The provided `docker-compose.yml` supports development with live code mounts and a shared Go module cache. It expects an external Docker network named `grpc_sample_network` (so the gateway can reach the server container by name if you run the server with the same network).

Steps:

1) Create the external network once (if you don’t already have it):
	- `docker network create grpc_sample_network`
2) Provide a `.env` file (or environment variables) containing at least `GATEWAY_PORT`.
3) Start the dev container:
	- `docker compose up --build dev`

The gateway will be available on `http://localhost:${GATEWAY_PORT}`.

## Run with Docker Compose (prod-like)

`docker-compose.prod.yml` builds a production image and exposes the service on port 6000 by default. It includes Traefik labels for routing under `grpc-gateway.cutbray.tech` when connected to a Traefik-managed network.

- Image name: `${IMAGE_NAME:-grpc-sample-gateway}:${IMAGE_VERSION:-latest}`
- Container port: `6000` (env `GATEWAY_PORT=6000`)

Example:

1) Ensure your Traefik network exists and is named `traefik-network`.
2) Build and run:
	- `docker compose -f docker-compose.prod.yml up --build -d`

## 🧪 Testing & Integration

### Perintah Make yang Tersedia:
```bash
make install-tools     # Install gow untuk live reload
make install-deps      # Update dependencies + grpc-sample module
make dev-server        # Development server dengan live reload
make build-server      # Build production binary
make prod-server       # Build + run production binary
make clean-server-bin  # Clean build artifacts
```

### 🔗 Ketergantungan Service
- **gRPC Backend:** Gateway membutuhkan grpc-sample-server yang bisa diakses di `GRPC_REMOTE_SERVER`
- **Shared Module:** `github.com/achtarudin/grpc-sample` untuk generated gateway handlers
- **OpenAPI Embedded:** Dokumentasi Swagger yang sudah disediakan dari shared module

### 🌐 Cara Kerja Integration
```
HTTP Client → Gateway (port 8081) → gRPC Server (port 9000)
     ↓
Swagger UI (/doc/) + OpenAPI (/swagger.json)
```

## 🔍 Fitur Unggulan

### 1. **Dokumentasi Swagger Terintegrasi**
- Spesifikasi API sudah tersedia dari module `grpc-sample`
- Memungkinkan testing API langsung dari browser lewat Swagger UI
- Fitur persistent auth untuk kemudahan testing
- Tampilan UI yang sederhana (model API otomatis tersembunyi)

### 2. **Registrasi Service Otomatis**
Gateway mencoba menghubungkan ke service:
- **Hello Service** - 4 cara komunikasi gRPC dengan mapping HTTP
- **Resiliency Service** - Tool untuk testing ketahanan sistem

### 3. **Beberapa Fitur Production**
- **Integrasi Traefik** - Routing lewat label Docker
- **Shutdown Aman** - Menunggu 5 detik untuk cleanup sebelum berhenti
- **Konfigurasi Fleksibel** - Setting yang berbeda untuk development vs production
- **Dukungan TLS** - Opsi enkripsi untuk koneksi ke gRPC backend

## 🤝 Integrasi dengan Sistem Lain

Project ini adalah bagian dari percobaan **ekosistem gRPC Microservices** yang terdiri dari:

- 📦 **[grpc-sample](https://github.com/achtarudin/grpc-sample)** - Definisi API dan code generator
- 🚀 **[grpc-server.cutbray.tech](https://grpc-server.cutbray.tech)** - Server gRPC backend  
- 🌐 **[grpc-gateway.cutbray.tech/doc/](https://grpc-gateway.cutbray.tech/doc/)** - Gateway dengan dokumentasi (project ini)
- 📱 **grpc-sample-client** - Aplikasi client untuk panggil gRPC langsung

## 👨‍💻 Author

**Achtarudin**
- 🌐 Live Gateway: [grpc-gateway.cutbray.tech/doc/](https://grpc-gateway.cutbray.tech/doc/)
- 🚀 gRPC Backend: [grpc-server.cutbray.tech](https://grpc-server.cutbray.tech)
- 🐙 GitHub: [@achtarudin](https://github.com/achtarudin)

---

> **💼 Catatan:** Project ini adalah hasil belajar dan eksplorasi dalam **API Gateway**, **HTTP ke gRPC translation**, **dokumentasi API interaktif**, **Docker deployment**, dan **microservices integration**. Merupakan upaya untuk memahami dan mengimplementasikan gateway service dengan fitur **routing**, **service discovery**, dan **dokumentasi** yang memadai.

