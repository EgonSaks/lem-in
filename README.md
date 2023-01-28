## lem-in

### Objectives

This project is meant to make you code a digital version of an ant farm.

Finding the quickest path, `lem-in` will display the content of the file passed as argument and each move the ants make from room to room.

How does it work?

- Program makes an ant farm with tunnels and rooms.
- Places the ants on one side and and passes ants through so they find the exit.

Quickest way is used to get `n` ants across a colony (composed of rooms and tunnels).

- At the beginning of the game, all the ants are in the room `##start`. The goal is to bring them to the room `##end` with as few moves as possible.
- The shortest path is not necessarily the simplest.
- Some colonies will have many rooms and many links, but no path between `##start` and `##end`.
- Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no `##start` or `##end`, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message `ERROR: invalid data format`. If you wish, you can elaborate a more specific error message (example: `ERROR: invalid data format, invalid number of Ants` or `ERROR: invalid data format, no start room found`).

Results is shown on the standard output in the following format:

```console
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

- x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

- A room is defined by `"name coord_x coord_y"`, and will usually look like `"Room 1 2", "nameoftheroom 1 6", "4 6 7"`.

- The links are defined by `"name1-name2"` and will usually look like `"1-2", "2-5"`.

Here is an example of this in practice:

```
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```

Which corresponds to the following representation :

```console
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

### Usage



```
❯ go run main.go example00.txt
```
output:

```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

L1-2 
L1-3 L2-2 
L1-1 L2-3 L3-2 
L2-1 L3-3 L4-2 
L3-1 L4-3 
L4-1 

Found 1 paths in 34.875µs.
Used quickest path possible with 6 turns
```
or 

```
❯ sh test.sh
```

Learned about:

- Algorithms: Depth-first search (DFS), Breadth-first search (BFS), Dijkstra's algorithm, Bubble sort algorithm, greedy algorithm
- Ways to receive data
- Ways to output data
- Manipulation of strings
- Manipulation of structures

### Author

Egon Saks