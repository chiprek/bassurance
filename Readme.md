# Bassurance

Bassurance is a my second personal project For Boot.dev
## I am still working on this but i have taken way to much time to complete this to a version I am happy with releasing. If you have stumbled upon this do note I am still working on this. not that I feel anyone wants this oddly specific tooling in their production line. 

## What does it do:
Bassurance is an restAPI for a having an audit log for factory floor workers, it tracks the progress of a unit in its creation on the factory floor.
this entails creation of new jobs , units, attaching units to many diffrent jobs e.g: inital creation job, and a warrnty call later in the units life span. uploading proof of work and pictures of inprogress and completed products with floor notes (notes from the asembly).

# Setup
things you will need:
- a postgres database and the network address for where the database lives.
- a .env file with such variables:
  - `DB_URL="<network location of running postgres database>"`
  - `PLATFORM="<either dev or prod>"`
# Bassurance cli

## Installation:
To install the cli tool run
```bash
go install github.com/chiprek/bassurance/cmd/bassurance@latest
go install github.com/chiprek/bassurance/cmd/api_server@latest
```

Using Goose run the available sql migrations via goose up 
example
```bash
  goose up <location of sql database> 
```

## Todo: 
- Finish this project.
- Incoperate a front end. 

## Operations currently available
- create and modify:
  - UNITS
  - Sub Asemblies
  - Jobs
