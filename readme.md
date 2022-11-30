
# Tic Tac Toe

In a game of Tic-Tac-Toe, two players take turns marking an available cell in a 3-by-3 grid with 
their respective tokens (either X or O). When one player has placed three tokens in a horizontal, 
vertical, or diagonal row on the grid, the game is over and that player has won. A draw (no winner) 
occurs when all the cells on the grid have been filled with tokens and neither player has achieved 
a win. More details on the game here : https://en.wikipedia.org/wiki/Tic-tac-toe

![Logo](./assets/tic_tac_toe.png)

## Requirements
Problem statement: 
  Given a string of ‘X’, ‘O’ and ‘-’ as input, 

  which represents a game board, print the state of the 

  game as output. The states must be one of:

  - X wins 

  - O wins

  - Draw 

  - Game in progress

  - Invalid grid

## Input Output
Examples :

1. X wins - when X occupies either all cells in a vertical line, or all cells in a horizontal line, or all 
cells in a diagonal.

    ```
    Sample Input :
    XOXXOOXXO
    ```

    Which represents the board:
    | X | O | X |
    |---|---|---|
    | X | O | O |
    | X | X | O |

    ```
    Sample output:
    X wins!
    ```

2) O wins - when O occupies either all cells in a vertical line, or all cells in a horizontal line, or all cells in a diagonal.

    ```
    Sample input:
    XOOXOXOXO
    ```

    Which represents the board:

    | X | O | O |
    |---|---|---|
    | X | O | X |
    | O | X | O |

    ```
    Sample output:
    O Wins!
    ```

3) Draw - no player has won.

    ```
    Sample input:
    OXOXOXXOX
    ```

    Which represents the board:

    | O | X | O |
    |---|---|---|
    | X | O | X |
    | X | O | X |

    ```
    Sample output:
    Its a draw!
    ```

4) Game in progress - if no player has won and its not a draw

    ```
    Sample input:
    XOXX--O--
    ```

    Which represents the board:

    | X | O | X |
    |---|---|---|
    | X |   |   |
    | O |   |   |

    ```
    Sample output:
    Game still in progress!
    ```

5) Invalid Grid - a grid thats not possible to achieve in a real game

    ```
    Sample Input:
    XXXOOOXXO
    ```

    Which represents the board:

    | X | X | X |
    |---|---|---|
    | O | O | O |
    | X | X | O |

    ```
    Sample Output:
    Invalid game board
    ```
