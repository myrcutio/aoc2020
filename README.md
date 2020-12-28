Advent of Code 2020
===================

Solved up to day 10, part 1.  Calculating permutations defeated me.  
If you're going to run this, create an `input.csv` file in the relevant day and paste in your input.

Example usage:
```shell script
go run ./pkg/day1

# pair results of 1473 and 547 is 805731
# pair results of 438, 1222, and 360 is 192684960%
```

# Useful Links
https://adventofcode.com/2020

Private leaderboard:  
https://adventofcode.com/2020/leaderboard/private/view/1119121


# Boilr usage
Installing (docs https://github.com/tmrts/boilr/wiki/Installation)
```shell script
go get github.com/tmrts/boilr
```

Start up a brand new day (with csv input)
```shell script
make csvTemplate
make pkg/day42
```

Sanity check, confirm the template was created successfully
```shell script
go test ./pkg/day42
```