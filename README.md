# Pokedex CLI
A simple CLI that allows you explore locations, catch pokemon, and keep track of all caught pokemon. It makes use of [PokeAPI](https://pokeapi.co/) to get data from the Pokemon universe.
# Build

```bash
go build -o pokedex
```
# Usage
```bash
./pokedex
```
Commands:
| Command | Description |
| --- | --- |
| help | Displays a help message. |
| exit | Exit the pokedex. |
| map | Displays the names of 20 location areas in the Pokémon world. |
| mapb | Displays the previous 20 location areas in the Pokémon world. |
| explore | Displays a list of pokemon that can be found in a given area. |
| catch | Attempts to catch a pokemon. If successful, the pokemon will be added to your pokedex. |
| inspect | Displays information about the pokemon if it has been caught. |
| pokedex | Displays a list of pokemon that have been caught. |
