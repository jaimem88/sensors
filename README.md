# Sensors

Solution to SafetyCulture's coding challenge

## Problem description

We are collecting temperature data about fridges in a supermarket. Imagine we have data from.different fridge sensors aggregated into a single JSON array (where an individual sensor is identified by an id):

### Input

```json
[{"id": "a","timestamp": 1509493641,"temperature": 3.53},
{"id": "b","timestamp": 1509493642,"temperature": 4.13},
{"id": "c","timestamp": 1509493643,"temperature": 3.96},
{"id": "a","timestamp": 1509493644,"temperature": 3.63},
{"id": "c","timestamp": 1509493645,"temperature": 3.96},
{"id": "a","timestamp": 1509493645,"temperature": 4.63},
{"id": "a","timestamp": 1509493646,"temperature": 3.53},
{"id": "b","timestamp": 1509493647,"temperature": 4.15},
{"id": "c","timestamp": 1509493655,"temperature": 3.95},
{"id": "a","timestamp": 1509493677,"temperature": 3.66},
{"id": "b","timestamp": 1510113646,"temperature": 4.15},
{"id": "c","timestamp": 1510127886,"temperature": 3.36},
{"id": "c","timestamp": 1510127892,"temperature": 3.36},
{"id": "a","timestamp": 1510128112,"temperature": 3.67},
{"id": "b","timestamp": 1510128115,"temperature": 3.88}]
```

### Output

```json
[
 {
  "id": "b",
  "average": 4.08,
  "median": 4.15,
  "mode": [
   4.15
  ]
 },
 {
  "id": "c",
  "average": 3.72,
  "median": 3.95,
  "mode": [
   3.36,
   3.96
  ]
 },
 {
  "id": "a",
  "average": 3.78,
  "median": 3.66,
  "mode": [
   3.53
  ]
 }
]
```

## How to run

From a terminal run with optional parameters. If no parameters are given then the program assumes the file `input.json` to be present in the same directory.

```sh
go run ./cmd/cli.sensor.data/main.go -input input.json -output output.json
```

## Tests

```sh
go test
```

## Build

Building this file will generate a `cli.sensor.data` binary file

```sh
go build ./cmd/cli.sensor.data/
```

## Assumptions

1. I decided to use `float64` to represent the temperature. The result may vary depending on where the program runs due to actual representation of floating point numbers. For a more accurate representation I would have used the [decimal package](https://github.com/shopspring/decimal)

2. I did not implement my own Sort nor Round functions

3. Reading a file into memory could be dangerous if it's too big and can cause a memory issue depending on the instance where it runs

4. The `mode` function is not the most optimal implementation as various iterations over different arrays isn't the best approach. There is probably a better or more complicated mathematical solution to obtain that value