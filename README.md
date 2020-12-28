Advent of Code 2020
===================

Solved up to day 10, part 1.  Calculating permutations defeated me.

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