package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codersgyan/students-api/internal/config"
	"github.com/codersgyan/students-api/internal/http/handlers/student"
	"github.com/codersgyan/students-api/internal/storage/sqlite"
)

func main() {
	// load config
	// config is a package identifier. When you write config.MustLoad(), you are instructing Go to look inside the imported config package toolbox and execute its exported MustLoad function. It works the exact same way as fmt.Println().
	cfg := config.MustLoad()
	// database setup

	// Create one SQLite instance for the entire application.
// main() owns the dependency and injects it wherever needed.
// Handlers never create the database themselves (Dependency Injection).
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	// Inject the Storage dependency into the handler.
// The handler depends on the Storage interface, not on SQLite.
// This allows SQLite, PostgreSQL, MongoDB or a fake implementation
// to be swapped without changing handler code.
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	// setup server

	// Addr: cfg.Addr tells the server where to listen (e.g., port :8080).Handler: router connects the Building (Server) to the Receptionist (Router). It tells the server: "Whenever someone connects, hand them over to this specific router to find the right endpoint."
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	// This line sets up a tripwire that catches shutdown signals sent by your operating system (OS) or when you press Ctrl+C in your terminal.
	// Without this line, your program would crash instantly when stopped, potentially cutting off a database save mid-sentence.
	// done: It acts like a mailbox. When the operating system triggers a shutdown signal, signal.Notify catches it and drops a message into this done mailbox.
	// os.Interrupt: A standard Go identifier for a generic process interruption. this represents the standard signal sent when a user requests a program to stop.
	// syscall.SIGINT: called as signal interrupt, this specifically listens for when you press Ctrl+C in your terminal while the server is running.
	// syscall.SIGTERM: short form for Signal Termination, this listens for external management systems telling your app to stop. For example, if you deploy this API to Docker, Kubernetes, or Render, those systems send a SIGTERM when they want to update or restart your application.
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// You press Ctrl+C, syscall.SIGINT fires, signal.Notify catches it, and throws it into the done channel, The <-done line finally receives the message, unblocks, and lets the code move down to safely run your server.Shutdown(ctx) function.


	// server.ListenAndServe(): This single function opens the network port, listens for incoming traffic, creates background threads (Goroutines) for users, and translates raw internet data streams into Go requests.go func() { ... }() : This runs the server inside a background worker. It prevents the server from blocking your program, allowing the code below it (the shutdown listener) to run at the same time.
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	// It tells your application: "You have exactly 5 seconds to finish cleaning up and shutting down, or else I will cut the power automatically."
	// context.Background() (The Starting Point): This creates an empty, blank context. Think of it as a blank sheet of paper. It serves as the base layer that you can add rules to (like a timeout).
	// context.WithTimeout(..., 5*time.Second) (The Timer Rule)
		// What it does: It takes that blank paper and writes a strict rule on it: "Expire in exactly 5 seconds."
		// Why it matters: In web development, you never want a shutdown process to hang forever. If a database connection freezes during shutdown, this 5-second limit ensures your application will still force-quit instead of getting stuck in limbo.
	// ctx (The Context Object): This is the actual context variable holding the 5-second countdown timer. You immediately pass it into your server on the very next line: server.Shutdown(ctx). The server looks at ctx, sees the 5-second rule, and begins tracking the time.
	// If your server finishes shutting down in 0.5 seconds, you don't want that 5-second timer continuing to tick in the computer's memory. Calling cancel() tells the operating system: "We finished early, you can destroy this timer now and free up system memory."
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// It tells the server to stop accepting new traffic, but gives existing users 5 seconds (time.Second * 5 from your context) to finish saving or fetching their student data before fully cutting the power.
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
