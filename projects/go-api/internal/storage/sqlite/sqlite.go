package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/codersgyan/students-api/internal/config"
	"github.com/codersgyan/students-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
	// The underscore (_) is called a blank import. We don't call functions directly from go-sqlite3 in our code. However, Go's standard database/sql package is generic—it doesn't know how to talk to SQLite, MySQL, or Postgres out of the box. By using _, we force the go-sqlite3 driver to execute its hidden init() function, which registers the SQLite engine driver with Go's standard database package behind the scenes.
	// Without that single underscore import, when you later try to run sql.Open("sqlite3", ...), Go will panic with the message: sql: unknown driver "sqlite3".
)
// It defines a structural type named Sqlite which encapsulates a pointer to a SQL database pool (*sql.DB).
// Sqlite is one concrete implementation of the Storage interface.
// It knows HOW to store data in SQLite.
// Handlers only know about Storage, not Sqlite.
type Sqlite struct {
	Db *sql.DB
}


// It opens a file-based connection pool to SQLite using the database file path stored inside our configuration (cfg.StoragePath). It then immediately triggers a raw SQL command (db.Exec) to guarantee that a students table exists in that database file before returning an initialized instance of our Sqlite struct.
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}


// ### 1. What is `s.Db.Prepare`? (The Template Maker)

// When you want a database to execute a query, you can do it two ways:

// 1. Send the full query with the data combined together instantly.
// 2. Send a **blueprint/template** of the query first, let the database compile it, and then send the data later.

// `s.Db.Prepare` is the second way. It takes a raw SQL string containing placeholders (`?` in SQLite) and pre-compiles it. Instead of executing the query immediately, it asks the database: *"Hey, look at the structure of this query, verify the syntax is correct, optimize the execution path, and hold onto it."*

// ---

// ### 2. What is `stmt`? (The Remote Control)

// `stmt` is short for **Statement**. Technically, it is a variable holding a pointer to an `sql.Stmt` struct object.

// Think of `stmt` as a **specialized remote control** bound exclusively to that pre-compiled SQL blueprint. Once the database creates the blueprint, it hands this `stmt` object back to Go. You use this `stmt` object to securely shoot data variables into those empty `?` placeholders and trigger the execution.

// ---

// ### 3. Concrete Example Comparison

// To understand the practical difference, let's look at the **Unsafe Way** versus the **Safe Way (`Prepare`)** that your project uses.

// #### The Bad & Unsafe Way (Without `Prepare`)

// Imagine a function written like this using pure string concatenation:

// ```go
// // DO NOT DO THIS IN PRODUCTION
// func (s *Sqlite) UnsafeCreateStudent(name string, email string, age int) (int64, error) {
//     // If name is: "John", email is: "john@example.com", age is: 20
//     // The query string becomes: INSERT INTO students (name, email, age) VALUES ('John', 'john@example.com', 20)
//     query := fmt.Sprintf("INSERT INTO students (name, email, age) VALUES ('%s', '%s', %d)", name, email, age)
    
//     result, err := s.Db.Exec(query)
//     // ...
// }

// ```

// * **The Problem:** If a malicious hacker registers with a name like `", "", 0); DROP TABLE students; --`, that name gets pasted directly into the SQL string. The database reads `DROP TABLE students;` as a valid direct instruction and deletes your data. This is a **SQL Injection Attack**.

// #### The Project's Way (With `Prepare` and `stmt`)

// Here is your actual code broken down line-by-line:

// ```go
// func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
//     // Step 1: Tell the database to compile the query structural template.
//     // The database sees the '?' and knows they represent absolute raw values, NOT SQL commands.
//     stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
//     if err != nil {
//         return 0, err
//     }
    
//     // Step 2: Ensure that we close the statement blueprint when the function finishes, 
//     // releasing database resources.
//     defer stmt.Close()

//     // Step 3: Use the 'stmt' remote control to execute the query.
//     // Go sends 'name', 'email', and 'age' separately. The database inserts them strictly as text strings.
//     result, err := stmt.Exec(name, email, age)
//     if err != nil {
//         return 0, err
//     }

//     lastId, err := result.LastInsertId()
//     if err != nil {
//         return 0, err
//     }

//     return lastId, nil
// }

// ```

// ---

// ### 4. Summary of Connections

// * **What it does:** `s.Db.Prepare` pre-compiles a SQL statement template containing placeholders, and returns an `sql.Stmt` tracker object (`stmt`).
// * **Why it is there:** Security and performance. It forces the database to split the query logic from the raw data variables, rendering SQL injection attacks impossible.
// * **How it connects:** In your project, `stmt` is created locally within functions like `CreateStudent` or `GetStudentById`. It safely binds your incoming HTTP request data parameters (like parameters parsed from the JSON body) before saving them down into the physical disk space allocated to SQLite.
// -------


// Concrete implementation of Storage.CreateStudent() for SQLite.
// When a Storage interface contains *Sqlite,
// storage.CreateStudent() dispatches to this method.
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

// 	When Go executes s.Db.Prepare(), it sends only the SQL query template to the database (e.g., SQLite). It does not send Karan's name, email, or age yet.

// The database takes this template:
// "INSERT INTO students (name, email, age) VALUES (?, ?, ?)"

// And does three things:

// Parses it: It reads the structure to ensure it's valid SQL syntax.

// Compiles it: It creates an optimized execution plan (a blueprint) for how to insert a row into the students table.

// Locks the doors (?): It treats the three ? marks as strictly locked placeholders. The database says: "I have created a secure blueprint. I am expecting exactly 3 raw pieces of data to fit into these slots later, but I will never interpret whatever comes next as SQL instructions."

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
