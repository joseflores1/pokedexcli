# Pokedex CLI

A robust and lightweight CLI application designed for interacting with the PokeAPI and managing a local Pokedex.

## Description

This application is a full-featured tool that provides a seamless workflow for users to explore the Pokemon world from their terminal. It can list location areas, explore the Pokemon within them, catch Pokemon, and inspect their details. It leverages the PokeAPI and implements a custom caching mechanism to ensure fast performance and reduced API load.

## Motivation

Interacting with the PokeAPI through a browser or raw curl commands can be tedious. This project was developed to provide a minimalist yet powerful alternative by:

*   **Ensuring Simplicity**: Leveraging a clean Go structure for maximum performance and portability.
*   **Streamlining Exploration**: Automating the process of discovering location areas and their inhabitants.
*   **Game-like Experience**: Implementing a catching mechanic and a personal Pokedex to track progress.
*   **Optimized Performance**: Implementing a background-reaping cache to keep the application responsive and respect API rate limits.

## Quick Start

### Prerequisites

*   Go 1.23+

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/joseflores1/pokedexcli
    cd pokedexcli
    ```

2.  Install dependencies:
    ```bash
    go mod download
    ```

3.  Run the application:
    ```bash
    go run .
    ```

## Usage

### Project Structure

*   `internal/pokeapi/`: Contains the client logic for interacting with the PokeAPI.
*   `internal/pokecache/`: Implements a thread-safe caching mechanism with automatic expiration.
*   `main.go`: The entry point for the application.
*   `repl.go`: The orchestration layer for the interactive REPL and command routing.
*   `command_*.go`: Individual implementations for each CLI command.

### Commands

*   `help`: Displays a help message with available commands.
*   `map`: Displays the next 20 location areas.
*   `mapb`: Displays the previous 20 location areas.
*   `explore <location_name>`: Lists all Pokemon found in a specific location area.
*   `catch <pokemon_name>`: Attempts to catch a Pokemon. Success depends on its base experience.
*   `inspect <pokemon_name>`: Shows stats and details of a caught Pokemon.
*   `pokedex`: Lists all Pokemon currently in your Pokedex.
*   `exit`: Exits the application.

## Key Features

*   **Interactive REPL**: A user-friendly shell-like interface for executing commands.
*   **Automated Caching**: Efficiently caches API responses with a background cleaner to minimize network requests.
*   **Catching Mechanic**: A probability-based system for catching Pokemon, adding a layer of gamification.
*   **Persistent Session**: Keeps track of caught Pokemon throughout the session.

## Architecture

The application is organized into modular components:

*   **REPL Layer**: Handles user input, command parsing, and dispatching in `repl.go`.
*   **Command Layer**: Discrete functions in `command_*.go` implementing specific logic for each command.
*   **API Client**: A dedicated package in `internal/pokeapi/` for PokeAPI interaction, handling HTTP requests and JSON decoding.
*   **Cache System**: A thread-safe, time-based cache in `internal/pokecache/` that automatically removes stale entries.
