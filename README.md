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
- Install required packages: 
```bash 
make build
```

To run the server you need simply to execute the following command:
```bash
make run
```

```
Or if you want to customize the port of the server (port 3001 for example), run:
```bash
PORT=3001 make run
```

## Execute the GoCar web application in root folder

```bash
go run . 
```

```bash 
Server is running at http://localhost:8080 ...
```
