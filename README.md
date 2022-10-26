# Proxx game
## Description
Proxx game simulation. Just to demonstrate Golang skills. Contains an implementation of the Proxx game algorithm and also includes a console representation of the game.
## Configuration
Configuration is done through the environment variables. Supported configuration parameters:  
- ```BOARD_WIDTH``` - determines game board width.
- ```BOARD_WIDTH_MAX``` - determines game board MAX width that could be entered in console.
- ```BOARD_HEIGHT``` - determines game board height.
- ```BOARD_HEIGHT_MAX``` - determines game board MAX height that could be entered in console.
- ```BOARD_BLACK_HOLES_COUNT``` - determines number of black holes on the gaming board.
- ```LOG_LEVEL``` - determines the logs level of the [logrus](https://github.com/sirupsen/logrus) logger. Could have values such as  ```panic```, ```fatal```, ```error```,  ```warn```, ```info```, ```debug```, ```trace```. By default it is set to ```debug```.
## Usage
### Command line
Direct usage could be done with the following command:  

```go run main.go --env ./config/.env.dev```  

```--env``` (shot version is ```-e```) parameter should take the path to the config file.
### Makefile
Use makefile for a quick run of the program. There are the following rules:
- ```run``` - runs the game with the DEV configuration file.  
- ```test``` - runs unit-tests.
- ```godoc``` - generates docs and launches it on http://localhost:6060.  
- ```lint``` - runs linting.  