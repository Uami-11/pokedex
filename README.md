# CLI Pokedex

This is the cli pokedex I made using go, it uses http requests from the [PokeAPI](https://pokeapi.co/) to get locations and pokemon information.

<img width="633" height="521" alt="image" src="https://github.com/user-attachments/assets/bcd40dcf-b08f-42ae-812e-c3d7c454ac81" />

## Installation
Clone the repository and do 
```
go build 
./pokedex
```

to build it and use it. You will need to have go tho.

## Features

### Exploring the map

I have two commands to go back and forth through the [PokeAPI](https://pokeapi.co/) map. map and mapb. It pages through the locations

<img width="712" height="718" alt="image" src="https://github.com/user-attachments/assets/dbc46e27-5125-4234-971d-38f5cc6f262c" />

And to get more details on a specific location. You can use the explore command on a specific location

<img width="813" height="830" alt="image" src="https://github.com/user-attachments/assets/405da815-00ce-47eb-85b6-c54724c6a1ba" />

### Catching and Looking at Pokemon

Then you can use the catch command to catch a specific pokemon you would like. The catch rate is based on the pokemon's base experience, so tougher 'mons will be harder to catch.

<img width="650" height="553" alt="image" src="https://github.com/user-attachments/assets/658f40c6-364b-40e0-839e-7b9a07af3ca6" />

You can check all your caught pokemon through the pokedex command, and then you can check for a specific pokemon's info using the inspect command.

<img width="636" height="725" alt="image" src="https://github.com/user-attachments/assets/2cd1acee-8f8b-4e5b-8f40-3baa81f6c05c" />


