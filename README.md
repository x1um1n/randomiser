# randomiser
For some bizarre reason, the mp3 player built into my car doesn't have a shuffle function, so I've built this little app to force a random sort order on a dir.

The app scans the supplied dir into a slice, then picks elements at random to be inserted into a second slice with an index number.  The user is prompted to confirm they approve of the new filenames, before all the files are renamed.

### usage
```
randomiser <relative/or/absolute/path/to/the/dir>
```

### build/install instructions
If you haven't already got golang installed, first follow this guide: https://golang.org/doc/install
```
##build:
go build randomiser.go

##install:
mv randomiser /usr/local/bin/ 
```
