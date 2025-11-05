# List Ranker

List ranker is a game to help you order a list of related items by preference. The inspiration behind this project was [this Pokemon ranker](https://fio4ri.github.io/FavoritePokemon/).

The main purpose of this project was to utilize the knowledge I have gained in the use of Go as a backend language.

## How to play

Beginning a game of List Ranker is easy. All you have to do is click one of the premade lists to start a game. You will presented with two options from a pool of items. You then click the item that you prefer. If you can't choose between the two, there is an option to skip and be presented with a new pair of options.

## Tech Stack

All logic for the game is written in Go, which is sent to the frontend using Gin. Axios is used to make API calls to the backend. I used [shadcn](https://ui.shadcn.com/) to speed up the process of building ui components.

## Future Features

[] More premade lists
<br>
[x] Dark Mode
<br>
[x] Improved ranking algorithm
<br>
[x] Custom Lists
<br>
[] Better results display
<br>
[] Save results to files
<br>
[] Analytics
