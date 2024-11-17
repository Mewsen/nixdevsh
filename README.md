# nixdevsh

A tool for simplifying and streamlining development workflows using Nix. This project is focused on automating environment setup, dependency management, and deployment processes for developers.

---

## Features
  - **Automated Development Environment**: Configure and provision consistent environments with ease.
  - **Nix Integration**: Leverages the power of Nix flakes for reproducibility and efficient dependency management.
  - **Interactive CLI**: A user-friendly Command-Line Interface (CLI) for interacting with the tool.
  - **Modular Design**: Clean separation of concerns across cmd, logic, and ui layers.

## Demo
![nixdevsh demo](demo.gif)

## Requirements
  - Nix installed on your system.[Nix Installation Guide](https://nixos.org/download.html)
  - Go (latest version recommended) for building and running the CLI.

## Installation
- Clone the repository:

```bash
git clone https://github.com/Mewsen/nixdevsh.git
cd nixdevsh
 ```
- Activate the Nix shell environment:

```bash
direnv allow .
```

## Usage
- Running the Tool
Run the main program using:

``` bash
go run main.go
```
- Explore the available commands using:

```bash
./nixdevsh --help
```

## Development
The project structure is as follows:

**cmd/**: Contains CLI-related logic and command definitions.**
logic/**: Core application logic and utilities.
**ui/**: User interface components for terminal interactions.

## Contributing
Contributions are welcome! Here's how you can help:
  
  - Fork the repository.
  - Create a new branch (git checkout -b feature/your-feature).
  - Make your changes and test them thoroughly.
  - Commit your changes (git commit -m "Add some feature").
  - Push to the branch (git push origin feature/your-feature).
  - Open a pull request.

## License
This project is licensed under the [MIT License](LICENSE)








