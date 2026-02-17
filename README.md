# warpget
warpget is a high-performance concurrent download manager written in Go.

Think wget on steroids: parallel chunked downloads, resume support, smart retries, and modern concurrency primitives.

Built to be fast, reliable, and developer-grade.

---
## Why warpget?

Traditional download tools are largely sequential and blocking.  
warpget is designed to:

- Maximize network throughput
- Handle large files efficiently
- Scale across multiple downloads without blocking
- Recover gracefully from failures

In short: bytes move faster, CPU stays busy, users stay productive.

---
## Features

- Parallel chunked downloads
- Pause & resume support
- Real-time progress bars (speed, ETA, completion)
- Automatic retries with exponential backoff
- Multiple concurrent downloads
- Configurable concurrency & rate limits
- Smart HTTP handling (Range, ETag, Last-Modified)

**Planned:**

- Checksum verification (SHA256)
- Proxy support
- Website mirroring mode
- Config file support
- Plugin-friendly architecture

---
## Installation

### From source (recommended)

```bash
git clone https://github.com/yourusername/warpget.git
cd warpget
go build -o warpget ./cmd/warpget
```

The binary will be available at:

`./warpget`

(Alternatively, once published: `go install github.com/idorocodes/warpget/cmd/warpget@latest`)

---
## Usage

### Basic download

```bash
warpget https://example.com/file.zip
```

### Specify output file

```bash
warpget https://example.com/file.zip -o file.zip
```

### Set chunk concurrency

```bash
warpget https://example.com/file.zip -c 8
```

### Download multiple files

```bash
warpget https://a.com/a.zip https://b.com/b.zip
```

### Resume an interrupted download

```bash
warpget https://example.com/file.zip --resume
```

### Rate limiting (e.g. 2 MiB/s)

```bash
warpget https://example.com/file.zip --limit 2M
```

---
## How It Works (High-Level)

- File size is discovered via HEAD request (or first GET)
- File is split into byte ranges (equal-sized chunks)
- Each range is downloaded by a goroutine using `http.Client` with `Range` header
- Downloads are coordinated using a `sync.WaitGroup` + buffered channels
- Progress is reported in real time via terminal UI (using a library such as `github.com/vbauerster/mpb` or `github.com/cheggaaa/pb/v3`)
- Partial state (completed ranges, ETag, etc.) is persisted in a small `.warpget` metadata file for resume support

**Mental model:**

> URLs are jobs → chunks are goroutines → HTTP range streams → concurrent file writes (using `os.File` + `io.WriteSeeker`)

---
## Tech Stack

- Go (1.21+)
- `net/http` / `github.com/go-resty/resty/v2` (recommended for cleaner retry & range handling)
- `github.com/spf13/cobra` or `github.com/urfave/cli/v2` (CLI interface)
- `github.com/vbauerster/mpb/v8` or `github.com/cheggaaa/pb/v3` (progress bars)
- `golang.org/x/time/rate` (for rate limiting)
- `sync`, `context`, channels, `io`, `os` (core concurrency & I/O)

---
## Project Status

🚧 Active development  
Core concurrent download engine is the current focus.  
Stability, performance, and correctness take priority over features.

---
## Contributing

Contributions are welcome.  

Good areas to contribute:

- Robust error handling & retry strategies
- Performance optimization (connection pooling, HTTP/2 prioritization)
- Cross-platform testing (Windows resume paths, symlink handling)
- Documentation & usage examples

Please open an issue before starting major changes.
