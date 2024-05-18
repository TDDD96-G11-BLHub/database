# database
The database handler for BLHub.

# MongoDB
This is the usage of the db package

## How to run
Create a .env file in your root folder and set a uri variable like this:

```python
MONGODB_URI=<connection string goes here>
```

## Function usage
See the documentation for each function.
Configure the function calls with the DBFunction enum.

Ex:
```python
lib.FetchDocument(*client, "Sensordata", "deepoidsensor", filter, lib.FnFindOne)
```

# Json and Peer to Peer
This is the usage of the peerdb package

## How to run
To run this package u must create a local db with the CreateLocalDatabase() function.
This will create a new directory in the designated filaddress.
This directory can then be opened with LoadLocalDatabase if it already exists.
The directory is meant to be filled with json files that contain data, see the code for how
to create and load such files.

All other functions can then be run via the LocalDB object struct in the code.
The only supported sensortype is currently the deepoidsensor.

## Function usage
The two noteworthy structs are LocalDB and Collection, se the code for further info.
All LocalDB functions can be found in the LocalDBBase interface

Be careful with choosing file address and creating/removing directories as this package
has all the necessary permissions to edit the filesystem.

## Peer to Peer
The peer to peer functionality has yet to be implemented.
The package currently contains the functionality to set up a connection node.