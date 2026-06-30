# GoCars

A website that showcases information about different car models, their specifications, manufacturers, and more.

## Setup and Installation

1.Clone the repository:
```bash  
git clone https://gitea.kood.tech/ashagaire/GoCars.git
cd GoCars
```

## Installation of required packages for database apis
```bash
cd api
```
- Install [NodeJS](http://nodejs.org)
- Install [NPM](https://www.npmjs.com/package/npm) package manager
- Install required packages: 

```bash
npm install
```

```bash 
make build
```



To run the server you need simply to execute the following command:
```bash
make run
```

```bash
Server is running on http://localhost:3000
```


Or if you want to customize the port of the server (port 3001 for example), run:

```bash
PORT=3001 make run
```
---
## Execute the GoCars web application in root folder

```bash
cd GoCars 
```

```bash
go run . 
```

```bash 
Server is running at http://localhost:8080 ...
```

# Usage Guide

## Browse Cars

Open:

```
http://localhost:8080
```

The homepage displays all available car models.

## Search

Use the search bar to find cars by model name.

Examples:

- Corolla
- BMW
- Civic

## Filter Cars

Use the filter panel to narrow results by:

- Manufacturer
- Category
- Minimum year
- Maximum year

Filters can be combined with the search function.

## View Car Details

Click on any car card to view:

- Engine
- Horsepower
- Transmission
- Drivetrain
- Manufacturer
- Category
- Model year

## Recommended Cars

Each car details page displays similar recommended vehicles based on shared characteristics.

---

# Cars API

The API exposes the following endpoints:

```
GET /api/models
GET /api/models/{id}

GET /api/manufacturers
GET /api/manufacturers/{id}

GET /api/categories
GET /api/categories/{id}
```

Images are available from:

```
/api/images
```

---

# Project Structure

```
GoCars/
│
├── api/                 # Node.js REST API
├── handlers/            # HTTP handlers
├── models/              # Data models
├── services/            # Business logic
├── templates/           # HTML templates
├── static/              # CSS and assets
├── main.go              # Application entry point
└── README.md
```

---

# Additional Features

Besides the required functionality, the project also includes several improvements:

- Responsive user interface
- Search by car model name
- Multiple filter options
- Combined search and filtering
- Car recommendation section
