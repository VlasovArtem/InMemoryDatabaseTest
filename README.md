<!-- TOC -->

* [Information](#information)
* [Run](#run)
* [Description](#description)
    * [Question](#question)
    * [Task](#task)
    * [Examples](#examples)
        * [Example 1 for commit a transaction](#example-1-for-commit-a-transaction)
        * [Example 2 for roll_back()](#example-2-for-roll_back)
        * [Example 3 for nested transactions](#example-3-for-nested-transactions)
        * [Example 4 for nested transactions with roll_back()](#example-4-for-nested-transactions-with-roll_back)

<!-- TOC -->

# Information

The implementation assumes that there is no multi-thread access to the database.  

The implementation is based on snapshot isolation.  At the start of each transaction, the system creates a new transaction with a new snapshot of the database.  

If a user commits changes, then the changes are applied from the ongoing transaction.  

If a user aborts changes, then the changes of an ongoing transaction are discarded.

* [Main](./cmd/main.go)
* [Database Interface](./internal/domain/inmemory_database.go)
* [Database Implementation](./internal/repository/inmemory_database.go)
* [Database Tests](./internal/repository/inmemory_database_test.go)

# Run

```shell
go run cmd/main.go
```

# Description

Please include a working coded solution along with an explanation for choosing a certain approach.

## Question

You are tasked with designing and implementing an in-memory database that supports key-value storage and transactions,
including operations like get(key), set(key, value), delete(key), start_transaction(), commit() and roll_back().
A user is able to start a transaction with start_transaction(). Then the user could use commit() to commit all changes
in the current transaction, or use roll_back() to discard all changes in the current transaction.
In addition, the database supports nested transactions, which means you are able to create a transaction within a
transaction.
In this question, you can assume there is no multithread access.
Please refer to the examples section to get a better idea how the system works.

## Task

Your task is to implement the InMemoryDatabase Class that supports the above database's core functionality. For example,
the class has an interface defined as below in Python. You should use GOLANG.

```python
class InMemoryDatabase:
    def __init__(self):

    def get(self, key):
        """
        Get the value associated with the given key.
        :param key: The key to retrieve.
        :return: The value associated with the key or None if the key does not exist.
        """

    def set(self, key, value):
        """
        Store a key-value pair in the database.
        :param key: The key to store.
        :param value: The value to associate with the key.
        :return: None
        """

    def delete(self, key):
        """
        Delete the key-value pair associated with the given key.
        :param key: The key to delete.
        :return: None
        """

    def start_transaction(self):
        """
        Start a new transaction. All operations within this transaction are isolated from others.
        :return: None
        """

    def commit(self):
        """
        Commit all changes made within the current transaction to the database.
        :return: None
        """

    def rollback(self):
        """
        Roll back all changes made within the current transaction and discard them.
        :return: None
        """
```

## Examples

### Example 1 for commit a transaction

```
db = InMemoryDatabase()
db.set("key1", "value1")
db.start_transaction()
db.set("key1", "value2")
db.commit()
db.get(“key1”)    -> Expect to get “value2”
```

### Example 2 for roll_back()

```
db = InMemoryDatabase()
db.set("key1", "value1")
db.start_transaction()
db.get("key1")    -> Expect to get “value1”
db.set("key1", "value2")
db.get("key1")    -> Expect to get ”value2”
db.roll_back()
db.get(“key1”)    -> Expect to get “value1”
```

### Example 3 for nested transactions

```
db = InMemoryDatabase()
db.set("key1", "value1")
db.start_transaction()
db.set("key1", "value2")
db.get("key1")    -> Expect to get ”value2”
db.start_transaction()
db.get("key1")    -> Expect to get ”value2”
db.delete(“key1“)
db.commit()
db.get(“key1”)    -> Expect to get None
db.commit()
db.get(“key1”)     -> Expect to get None
```

### Example 4 for nested transactions with roll_back()

```
db = InMemoryDatabase()
db.set("key1", "value1")
db.start_transaction()
db.set("key1", "value2")
db.get("key1")    -> Expect to get ”value2”
db.start_transaction()
db.get("key1")    -> Expect to get ”value2”
db.delete(“key1“)
db.roll_back()
db.get(“key1”)    -> Expect to get “value2”
db.commit()
db.get(“key1”)     -> Expect to get “value2”
```