# Pokedex CLI

A command-line Pokedex application written in Go. This CLI interacts with [PokéAPI](https://pokeapi.co/) to let you:

- Explore location areas and see which Pokémon appear there.
- Catch Pokémon (with a bit of luck!).
- Inspect caught Pokémon to see their stats, types, height, and weight.
- Maintain an in-memory Pokédex of all Pokémon you’ve caught.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Commands](#commands)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

This CLI allows you to interact with PokéAPI’s endpoints to list locations, explore them, catch Pokémon, and inspect the details of any Pokémon you’ve caught — all from your terminal.

You start by running the program, which provides a REPL-style prompt `Pokedex >`. Type commands (for example, `help` or `catch pikachu`) and press Enter to execute them.

Under the hood, the application:
- Uses Go’s `net/http` package to make requests to the PokéAPI.
- Caches requests using a simple in-memory cache (defined in `internal/pokecache`).
- Stores caught Pokémon information in a local map (memory-based).

---

## Features

1. **View Location Areas**  
   Use `map` to view the next set of location areas, or `mapb` to go back (if available).
2. **Explore a Location**  
   Use `explore <location-area>` to see which Pokémon appear in that area.
3. **Catch a Pokémon**  
   Use `catch <pokemon-name>` to (attempt to) catch a Pokémon. Catch success is randomized based on the Pokémon’s base experience.
4. **Inspect a Pokémon**  
   Use `inspect <pokemon-name>` to display stats, types, height, and weight for a previously caught Pokémon.
5. **List All Caught Pokémon**  
   Use `pokedex` to see a list of all Pokémon you’ve caught so far.
6. **Help**  
   Use `help` to display details about all available commands.
7. **Exit**  
   Use `exit` to close the CLI.

---

## Requirements

- **Go** (1.18 or higher should work fine, though the `go.mod` suggests Go 1.23.6).
- Internet connectivity to query data from [PokéAPI](https://pokeapi.co).

---

## Installation

1. **Clone the repository** (or place these files in a directory):
   ```bash
   git clone https://github.com/mu7ammad1951/pokedex.git
   ```
   (Adjust the URL if your fork or remote is different.)

2. **Navigate to the project directory**:
   ```bash
   cd pokedex
   ```

3. **(Optional) Tidy up any dependencies**:
   ```bash
   go mod tidy
   ```

4. **Build the project**:
   ```bash
   go build -o pokedex
   ```
   This will produce an executable named `pokedex` (or `pokedex.exe` on Windows).

---

## Usage

Once you’ve built the project, run the executable:

```bash
./pokedex
```

You’ll see a prompt:
```
Pokedex >
```
Now you can enter commands (see the list below).  

### Commands

- **help**  
  Displays all available commands and their descriptions.

- **map**  
  Retrieves and displays a list of Pokémon location areas (next page).

- **mapb**  
  Retrieves and displays the previous page of Pokémon location areas (if available).

- **explore <location-area>**  
  Lists all Pokémon found in the specified location area.

- **catch <pokemon-name>**  
  Attempts to catch the specified Pokémon (using a random check based on their base experience).

- **inspect <pokemon-name>**  
  Shows detailed information about a Pokémon (stats, types, etc.) that you’ve previously caught.

- **pokedex**  
  Lists all of the Pokémon you have caught so far.

- **exit**  
  Exits the CLI.

#### Example REPL session

```
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Retrieves and displays Pokémon location areas from the next available page in the Pokédex API
mapb: Retrieves and displays Pokémon location areas from the previous available page in the Pokédex API.
explore: Lists all pokemon in a given location area. Usage: explore <location-area>
catch: Attempt to catch a pokemon and add it to your pokedex. Usage: catch <pokemon_name>
inspect: Inspect a Pokemon you've caught. Usage: inspect <pokemon_name>
pokedex: Show a list of all the Pokemon you've caught

Pokedex > map
...
Pokedex > explore route-1
Exploring route-1...
Found Pokemon:
 - pidove
 - rattata
 - ...
Pokedex > catch pidove
Throwing a Pokeball at pidove...
pidove was caught!

Pokedex > pokedex
Your Pokedex:
 - pidove
Pokedex > exit
Closing the Pokedex... Goodbye!
```

---

## Project Structure

```
.
├── command_catch.go
├── command_exit.go
├── command_explore.go
├── command_help.go
├── command_inspect.go
├── command_map.go
├── command_pokedex.go
├── go.mod
├── internal
│   ├── pokeapi
│   │   ├── client.go
│   │   ├── explore_location.go
│   │   ├── location_list.go
│   │   ├── pokeapi.go
│   │   ├── pokemon_info.go
│   │   ├── types_area.go
│   │   ├── types_locations.go
│   │   └── types_pokemons.go
│   └── pokecache
│       ├── cache_methods.go
│       ├── cache_test.go
│       └── types_cache.go
├── main.go
├── repl.go
├── repl_test.go
└── ...
```

- **main.go**  
  Entrypoint; sets up the config, initializes the `pokeapi.Client`, and starts the REPL.
- **repl.go**  
  Implements the CLI loop, reading user input and routing to commands.
- **command_*.go**  
  Each file defines a specific CLI command and its callback logic.
- **internal/pokeapi/**  
  Interacts with the [PokéAPI](https://pokeapi.co/), handling HTTP requests, unmarshaling JSON responses, etc.
- **internal/pokecache/**  
  Implements a thread-safe, in-memory cache with a reaper goroutine that periodically removes stale entries.
- **go.mod**  
  Go module file; identifies the module name and required Go version.

---

## Testing

You can run tests (both for the `pokecache` and REPL logic) using:

```bash
go test ./...
```

This will find and run all Go tests in the current module’s subdirectories.

---

## Contributing

Pull requests, bug reports, and feature suggestions are welcome!  

1. **Fork** the repo.
2. Create a new branch for your feature/bugfix: `git checkout -b feature/some-improvement`.
3. **Commit** your changes: `git commit -m 'Add some improvement'`.
4. **Push** to your fork: `git push origin feature/some-improvement`.
5. Open a **Pull Request** against this repo.

---



**Enjoy your command-line Pokémon adventure!**