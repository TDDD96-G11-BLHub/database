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
