# GoCars

GoCars is a web application built with Go that allows users to explore different car models, compare specifications, browse manufacturers, and search for vehicles based on multiple criteria. The application retrieves data from a REST API and presents it through a clean and responsive web interface.

---

# Project Overview

GoCars consists of two components:

### Go Web Application
The frontend is built with Go using HTML templates and serves dynamic web pages that allow users to browse and explore car information.

### Cars REST API
A Node.js + Express API provides data about:

- Car models
- Manufacturers
- Categories
- Vehicle images

The Go application consumes this API to render all pages.

## Features

- Browse all available cars
- Search and advanced filtering
- View detailed specifications for each vehicle
- Compare cars side-by-side
- View recently viewed cars
- Receive recommended cars based on the currently viewed vehicle

---

# Setup and Installation

## 1. Clone the repository

```bash
git clone https://gitea.kood.tech/ashagaire/GoCars.git
cd GoCars
```

## 2. Install and run the Cars API

Navigate to the API folder.

```bash
cd api
```

Install dependencies:

```bash
npm install
```

or

```bash
make build
```

Run the API:

```bash
make run
```

The API will run on:

```
http://localhost:3000
```

To use another port:

```bash
PORT=3001 make run
```

## 3. Run the Go application

Return to the project root:

```bash
cd ..
```

Start the application:

```bash
go run .
```

The web application will be available at:

```
http://localhost:8080
```

---

# Usage Guide

## Home Page

The homepage displays all available cars. Users can browse the collection, search for vehicles, apply filters, or navigate to other pages.

## Search & Filter

The search bar supports keyword searches across multiple car attributes, allowing users to quickly find relevant vehicles.

Searchable information includes:

- Model name
- Manufacturer
- Category
- Year
- Transmission
- Drivetrain
- Horsepower

Users can further narrow the results by applying filters for:

- Manufacturer
- Category
- Production year range
- Transmision
- Drivetrain
- Horsepower

Search and filters work together and update the displayed results accordingly.

## Car Details

Selecting a car opens its detail page, where users can view:

- Manufacturer info
- Category
- Engine
- Horsepower
- Transmission
- Drivetrain
- Production year

The page also displays:

- Recommended cars with similar characteristics
- Recently viewed cars for quick navigation back to previously visited vehicles

## Manufacturer Page

Users can browse manufacturers and view:

- Country of origin
- Founding year
- All cars produced by the selected manufacturer

## Compare Cars

Users can compare two vehicles side-by-side, including:

- Manufacturer
- Category
- Engine
- Horsepower
- Transmission
- Drivetrain
- Production year

This allows users to easily identify differences between two models.

---

# API Endpoints

The API provides the following endpoints:

```
GET /api/models
GET /api/models/{id}

GET /api/manufacturers
GET /api/manufacturers/{id}

GET /api/categories
GET /api/categories/{id}
```

Vehicle images are served from:

```
/api/images
```

---

# Project Structure

```
GoCars/
│
├── api/                # Node.js REST API
├── handlers/           # HTTP handlers
├── models/             # Data models
├── services/           # Business logic
├── templates/          # HTML templates
├── static/             # CSS, JavaScript and images
├── main.go             # Application entry point
└── README.md
```

---

# Additional Features

In addition to the required project functionality, GoCars includes:

- Full-text search across multiple vehicle attributes
- Multi-criteria filtering
- Car comparison page
- Manufacturer information page
- Recently viewed cars history
- Recommended cars based on the selected vehicle
