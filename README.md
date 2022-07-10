# Cosmo
Cosmo is my personal distributed file system. The purpose of this file system is to use perisistent and in memory key value storage. The data storage paradigm I'm going to build off of is "Key Value Storage".  A lot of the functionality for this project will be mimicked from S3 to create my own simple storage unit. Similar to S3, we should be able to create buckets and store files in subprefixes or "Folders". The file system will have multiple endpoints for access. It can be used via API calls (webserver) or command line interface. Eventually, we'll want to add a distributed networking component and be able to sync replicas across multiple volumns. For now, we're focusing on creating a reliable key value storage pattern. 



## Key Value Methods

1. **Put**: Store objects, files and serialize to bytes. This is the equivalent of creating a file or folder for the system. Associate a "key" with whatever we're storing. 
2. **Get**: Retrive a file/contents of a folder using an associated key. 
3. **Delete**: Delete a value with a specified key. Keys should be unique


## Web Server
All the key value methods mentioned above should be callable using a rest api. There should be endpoints for each key value method. The user should be able to create,delete, copy keys using the API endpoints. Each API endpoint should return JSON responses that indicate the status of the method call. 
	* Create HTTP endpoints for PUT, GET, DELETE

## Replication
There should be a master server that is the "leader" of all the other volume servers. The data we store in the file system should be replicated to the volume servers. Whenever a get, put or delete happens, there should be remote calls to the volume servers that replicate those actions. For replication, we should implement a consensus algorithm in case the master server fails. A new master server should automatically be selected in the case of a failure. 
	* Implement RAFT consensus algorithm
	* Create a replicated log for all the volume/master servers
	* https://github.com/eliben/raft/blob/master/part1/server.go

1. **Raft**: Each server stores a log containing a series of commands, which it's state machines execute in order. Each log contains the same commands in the same order. Raft is supposed to keep this log consistent and replicated across all machines. 

## Rebalancing
If a volume server goes offline or, we want to change the amount of servers to distribute the data across, we'll need functionality to rebalance and distribute the data across the new amount of servers desired. For volume servers that fail or crash, their data needs to be backed up and replicated to the currently available volumes. 
