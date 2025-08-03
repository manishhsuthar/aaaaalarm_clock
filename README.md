# Alarm Clock Application

This is a simple alarm clock application built in Go. It allows users to set alarms, activate or deactivate them, and provides utility functions for time manipulation.

## Project Structure

```
alarm-clock-go
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── alarm
│   │   └── alarm.go     # Alarm management logic
│   └── utils
│       └── time.go      # Time utility functions
├── go.mod                # Module dependencies
└── README.md             # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd alarm-clock-go
   ```

2. Initialize the Go module:
   ```
   go mod tidy
   ```

3. Build the application:
   ```
   go build -o alarm-clock cmd/main.go
   ```

4. Run the application:
   ```
   ./alarm-clock
   ```

## Usage Examples

- To set an alarm, use the `SetTime` method from the `Alarm` struct.
- Activate the alarm using the `Activate` method.
- Deactivate the alarm with the `Deactivate` method.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.