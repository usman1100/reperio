# Reperio - Go Port Scanner

Reperio is a lightweight and efficient port scanner written in Go. It can concurrently scan multiple ports on a given host, leveraging Go's goroutines for fast and parallel scanning. Reperio provides quick feedback about the status of specified ports, whether they are open or closed.

## Features

- **Concurrent Scanning**: Scans multiple ports concurrently for faster results.
- **Customizable Port List**: Allows you to specify a list of ports to scan.
- **Timeout Handling**: Each port scan has a 5-second timeout for better control over slow or unresponsive hosts.
- **Colorized Output**: Uses color to differentiate between open and closed ports for better readability.

## Installation

To get started with Reperio, clone the repository and build the project:

```bash
git clone https://github.com/usman1100/reperio.git
cd reperio
go build
```

Alternatively, you can directly run it using:

```bash
go run main.go
```

## Usage

You can run Reperio using the following options:

```bash
reperio -host <hostname> -ports <ports>
```

- `-host`: The hostname or IP address of the target system (defaults to `localhost`).
- `-ports`: A comma-separated list of ports to scan (defaults to `80,443,22,8080`).

### Example Usage

1. **Scanning the default ports on localhost**:

   ```bash
   reperio
   ```

   This will scan ports 80, 443, 22, and 8080 on `localhost`.

2. **Scanning specific ports on a remote host**:

   ```bash
   reperio -host 192.168.1.1 -ports 21,22,80,443
   ```

   This command will scan ports 21, 22, 80, and 443 on the IP address `192.168.1.1`.

### Output

Reperio provides colorized output to clearly indicate the status of each port:

- **Green**: Port is open
- **Red**: Port is closed

Example output:

```bash
Scanning ports on localhost
Port 80    is open
Port 443   is closed
Port 22    is open
Port 8080  is closed
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---
