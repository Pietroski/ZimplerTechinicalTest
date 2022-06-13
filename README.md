# Zimpler Technical Test

This is a technical test proposed at [zimpler's technical test page](https://candystore.zimpler.net/).

To run the script is simple, once you have go runtime installed, "zimply":
```bash
make run
```

After that, there will be printed two data structures, the first is the overall overview of the csv analysis and the second one is the expected output :)

It will look something like: 
```bash
[
  {
    "Name": "Jonas",
    "FavouriteSnack": "Geisha",
    "TotalSnacks": 1982
  },
  {
    "Name": "Annika",
    "FavouriteSnack": "Geisha",
    "TotalSnacks": 208
  },
  {
    "Name": "Jane",
    "FavouriteSnack": "Nötchoklad",
    "TotalSnacks": 22
  },
  {
    "Name": "Aadya",
    "FavouriteSnack": "Center",
    "TotalSnacks": 11
  }
]
```

## Final technical test considerations
There is still plenty of possibilities and room for improvements, for sure. However, as a script and a "simple" technical test I tried to keep things as "zimple" as possible :)

My main initial doubt was regarding whether I would store the given data into memory, csv ou database... so, to make things "zimple" but not very "zimple", I decided to go with the csv approach :)

The script has 100% test coverage.
To run the test suites
```bash
make test-unit-cover
```
or for verbose mode
```bash
make test-unit-cover-silent
```

For you to be able to see the coverage report, run
```bash
make test-unit-cover-report
```

Any doubts or willingness to explore, please, have a look in the Makefile :)

```bash
# Ps. -> I wish I had done a more robust architecture design but again... I tried to make things as "zimple" as possible.
```
