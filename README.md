# chronarch

For development, run from the repo root with `CHRONARCH_DEV=1 go run ./cmd/chronarch`.
This uses `.chronarch/config.toml` and `.chronarch/chronarch.db` in the current
working directory (ignored by Git). Without `CHRONARCH_DEV=1`, config and data
use OS user directories. An explicit `storage.path` in the selected config
file overrides the default database path. The CLI commands are not yet backed
by storage; do not rely on them to record work hours.

`go-sqlite3` requires CGO and a C compiler. If macOS reports
`ld: library not found for -lresolv`, use
`CC=/usr/bin/clang SDKROOT="$(xcrun --show-sdk-path)" CHRONARCH_DEV=1 go run ./cmd/chronarch`.

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
│   │   └── session.go
│   │   └── id.go       # Interfaces owned by the domain/app layer
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
