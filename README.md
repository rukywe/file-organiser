# File Organiser

File Organiser is a Go application that helps you organise files in a directory based on their extensions.

## Features

- Organises files into subdirectories based on their extensions
- Configurable via YAML file
- Logging with different log levels

## Prerequisites

- Go 1.16 or higher

## Installation

Clone the repository:

```sh
git clone https://github.com/rukywe/file-organiser.git
cd file-organiser
```

Install dependencies:

```sh
go mod tidy
```

## Configuration

Edit the `config.yaml` file to set your preferences:

```yaml
source_dir: '/path/to/source/directory'
log_level: 'info'
file_extensions:
  '.txt': 'documents'
  '.jpg': 'images'
  '.pdf': 'documents'
```

- `source_dir`: The directory to organize
- `log_level`: Logging level ("info" or "debug")
- `file_extensions`: Map of file extensions to target directories

## Usage

Build and run the application:

```sh
go build ./cmd/fileorganiser
./fileorganiser -config config.yaml
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
