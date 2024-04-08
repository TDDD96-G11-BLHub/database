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
lib.FetchDocument(*client, "Sensordata", "deepoidsensor", document, lib.FnFindOne)
```

## InsertDocument
This function currently has two parameters for the documents to be inserted. One is for a single document and one is for multiple. When calling this function set the parameter not used to nil.

## DeleteDocument
When calling this function with FnDeletOne the filter MUST be set to the ObjectId of the document to be deleted. Otherwise a random matching document will be deleted. Also keep in mind that calling this function with FnDeleteMany will delete ALL matching documents so be careful!