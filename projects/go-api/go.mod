// Without a go.mod file, Go would have no idea how to locate your internal packages or where to download third-party libraries. It acts as the manifest file for your application's ecosystem.
// Whenever you write import "some-external-package" in any of your project's .go files, Go looks at go.mod to know which version of that package to use.
// Before starting any project we need to initialize a module using go mod init github.com/karangupta982/go-api


// This defines the unique name (or module path) of your project.
// In Go, this acts as the "root prefix" for any custom package you import within your project. For example, if you create a folder called config, you will import it elsewhere using "github.com/codersgyan/students-api/config"


module github.com/codersgyan/students-api


// This specifies the minimum Go language version required to compile and run this project. It ensures that any developer working on this project uses a compatible compiler.
go 1.22.5



// This is the list of external "toolboxes" (third-party dependencies) your project needs to function. Think of them as libraries you downloaded from the internet.
// indirect comments: This comment means these packages are not directly imported in your source code, but are dependencies of the other tools you did import.

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.1 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.23 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
