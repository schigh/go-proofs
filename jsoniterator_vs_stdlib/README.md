# jsoniterator vs stdlib

This is a head-to-head comparison between jsoniterator (github.com/json-iterator/go) 
and the standard library marshaling structs with userland marshal overrides.  I'm 
using the aliasing pattern inside the marshaller to manipulate empty data output

## unscientific result

When providing custom marshalling behavior, the stdlib tends to be faster and 
allocates less memory, even when using jsoniterator's fastest config

Run `./bench.sh` to run the simple benchmarks.  Output will go to results.txt
