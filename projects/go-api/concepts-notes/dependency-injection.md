Dependency Injection (DI)

Problem:
- If handlers create SQLite themselves, they become tightly coupled,
  hard to test, and hard to replace.

Solution:
- main() creates dependencies.
- main() passes (injects) them into handlers.
- Handlers depend on interfaces, not concrete implementations.

Flow:
main()
   ↓
sqlite.New()
   ↓
Storage interface
   ↓
student.New(storage)
   ↓
storage.CreateStudent()
   ↓
Go dispatches to the concrete type (*Sqlite, *Postgres, etc.)

Key Rule:
"Don't create your dependencies; receive them."