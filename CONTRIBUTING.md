# Contributing

This process should be as easy as possible, because no one likes spending inordinate 
amounts of time making example code.

## New Proofs

To submit a new proof, create a subject off the main project.  Name the subject 
something pithy but also descriptive.

```
project root
└── subject
    ├── README.md
    ├── bench.sh        # use this for benchmarks
    ├── common
    │   └── common.go
    ├── run.sh          # use this to run your main
    ├── test.sh         # use this to run tests
    ├── use_case_one
    │   └── main.go
    └── use_case_two
        └── main.go
```

Each subject should contain a README that explains what the assumption is, how to 
test it, and a short TL;DR on the outcome.
