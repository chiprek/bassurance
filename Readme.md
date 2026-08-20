# Bassurance

Bassurance is a my second personal project For Boot.dev
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
```
