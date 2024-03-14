# Readme: Time Tracker CLI

This is a simple CLI program for tracking time spent on activities. It offers basic functionalities to start, pause, resume, stop, and show the recorded time.

## Installation

To use this program, you need to have Go installed on your system. Once you have Go set up, you can install the program with the following command:

```bash
go install github.com/Diegiwg/tt@latest
```

Make sure the `$GOPATH/bin` directory is added to your `$PATH` so you can run the `tt` command from anywhere on your system.

## Usage

### Available Commands

1. **start**: Starts a new time record.
   - **Description**: Clears the database and starts a new time record.
   - **Usage**: `tt start`

2. **pause**: Adds a pause to the time record.
   - **Description**: When you want to take a break, use this command.
   - **Usage**: `tt pause`

3. **resume**: Resumes counting time after a pause.
   - **Description**: When you want to resume after a pause, use this command.
   - **Usage**: `tt resume`

4. **stop**: Stops time counting, clearing the time record.
   - **Description**: When you want to finish the time record, use this command.
   - **Usage**: `tt stop`

5. **show**: Shows the time passed in the time record.
   - **Description**: When you want to see how much time has passed without finishing the record, use this command.
   - **Usage**: `tt show`

### Example Usage

To start a new time record:

```bash
tt start
```

To add a pause:

```bash
tt pause
```

To resume after a pause:

```bash
tt resume
```

To stop and clear the time record:

```bash
tt stop
```

To show the time passed:

```bash
tt show
```

## Contributing

This is an open-source project, and contributions are welcome! Feel free to fork this repository, implement improvements, and send a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
