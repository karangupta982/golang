package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type HttpServer struct{
	Addr string 
}

type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http-server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")

		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg) // &cfg memory location where data is to be stored

	
	if err != nil {
		log.Fatalf("Cant not read config file: %s", err.Error())
	}
	return &cfg
}





// Those backtick strings (e.g., yaml:"env" env:"ENV" env-required:"true") are called Struct Tags in Go.
// They are essentially metadata strings attached to struct fields. Go itself ignores them during normal execution, but outside libraries (like configuration parsers) read them using a mechanism called Reflection (reflect package) to understand how to map external data onto your struct. [1] 
// ------------------------------
// ## Why we write them in a Struct
// When you load a configuration from a file (like config.yaml) or from environment variables, Go needs to map those external text keys to your Go code properties.
// Without tags, a configuration library wouldn't know that an environment variable named ENV or a YAML key named storage_path belongs to your Env or StoragePath struct properties.
// Here is exactly what each part of your tags is doing:
// ## 1. yaml:"env" and yaml:"storage_path" [2] 
// This tells a YAML parsing library (like go-yaml or cleanenv) exactly what the key is named inside your config.yaml file. [3] 

// * In your YAML file, it expects to see lowercase keys like env: "production" and storage_path: "/data".
// * The tag maps env directly into the struct's Env field.

// ## 2. env:"ENV"
// This tells an environment variable library to look for a system environment variable explicitly named ENV (e.g., if you ran export ENV=production in your terminal). If found, it overrides whatever was in the YAML file. [4] 
// ## 3. env-required:"true"
// This is a custom validation rule used by specific config libraries (like ilyakaznacheev/cleanenv). It tells the library: "If this value is missing from both the YAML file and the environment variables, throw an immediate error and crash the application because we cannot run without it."
// ## 4. yaml:"http-server"
// This attaches to your HttpServer field. Because HttpServer doesn't have a data type listed next to it, it is an embedded struct (composition). This tag tells the YAML parser to look for a nested block in your YAML file called http-server, like this:

// env: "dev"storage_path: "./storage"http-server:
//   port: 8080
//   timeout: "4s"

// ------------------------------
// ## A Clear Visual Analogy
// Think of your Go struct fields as labeled storage boxes, and struct tags as barcodes sticky-noted to the outside of those boxes.

// * The Go compiler only cares about the box name (StoragePath).
// * The YAML delivery truck arrives with a package labeled storage_path.
// * The configuration library reads the barcode (yaml:"storage_path"), matches it to the truck's package, and places the contents inside your box.



// so you mean that the struct tags are just mapping based on the names present in yaml or env files, right? if a variable is present in both yaml and env then which would prioritised? clear me with complete picture
// Yes, you are exactly right about the mapping. Struct tags simply act as a bridge mapping external names to your Go properties. [1] 
// If a configuration variable exists in both your YAML file and as an Environment Variable, the Environment Variable always takes priority and overrides the YAML file. [2, 3] 
// Popular configuration packages (like [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv)) execute this logic in a specific, intentional sequence to give you a complete picture: [3, 4] 
// ------------------------------
// ## The Order of Priority (The Complete Picture)
// When your application boots up and calls a config-loading function, it overwrites data step-by-step from lowest priority to highest priority: [3] 

// [ Step 1: Default Values ]  --> (Lowest Priority)
//            ↓
// [ Step 2: YAML/JSON File ]  --> (Overwrites Defaults)
//            ↓
// [ Step 3: Environment Vars ] --> (Highest Priority - Overwrites YAML)

// ## Why is it designed this way?
// This hierarchy provides excellent flexibility across different environments: [5] 

// * 
// * Local Development: You rely on the local config.yaml file because it is easy to read and edit on your computer.
// * Production (Docker / Kubernetes): You use the exact same compiled Go binary, but you inject sensitive keys (like database passwords, API keys, or changing the ENV string) directly via Environment Variables. You don't need to rewrite or rebuild your YAML file for deployment. [6, 7, 8, 9, 10] 
// * 

// ------------------------------
// ## Code Execution Walkthrough
// Consider this struct layout:

// type Config struct {
//     Env  string `yaml:"env" env:"APP_ENV" env-default:"development"`
// }

// The parsing library handles the Env string through three distinct phases upon execution: [3, 11] 

//    1. Look for defaults: It reads env-default and fills the variable with "development".
//    2. Look at the YAML file: It opens your configuration file. If it finds env: "staging", it overwrites the default. The variable becomes "staging".
//    3. Look at the System Environment: It checks the operating system for a variable named APP_ENV. If it finds APP_ENV=production in the environment, it overwrites the YAML value. The final value becomes "production". [2, 3, 11, 12, 13] 

// If an option is marked as env-required:"true" and both the YAML file and the environment variables are completely blank, the library halts execution immediately and prints a missing property error. [14] 
// Would you like to see how to structurally group a database section or server section into nested structs using these tags? [15] 