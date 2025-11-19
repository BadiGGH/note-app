# Note App API

A simple RESTful API for managing notes, built with go and gin framework. Supports creating, reading, editing and deleting notes.

---

## Features

- Create new notes
- Read all notes or a single note by ID
- Update notes by ID
- Delete notes by ID
- Swagger API documentation included

---

## Getting Started

### Prerequisites

- Go 1.18+ installed
- MySQL or compatible database running

### Installation

1. Clone the repo

```bash
git clone https://github.com/BadiGGH/note-app.git
cd note-app
```

2. Set up your database and run the SQL schema in `/databse/note.sql`
 
3. Create the `.env` file with default values from `.example.env` file.(Change the values if you need to)

```bash
cp .example.env .env
```

4. Run the code

```bash
go run main.go
```

5. Visit `/swagger/index.html` and see the API documentation

---



