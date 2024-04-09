# database
The database handler for BLHub.

# How to run
Create a .env file in your root folder and set a uri variable like this:

```python
MONGODB_URI=<connection string goes here>
```

# Function usage
See the documentation for each function.
Configure the function calls with the DBFunction enum.

Ex:
```python
lib.FetchDocument(*client, "Sensordata", "deepoidsensor", filter, lib.FnFindOne)
```

## InsertDocument
This function currently has two parameters for the documents to be inserted. One is for a single document and one is for multiple. When calling this function set the parameter not used to nil.

## DeleteDocument
When calling this function with FnDeletOne the filter MUST be set to the ObjectId of the document to be deleted. Otherwise a random matching document will be deleted. Also keep in mind that calling this function with FnDeleteMany will delete ALL matching documents so be careful!

## Drop functions
Be VERY CAREFUL with this function since it is irreversible and will drop the entire database along with the collections or the collection with the documents. It will not crash if the database or collection name given is incorrect so double check the result in the cluster to see if it actually dropped a database or collection.

## GetAll functions
These functions are very usable for getting information about the cluster and the databases. They will be used for functions where the names of databases and collections are used so the user won't have to explicitly give the name of the queriable object.