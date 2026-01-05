# warpget

warpget is a high-performance async download manager written in Rust.  
Think wget on steroids: parallel chunked downloads, resume support, smart retries, and modern async I/O.  

Built to be fast, reliable, and developer-grade.

---

## Why warpget?

Traditional download tools are largely sequential and blocking.  
warpget is designed around Rust async primitives and Tokio, allowing it to:  

- Maximize network throughput  
- Handle large files efficiently  
- Scale across multiple downloads without blocking  
- Recover gracefully from failures  

In short: bytes move faster, CPU stays busy, users stay productive.

---

## Features

-  Parallel chunked downloads  
-  Pause & resume support  
-  Real-time progress bars (speed, ETA, completion)  
-  Automatic retries with backoff  
-  Multiple concurrent downloads  
-  Configurable concurrency & rate limits  

-  Smart HTTP handling (Range, ETag, Last-Modified)  

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
cargo build --release
```

Binary will be available at:  
`target/release/warpget`

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

### Rate limiting

```bash
warpget https://example.com/file.zip --limit 2M
```

---

## How It Works (High-Level)

- Files are split into byte ranges  
- Each range is downloaded by an async task  
- Tasks are scheduled and coordinated using Tokio  
- Progress is streamed live to the terminal  
- Partial state is persisted for resume support  

**Mental model:**  

> URLs are jobs → chunks are async tasks → network streams → async file writes  

---

## Tech Stack

- Rust  
- Tokio (async runtime)  
- reqwest (async HTTP)  
- clap (CLI interface)  
- indicatif (progress bars)  

---

## Project Status

🚧 Active development  
Core async download engine is the current focus.  
Stability, performance, and correctness take priority over features.

---

## Contributing

Contributions are welcome.  

Good areas to contribute:  
- Error handling & retries  
- Performance optimization  
- Cross-platform testing  
- Documentation & examples  

Open an issue before major changes.

---

## License

MIT License


