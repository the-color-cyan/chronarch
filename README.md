# chronarch

```text
chronarch/
├── cmd/
│   └── chronarch/
│       └── main.go             # Process entrypoint: config, DI, exit code
├── internal/                   # Private application code
│   ├── app/
│   │   └── service.go          # Use cases / orchestration
│   ├── domain/
│   │   ├── project.go          # Core entities and business rules
│   │   └── repository.go       # Interfaces owned by the domain/app layer
│   ├── config/
│   │   └── config.go
│   ├── storage/
│   │   ├── filesystem.go       # Repository implementation
│   │   └── sqlite.go
│   ├── cli/
│   │   ├── root.go             # Command setup
│   │   └── project.go          # CLI adapters / commands
│   ├── tui/                    # Add when needed
│   │   ├── model.go
│   │   └── views.go
│   └── gui/                    # Add only if/when needed
│       └── ...
├── pkg/                        # Optional: stable public Go libraries for others
│   └── api/
├── configs/                    # Example/default config, if applicable
├── scripts/                    # Build/release/dev scripts
├── testdata/                   # Fixtures consumed by tests
├── docs/
├── go.mod
├── go.sum
├── README.md
└── Makefile                    # Optional convenience targets
```
