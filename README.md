
# perlfunc

Implementation of some perl-like list and hash
functions in go using go1.18 generics.

It uses some experimental packages:

* golang.org/x/exp/constraints
* golang.org/x/exp/maps
* golang.org/x/exp/slices


## lists

contains perl list functions usable with all kinds of
go slices and a type 'List' that has these functions
as methods.


## hashes

contains some perl hash functions and grep and map 
for go maps.


## assert

contains generic test functions.

