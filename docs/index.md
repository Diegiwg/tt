# Time Tracker CLI

[![Brazil Flag](https://diegiwg.github.io/tt/public/brazil.png) Documentation in Portuguese](https://diegiwg.github.io/tt/pt)

## Project Overview 🚀

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

```bash
tt start
# Some time has passed
tt pause
# Took a break
tt resume
# Worked some more
tt show
# See that in 10 minutes it's time to stop working
tt stop
# Ends today's work, and see the result on the screen of how many hours have passed
```

## Contributing

This is an open-source project, and contributions are welcome! Feel free to fork this repository, implement improvements, and send a pull request.

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
