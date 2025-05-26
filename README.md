# Go Update Tool

The Go Update Tool is a command-line application that allows users to update devices over the network. It continuously pings the specified devices in the background and performs firmware updates when necessary.

## Project Structure

```
go-update-tool
├── cmd
│   └── main.go        # Entry point of the application
├── internal
│   ├── updater
│   │   └── updater.go # Implementation of the update tool
│   └── ping
│       └── ping.go    # Ping functionality
├── go.mod             # Module definition and dependencies
└── README.md          # Project documentation
```

## Requirements

- Go 1.16 or higher
- The `huh` library for pinging devices

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/go-update-tool.git
   cd go-update-tool
   ```

2. Install the dependencies:
   ```
   go mod tidy
   ```

## Usage

To run the update tool, use the following command:

```
go run cmd/main.go <device-ip> <firmware-file>
```

### Example

```
go run cmd/main.go 192.168.1.100 firmware.bin
```

This command will start the update process for the device at the specified IP address using the provided firmware file.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.