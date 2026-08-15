// The backticks (e.g., `yaml:"address"`) are called struct tags. They tell the configuration parser (cleanenv) how to translate keys from your local.yaml file into variables in Go.

// env-required:"true" tells Go that if this value is missing from the configuration, it should throw an error rather than booting up with empty/blank values
// This acts as a type-safe parser. Instead of manually parsing a file text line-by-line, we let a library do the heavy lifting of matching keys from a file to structured variables.
// env: "dev" maps to cfg.Env because of yaml:"env".

// storage_path: "storage/storage.db" maps to cfg.StoragePath because of yaml:"storage_path".

// http_server: -> address: "localhost:8082" maps to cfg.HTTPServer.Addr because HTTPServer is embedded in Config and maps to address.

// It gives you two flexible ways to tell the app where local.yaml is located:

// Method A (Env Var): By setting an environment variable on your computer (CONFIG_PATH="config/local.yaml").

// Method B (CLI Flag): By passing a flag when running your application in the terminal:
// go run cmd/students-api/main.go -config=config/local.yaml


// ## 1. How does Go find `local.yaml` if we don't "import" it?

// In programming, **importing** is only for other code files (like packages).

// Configuration files (like `.yaml`, `.json`, or `.txt`) are **data files**, not code. Go doesn't compile them into your app. Instead, it reads them at runtime (when the program is actually running), just like opening a text document in Notepad.

// 1. **You tell the program where the file is:** When you run your program in the terminal, you pass the file path as a flag:
// ```bash
// go run cmd/students-api/main.go -config=config/local.yaml

// ```


// 2. **Go grabs that path:** Inside `config.go`, this line captures that string:
// ```go
// flags := flag.String("config", "", "path to the configuration file")
// // ...
// configPath = *flags // This now holds the string "config/local.yaml"

// ```


// 3. **Go opens the file:** We pass that path to our third-party library (`cleanenv`):
// ```go
// cleanenv.ReadConfig(configPath, &cfg)

// ```


// Behind the scenes, `cleanenv` says: *"Ah, `configPath` is `config/local.yaml`. Let me open that file on the hard drive, read all of its text, and parse it."*

// ---

// ## 2. How does the mapping actually happen? (What are those backticks?)

// Inside `config.go`, you wrote:

// ```go
// StoragePath string `yaml:"storage_path" env-required:"true"`

// ```

// Everything inside the backticks ``...`` is called a **Struct Tag** in Go.

// Think of Struct Tags as **translation sticky notes** for external libraries.

// * **Without the tag:** Go only knows about a variable named `StoragePath` (using PascalCase).
// * **The YAML file:** The YAML file has a key named `storage_path` (using snake_case).
// * **The Bridge:** The tag `yaml:"storage_path"` is a message to `cleanenv`. It says:
// > *"Hey cleanenv, when you are reading the YAML file, look for a key named `storage_path`. When you find it, grab its value (e.g., `"storage/storage.db"`) and put it into my Go variable `StoragePath`."*



// If you didn't have that tag, `cleanenv` would look for a key literally named `StoragePath` in your YAML file, fail to find it, and your program would crash or leave the variable empty.

// ---

// ## 3. The Step-by-Step Data Flow (Trace the journey)

// Let's watch how the database path goes from a text file on your hard drive into your Go program's memory:

// | Step | What Happens | Code / File State |
// | --- | --- | --- |
// | **1. Storage** | The setting sits on your hard drive as raw text. | **`local.yaml`**:<br>

// <br>`storage_path: "storage/storage.db"` |
// | **2. Startup** | You launch the app and point it to the file path. | **Terminal**:<br>

// <br>`-config=config/local.yaml` |
// | **3. Capture** | Go captures the path as a string variable. | **`config.go`**:<br>

// <br>`configPath` is now `"config/local.yaml"` |
// | **4. Reading** | `cleanenv` reads the YAML file and looks at your struct tags. | **`config.go`**:<br>

// <br>Matches `yaml:"storage_path"` tag to `storage_path:` in the file. |
// | **5. Assignment** | `cleanenv` copies `"storage/storage.db"` into memory. | **`config.go`**:<br>

// <br>`cfg.StoragePath` is now `"storage/storage.db"` |

// ---

// ## Summary of the special keywords in your struct tags:

// yaml:"env": First, Go checks your local.yaml file for a key named env. (You got this 100% right!)

// env:"ENV": If it doesn't find it in the YAML file, Go checks your computer's operating system for an Environment Variable named ENV.

// env-required:"true": If it cannot find it in the YAML file and it cannot find it in the computer's environment variables, the program will crash and refuse to start.

// godotenv (The Loader): It reads your local, physical .env file (which contains key-value pairs like ENV=local), parses them, and injects/loads them directly into your computer's active operating system environment variables.

// cleanenv (The Reader/Parser): It does not know or care about your .env file. Instead, it looks at your configuration struct tags (like env:"ENV" or env-required:"true") and fetches those values directly from the operating system, mapping them into your Go code.

package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production"

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// used to define and capture command-line flags (arguments) passed to your executable when you run it from the terminal.
// 		"config" (The Name): This is the name of the flag. In the terminal, the user will trigger it using -config or --config.

// "" (The Default Value): If the user runs the application without providing this flag, Go will default this value to an empty string "".

// "path to the configuration file" (The Help Message): This is the documentation. If someone runs your application with the -help or -h flag, Go will automatically print this message to explain what the flag does.

		flags := flag.String("config", "", "path to the configuration file")

		// Defining the flag isn't enough. Go needs an explicit instruction to actually scan the terminal inputs. You must call:
		flag.Parse()
		// Without flag.Parse(), Go will completely ignore what you typed in the terminal, and *configPath will remain empty (its default value).


		// configPath is a pointer pointing to the memory address where the terminal input will eventually be stored. To get the actual text value later in your code, you must dereference it using an asterisk (*configPath).
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	// os.Stat(configPath) checks if a file actually exists at that location before reading.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	// cleanenv.ReadConfig(...) reads the YAML file, decodes it, performs the env-required validation checks, and fills the cfg memory address with the parsed settings.
	// we are passing cfg be reference so that parsed information is filled in the same variable and not in the copy of it.
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	// Finally, it returns a pointer containing all loaded configurations. So that values could be extracted from the same mermory addres.
	return &cfg
}
