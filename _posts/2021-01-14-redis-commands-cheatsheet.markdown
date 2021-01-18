---
layout: post
title:  "Redis Commands CheatSheet"
date:   2021-01-14 19:18:26 +0530
---

1. # Cluster

    | [__CLUSTER ADDSLOTS__ slot [slot ...]](https://redis.io/commands/cluster-addslots) | __*O(N)*__ where N is the total number of hash slot arguments | Assign new hash slots to receiving node |
    | [__CLUSTER BUMPEPOCH__](https://redis.io/commands/cluster-bumpepoch) | __*O(1)*__ | Advance the cluster config epoch |
    | [__CLUSTER COUNT-FAILURE-REPORTS__ node-id](https://redis.io/commands/cluster-count-failure-reports) | __*O(N)*__ where N is the number of failure reports | Return the number of failure reports active for a given node |
    | [__CLUSTER COUNTKEYSINSLOT__ slot](https://redis.io/commands/cluster-countkeysinslot) | __*O(1)*__ | Return the number of local keys in the specified hash slot |
    | [__CLUSTER DELSLOTS__ slot [slot ...]](https://redis.io/commands/cluster-delslots) | __*O(N)*__ where N is the total number of hash slot arguments | Set hash slots as unbound in receiving node |
    | [__CLUSTER FAILOVER__ [FORCE/TAKEOVER]](https://redis.io/commands/cluster-failover) | __*O(1)*__ | Forces a replica to perform a manual failover of its master |
    | [__CLUSTER FLUSHSLOTS__](https://redis.io/commands/cluster-flushslots) | __*O(1)*__ | Delete a node's own slots information |
    | [__CLUSTER FORGET__ node-id](https://redis.io/commands/cluster-forget) | __*O(1)*__ | Remove a node from the nodes table |
    | [__CLUSTER GETKEYSINSLOT__ slot count](https://redis.io/commands/cluster-getkeysinslot) | __*O(log(N))*__ where N is the number of requested keys | Return local key names in the specified hash slot |
    | [__CLUSTER INFO__](https://redis.io/commands/cluster-info) | __*O(1)*__ | Provides info about Redis Cluster node state |
    | [__CLUSTER KEYSLOT__ key](https://redis.io/commands/cluster-keyslot) | __*O(N)*__ where N is the number of bytes in the key | Returns the hash slot of the specified key |
    | [__CLUSTER MEET__ ip port](https://redis.io/commands/cluster-meet) | __*O(1)*__ | Force a node cluster to handshake with another node |
    | [__CLUSTER MYID__](https://redis.io/commands/cluster-myid) | __*O(1)*__ | Return the node id |
    | [__CLUSTER NODES__](https://redis.io/commands/cluster-nodes) | __*O(N)*__ where N is the total number of Cluster nodes | Get Cluster config for the node |
    | [__CLUSTER REPLICATE__ node-id](https://redis.io/commands/cluster-replicate) | __*O(1)*__ | Reconfigure a node as a replica of the specified master node |
    | [__CLUSTER RESET__ [HARD\|SOFT]](https://redis.io/commands/cluster-reset) | __*O(N)*__ where N is the number of known nodes. The command may execute a FLUSHALL as a side effect | Reset a Redis Cluster node |
    | [__CLUSTER SAVECONFIG__](https://redis.io/commands/cluster-saveconfig) | __*O(1*__) | Forces the node to save cluster state on disk |
    | [__CLUSTER SET-CONFIG-EPOCH__ config-epoch](https://redis.io/commands/cluster-set-config-epoch) | __*O(1)*__ | Set the configuration epoch in a new node |
    | [__CLUSTER SETSLOT__ slot IMPORTING\|MIGRATING\|STABLE\|NODE [node-id]](https://redis.io/commands/cluster-setslot) | __*O(1)*__ | Bind a hash slot to a specific node |
    | [__CLUSTER SLAVES__ node-id](https://redis.io/commands/cluster-slaves) | __*O(1)*__ | List replica nodes of the specified master node |
    | [__CLUSTER REPLICAS__ node-id](https://redis.io/commands/cluster-replicas) | __*O(1)*__ | List replica nodes of the specified master node |
    | [__CLUSTER SLOTS__](https://redis.io/commands/cluster-slots) | __*O(N)*__ where N is the total number of Cluster nodes | Get array of Cluster slot to node mappings |
    | [__READONLY__](https://redis.io/commands/readonly) | __*O(1)*__ | Enables read queries for a connection to a cluster replica node |
    | [__READWRITE__](https://redis.io/commands/readwrite) | __*O(1)*__ | Disables read queries for a connection to a cluster replica node |

1. # Connection

    | [__AUTH__ [username] password](https://redis.io/commands/auth) | - | Authenticate to the server |
    | [__CLIENT CACHING__ YES\|NO](https://redis.io/commands/client-caching) | __*O(1)*__ | Instruct the server about tracking or not keys in the next request |
    | [__CLIENT ID__](https://redis.io/commands/client-id) | __*O(1)*__ | Returns the client ID for the current connection |
    | [__CLIENT INFO__](https://redis.io/commands/client-info) | __*O(1)*__ | Returns information about the current client connection |
    | [__CLIENT KILL__ [ip:port] [ID client-id] [TYPE normal\|master\|slave\|pubsub] [USER username] [ADDR ip:port] [SKIPME yes\|no]](https://redis.io/commands/client-kill) | __*O(N)*__ where N is the number of client connections | Kill the connection of a client |
    | [__CLIENT LIST__ [TYPE normal\|master\|replica\|pubsub] [ID client-id [client-id ...]]](https://redis.io/commands/client-list) | __*O(N)*__ where N is the number of client connections | Get the list of client connections |
    | [__CLIENT GETNAME__](https://redis.io/commands/client-getname) | __*O(1)*__ | Get the current connection name |
    | [__CLIENT GETREDIR__](https://redis.io/commands/client-getredir) | __*O(1)*__ | Get tracking notifications redirection client ID if any |
    | [__CLIENT UNPAUSE__](https://redis.io/commands/client-unpause) | __*O(N)*__ Where N is the number of paused clients | Resume processing of clients that were paused |
    | [__CLIENT PAUSE__ timeout [WRITE\|ALL]](https://redis.io/commands/client-pause) | __*O(1)*__ | Stop processing commands from clients for some time |
    | [__CLIENT REPLY__ ON\|OFF\|SKIP](https://redis.io/commands/client-reply) | __*O(1)*__ | Instruct the server whether to reply to commands |
    | [__CLIENT SETNAME__ connection-name](https://redis.io/commands/client-setname) | __*O(1)*__ | Set the current connection name |
    | [__CLIENT TRACKING__ ON\|OFF [REDIRECT client-id] [PREFIX prefix [PREFIX prefix ...]] [BCAST] [OPTIN] [OPTOUT] [NOLOOP]](https://redis.io/commands/client-tracking) | __*O(1)*__ | Enable or disable server assisted client side caching support |
    | [__CLIENT TRACKINGINFO__](https://redis.io/commands/client-trackinginfo) | __*O(1)*__ | Return information about server assisted client side caching for the current connection |
    | [__CLIENT UNBLOCK__ client-id [TIMEOUT\|ERROR]](https://redis.io/commands/client-unblock) | __*O(log N)*__ where N is the number of client connections | Unblock a client blocked in a blocking command from a different connection |
    | [__ECHO__ message](https://redis.io/commands/echo) | - | Echo the given string |
    | [__HELLO__ [protover [AUTH username password] [SETNAME clientname]]](https://redis.io/commands/hello) | __*O(1)*__ | Handshake with Redis |
    | [__PING__ [message]](https://redis.io/commands/ping) | - | Ping the server |
    | [__QUIT__](https://redis.io/commands/quit) | - | Close the connection |
    | [__RESET__](https://redis.io/commands/reset) | - | Reset the connection |
    | [__SELECT__ index](https://redis.io/commands/select) | - | Change the selected database for the current connection |

1. # Geo

    | [__GEOADD__ key [NX\|XX] [CH] longitude latitude member [longitude latitude member ...]](https://redis.io/commands/geoadd) | __*O(log(N))*__ for each item added, where N is the number of elements in the sorted set | Add one or more geospatial items in the geospatial index represented using a sorted set |
    | [__GEOHASH__ key member [member ...]](https://redis.io/commands/geohash) | __*O(log(N))*__ for each member requested, where N is the number of elements in the sorted set | Returns members of a geospatial index as standard geohash strings |
    | [__GEOPOS__ key member [member ...]](https://redis.io/commands/geopos) | __*O(log(N))*__ for each member requested, where N is the number of elements in the sorted set | Returns longitude and latitude of members of a geospatial index |
    | [__GEODIST__ key member1 member2 [m\|km\|ft\|mi]](https://redis.io/commands/geodist) | __*O(log(N))*__ | Returns the distance between two members of a geospatial index |
    | [__GEORADIUS__ key longitude latitude radius m\|km\|ft\|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC\|DESC] [STORE key] [STOREDIST key]](https://redis.io/commands/georadius) | __*O(N+log(M))*__ where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index | Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point |
    | [__GEORADIUSBYMEMBER__ key member radius m/\|km\|ft\|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC\|DESC] [STORE key] [STOREDIST key]](https://redis.io/commands/georadiusbymember) | __*O(N+log(M))*__*__ where N is the number of elements inside the bounding box of the circular area delimited by center and radius and M is the number of items inside the index | Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member |
    | [__GEOSEARCH__ key [FROMMEMBER member] [FROMLONLAT longitude latitude] [BYRADIUS radius m\|km/|ft\|mi] [BYBOX width height m\|km\|ft\|mi] [ASC\|DESC] [COUNT count [ANY]] [WITHCOORD] [WITHDIST] [WITHHASH]](https://redis.io/commands/geosearch) | __*O(N+log(M))*__ where N is the number of elements in the grid-aligned bounding box area around the shape provided as the filter and M is the number of items inside the shape | Query a sorted set representing a geospatial index to fetch members inside an area of a box or a circle |
    | [__GEOSEARCHSTORE__ destination source [FROMMEMBER member] [FROMLONLAT longitude latitude] [BYRADIUS radius m\|km\|ft\|mi] [BYBOX width height m\|km\|ft\|mi] [ASC\|DESC] [COUNT count [ANY]] [WITHCOORD] [WITHDIST] [WITHHASH] [STOREDIST]](https://redis.io/commands/geosearchstore) | __*O(N+log(M))*__ where N is the number of elements in the grid-aligned bounding box area around the shape provided as the filter and M is the number of items inside the shape | Query a sorted set representing a geospatial index to fetch members inside an area of a box or a circle, and store the result in another key |

1. # Hashes

    | [__HDEL__ key field [field ...]](https://redis.io/commands/hdel) | __*O(N)*__ where N is the number of fields to be removed | Delete one or more hash fields |
    | [__HEXISTS__ key field](https://redis.io/commands/hexists) | __*O(1)*__ | Determine if a hash field exists |
    | [__HGET__ key field](https://redis.io/commands/hget) | __*O(1)*__ | Get the value of a hash field |
    | [__HGETALL__ key](https://redis.io/commands/hgetall) | __*O(N)*__ where N is the size of the hash | Get all the fields and values in a hash |
    | [__HINCRBY__ key field increment](https://redis.io/commands/hincrby) | __*O(1)*__ | Increment the integer value of a hash field by the given number |
    | [__HINCRBYFLOAT__ key field increment](https://redis.io/commands/hincrbyfloat) | __*O(1)*__ | Increment the float value of a hash field by the given amount |
    | [__HKEYS__ key](https://redis.io/commands/hkeys) | __*O(N)*__ where N is the size of the hash | Get all the fields in a hash |
    | [__HLEN__ key](https://redis.io/commands/hlen) | __*O(1)*__ | Get the number of fields in a hash |
    | [__HMGET__ key field [field ...]](https://redis.io/commands/hmget) | __*O(N)*__ where N is the number of fields being requested | Get the values of all the given hash fields |
    | [__HMSET__ key field value [field value ...]](https://redis.io/commands/hmset) | __*O(N)*__ where N is the number of fields being set | Set multiple hash fields to multiple values |
    | [__HSET__ key field value [field value ...]](https://redis.io/commands/hset) | __*O(1)*__ for each field/value pair added, so __*O(N)*__ to add N field/value pairs when the command is called with multiple field/value pairs | Set the string value of a hash field |
    | [__HSETNX__ key field value](https://redis.io/commands/hsetnx) | __*O(1)*__ | Set the value of a hash field, only if the field does not exist |
    | [__HSTRLEN__ key field](https://redis.io/commands/hstrlen) | __*O(1)*__ | Get the length of the value of a hash field |
    | [__HVALS__ key](https://redis.io/commands/hvals) | __*O(N)*__ where N is the size of the hash | Get all the values in a hash |
    | [__HSCAN__ key cursor [MATCH pattern] [COUNT count]](https://redis.io/commands/hscan) | __*O(1)*__ for every call. __*O(N)*__ for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection | Incrementally iterate hash fields and associated values |

1. # HyperLogLog

    | [__PFADD__ key element [element ...]](https://redis.io/commands/pfadd) | __*O(1)*__ to add every element | Adds the specified elements to the specified HyperLogLog |
    | [__PFCOUNT__ key [key ...]](https://redis.io/commands/pfcount) | __*O(1)*__ with a very small average constant time when called with a single key. __*O(N)*__ with N being the number of keys, and much bigger constant times, when called with multiple keys | Return the approximated cardinality of the set(s) observed by the HyperLogLog at key(s) |
    | [__PFMERGE__ destkey sourcekey [sourcekey ...]](https://redis.io/commands/pfmerge) | __*O(N)*__ to merge N HyperLogLogs, but with high constant times | Merge N different HyperLogLogs into a single one |

1. # Keys

    | [__COPY__ source destination [DB destination-db] [REPLACE]](https://redis.io/commands/copy) | __*O(N)*__ worst case for collections, where N is the number of nested items. __*O(1)*__ for string values | Copy a key |
    | [__DEL__ key [key ...]](https://redis.io/commands/del) | __*O(N)*__ where N is the number of keys that will be removed. When a key to remove holds a value other than a string, the individual complexity for this key is __*O(M)*__ where M is the number of elements in the list, set, sorted set or hash. Removing a single key that holds a string value is __*O(1)*__ | Delete a key |
    | [__DUMP__ key](https://redis.io/commands/dump) | __*O(1)*__ to access the key and additional __*O(NM)*__ to serialize it, where N is the number of Redis objects composing the value and M their average size. For small string values the time complexity is thus __*O(1)+O(M)*__ where M is small, so simply __*O(1)*__ | Return a serialized version of the value stored at the specified key |
    | [__EXISTS__ key [key ...]](https://redis.io/commands/exists) | __*O(1)*__ | Determine if a key exists |
    | [__EXPIRE__ key seconds](https://redis.io/commands/expire) | __*O(1)*__ | Set a key's time to live in seconds |
    | [__EXPIREAT__ key timestamp](https://redis.io/commands/expireat) | __*O(1)*__ | Set the expiration for a key as a UNIX timestamp |
    | [__KEYS__ pattern](https://redis.io/commands/keys) | __*O(N)*__ with N being the number of keys in the database, under the assumption that the key names in the database and the given pattern have limited length | Find all keys matching the given pattern |
    | [__MIGRATE__ host port key\|"" destination-db timeout [COPY] [REPLACE] [AUTH password] [AUTH2 username password] [KEYS key [key ...]]](https://redis.io/commands/migrate) | This command actually executes a DUMP+DEL in the source instance, and a RESTORE in the target instance. See the pages of these commands for time complexity. Also an __*O(N)*__ data transfer between the two instances is performed | Atomically transfer a key from a Redis instance to another one |
    | [__MOVE__ key db](https://redis.io/commands/move) | __*O(1)*__ | Move a key to another database |
    | [__OBJECT__ subcommand [arguments [arguments ...]]](https://redis.io/commands/object) | __*O(1)*__ for all the currently implemented subcommands | Inspect the internals of Redis objects |
    | [__PERSIST__ key](https://redis.io/commands/persist) | __*O(1)*__ | Remove the expiration from a key |
    | [__PEXPIRE__ key milliseconds](https://redis.io/commands/pexpire) | __*O(1)*__ | Set a key's time to live in milliseconds |
    | [__PEXPIREAT__ key milliseconds-timestamp](https://redis.io/commands/pexpireat) | __*O(1)*__ | Set the expiration for a key as a UNIX timestamp specified in milliseconds |
    | [__PTTL__ key](https://redis.io/commands/pttl) | __*O(1)*__ | Get the time to live for a key in milliseconds |
    | [__RANDOMKEY__](https://redis.io/commands/randomkey) | __*O(1)*__ | Return a random key from the keyspace |
    | [__RENAME__ key newkey](https://redis.io/commands/rename) | __*O(1)*__ | Rename a key |
    | [__RENAMENX__ key newkey](https://redis.io/commands/renamenx) | __*O(1)*__ | Rename a key, only if the new key does not exist |
    | [__RESTORE__ key ttl serialized-value [REPLACE] [ABSTTL] [IDLETIME seconds] [FREQ frequency]](https://redis.io/commands/restore) | __*O(1)*__ to create the new key and additional __*O(NM)*__ to reconstruct the serialized value, where N is the number of Redis objects composing the value and M their average size. For small string values the time complexity is thus __*O(1)+O(M)*__ where M is small, so simply __*O(1)*__. However for sorted set values the complexity is __*O(NMlog(N))*__ because inserting values into sorted sets is __*O(log(N))*__ | Create a key using the provided serialized value, previously obtained using DUMP |
    | [__SORT__ key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC\|DESC] [ALPHA] [STORE destination]](https://redis.io/commands/sort) | __*O(N+Mlog(M))*__ where N is the number of elements in the list or set to sort, and M the number of returned elements. When the elements are not sorted, complexity is currently __*O(N)*__ as there is a copy step that will be avoided in next releases | Sort the elements in a list, set or sorted set |
    | [__TOUCH__ key [key ...]](https://redis.io/commands/touch) | __*O(N)*__ where N is the number of keys that will be touched | Alters the last access time of a key(s). Returns the number of existing keys specified |
    | [__TTL__ key](https://redis.io/commands/ttl) | __*O(1)*__ | Get the time to live for a key |
    | [__TYPE__ key](https://redis.io/commands/type) | __*O(1)*__ | Determine the type stored at key |
    | [__UNLINK__ key [key ...]](https://redis.io/commands/unlink) | __*O(1)*__ for each key removed regardless of its size. Then the command does __*O(N)*__ work in a different thread in order to reclaim memory, where N is the number of allocations the deleted objects where composed of | Delete a key asynchronously in another thread. Otherwise it is just as DEL, but non blocking |
    | [__WAIT__ numreplicas timeout](https://redis.io/commands/wait) | __*O(1)*__ | Wait for the synchronous replication of all the write commands sent in the context of the current connection |
    | [__SCAN__ cursor [MATCH pattern] [COUNT count] [TYPE type]](https://redis.io/commands/scan) | __*O(1)*__ for every call. __*O(N)*__ for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection | Incrementally iterate the keys space |

1. # Lists

    | [__BLPOP__ key [key ...] timeout](https://redis.io/commands/blpop) | __*O(1)*__ | Remove and get the first element in a list, or block until one is available |
    | [__BRPOP__ key [key ...] timeout](https://redis.io/commands/brpop) | __*O(1)*__ | Remove and get the last element in a list, or block until one is available |
    | [__BRPOPLPUSH__ source destination timeout](https://redis.io/commands/brpoplpush) | __*O(1)*__ | Pop an element from a list, push it to another list and return it; or block until one is available |
    | [__BLMOVE__ source destination LEFT\|RIGHT LEFT\|RIGHT timeout](https://redis.io/commands/blmove) | __*O(1)*__ | Pop an element from a list, push it to another list and return it; or block until one is available |
    | [__LINDEX__ key index](https://redis.io/commands/lindex) | __*O(N)*__ where N is the number of elements to traverse to get to the element at index. This makes asking for the first or the last element of the list O(1) | Get an element from a list by its index |
    | [__LINSERT__ key BEFORE\|AFTER pivot element](https://redis.io/commands/linsert) | __*O(N)*__ where N is the number of elements to traverse before seeing the value pivot. This means that inserting somewhere on the left end on the list (head) can be considered __*O(1)*__ and inserting somewhere on the right end (tail) is __*O(N)*__ | Insert an element before or after another element in a list |
    | [__LLEN__ key](https://redis.io/commands/llen) | __*O(1)*__ | Get the length of a list |
    | [__LPOP__ key [count]](https://redis.io/commands/lpop) | __*O(N)*__ where N is the number of elements returned | Remove and get the first elements in a list |
    | [__LPOS__ key element [RANK rank] [COUNT num-matches] [MAXLEN len]](https://redis.io/commands/lpos) | __*O(N)*__ where N is the number of elements in the list, for the average case. When searching for elements near the head or the tail of the list, or when the MAXLEN option is provided, the command may run in constant time | Return the index of matching elements on a list |
    | [__LPUSH__ key element [element ...]](https://redis.io/commands/lpush) | __*O(1)*__ for each element added, so __*O(N)*__ to add N elements when the command is called with multiple arguments | Prepend one or multiple elements to a list |
    | [__LPUSHX__ key element [element ...]](https://redis.io/commands/lpushx) | __*O(1)*__ for each element added, so __*O(N)*__ to add N elements when the command is called with multiple arguments | Prepend an element to a list, only if the list exists |
    | [__LRANGE__ key start stop](https://redis.io/commands/lrange) | __*O(S+N)*__ where S is the distance of start offset from HEAD for small lists, from nearest end (HEAD or TAIL) for large lists; and N is the number of elements in the specified range | Get a range of elements from a list |
    | [__LREM__ key count element](https://redis.io/commands/lrem) | __*O(N+M)*__ where N is the length of the list and M is the number of elements removed | Remove elements from a list |
    | [__LSET__ key index element](https://redis.io/commands/lset) | __*O(N)*__ where N is the length of the list. Setting either the first or the last element of the list is __*O(1)*__ | Set the value of an element in a list by its index |
    | [__LTRIM__ key start stop](https://redis.io/commands/ltrim) | __*O(N)*__ where N is the number of elements to be removed by the operation | Trim a list to the specified range |
    | [__RPOP__ key [count]](https://redis.io/commands/rpop) | __*O(N)*__ where N is the number of elements returned | Remove and get the last elements in a list |
    | [__RPOPLPUSH__ source destination](https://redis.io/commands/rpoplpush) | __*O(1)*__ | Remove the last element in a list, prepend it to another list and return it |
    | [__LMOVE__ source destination LEFT\|RIGHT LEFT\|RIGHT](https://redis.io/commands/lmove) | __*O(1)*__ | Pop an element from a list, push it to another list and return it |
    | [__RPUSH__ key element [element ...]](https://redis.io/commands/rpush) | __*O(1)*__ for each element added, so __*O(N)*__ to add N elements when the command is called with multiple arguments | Append one or multiple elements to a list |
    | [__RPUSHX__ key element [element ...]](https://redis.io/commands/rpushx) | __*O(1)*__ for each element added, so __*O(N)*__ to add N elements when the command is called with multiple arguments | Append an element to a list, only if the list exists |

1. # Pub/Sub

    | [__PSUBSCRIBE__ pattern [pattern ...]](https://redis.io/commands/psubscribe) | __*O(N)*__ where N is the number of patterns the client is already subscribed to | Listen for messages published to channels matching the given patterns |
    | [__PUBSUB__ subcommand [argument [argument ...]]](https://redis.io/commands/pubsub) | __*O(N)*__ for the CHANNELS subcommand, where N is the number of active channels, and assuming constant time pattern matching (relatively short channels and patterns). __*O(N)*__ for the NUMSUB subcommand, where N is the number of requested channels. __*O(1)*__ for the NUMPAT subcommand | Inspect the state of the Pub/Sub subsystem |
    | [__PUBLISH__ channel message](https://redis.io/commands/publish) | __*O(N+M)*__ where N is the number of clients subscribed to the receiving channel and M is the total number of subscribed patterns (by any client) | Post a message to a channel |
    | [__PUNSUBSCRIBE__ [pattern [pattern ...]]](https://redis.io/commands/punsubscribe) | __*O(N+M)*__ where N is the number of patterns the client is already subscribed and M is the number of total patterns subscribed in the system (by any client) | Stop listening for messages posted to channels matching the given patterns |
    | [__SUBSCRIBE__ channel [channel ...]](https://redis.io/commands/subscribe) | __*O(N)*__ where N is the number of channels to subscribe to | Listen for messages published to the given channels |
    | [__UNSUBSCRIBE__ [channel [channel ...]]](https://redis.io/commands/unsubscribe) | __*O(N)*__ where N is the number of clients already subscribed to a channel | Stop listening for messages posted to the given channels |

1. # Scripting

    | [__EVAL__ script numkeys key [key ...] arg [arg ...]](https://redis.io/commands/eval) | Depends on the script executed | Execute a Lua script server side |
    | [__EVALSHA__ sha1 numkeys key [key ...] arg [arg ...]](https://redis.io/commands/evalsha) | Depends on the script executed | Execute a Lua script server side |
    | [__SCRIPT__ DEBUG YES\|SYNC\|NO](https://redis.io/commands/script-debug) | __*O(1)*__ | Set the debug mode for executed scripts |
    | [__SCRIPT__ EXISTS sha1 [sha1 ...]](https://redis.io/commands/script-exists) | __*O(N)*__ - N number of scripts to check | Check existence of scripts in the script cache |
    | [__SCRIPT__ FLUSH](https://redis.io/commands/script-flush) | __*O(N)*__ with N being the number of scripts in cache | Remove all the scripts from the script cache |
    | [__SCRIPT__ KILL](https://redis.io/commands/script-kill) | __*O(1)*__ | Kill the script currently in execution |
    | [__SCRIPT__ LOAD script](https://redis.io/commands/script-load) | __*O(N)*__ with N being the length in bytes of the script body | Load the specified Lua script into the script cache |

1. # Server

    | [__ACL LOAD__](https://redis.io/commands/acl-load) | __*O(N)*__. Where N is the number of configured users | Reload the ACLs from the configured ACL file |
    | [__ACL SAVE__](https://redis.io/commands/acl-save) | __*O(N)*__. Where N is the number of configured users | Save the current ACL rules in the configured ACL file |
    | [__ACL LIST__](https://redis.io/commands/acl-list) | __*O(N)*__. Where N is the number of configured users | List the current ACL rules in ACL config file format |
    | [__ACL USERS__](https://redis.io/commands/acl-users) | __*O(N)*__. Where N is the number of configured users | List the username of all the configured ACL rules |
    | [__ACL GETUSER__ username](https://redis.io/commands/acl-getuser) | __*O(N)*__. Where N is the number of password, command and pattern rules that the user has | Get the rules for a specific ACL user |
    | [__ACL SETUSER__ username [rule [rule ...]]](https://redis.io/commands/acl-setuser) | __*O(N)*__. Where N is the number of rules provided | Modify or create the rules for a specific ACL user |
    | [__ACL DELUSER__ username [username ...]](https://redis.io/commands/acl-deluser) | __*O(1)*__ amortized time considering the typical user | Remove the specified ACL users and the associated rules |
    | [__ACL CAT__ [categoryname]](https://redis.io/commands/acl-cat) | __*O(1)*__ since the categories and commands are a fixed set | List the ACL categories or the commands inside a category |
    | [__ACL GENPASS__ [bits]](https://redis.io/commands/acl-genpass) | __*O(1)*__ | Generate a pseudorandom secure password to use for ACL users |
    | [__ACL WHOAMI__](https://redis.io/commands/acl-whoami) | __*O(1)*__ | Return the name of the user associated to the current connection |
    | [__ACL LOG__ [count or RESET]](https://redis.io/commands/acl-log) | __*O(N)*__ with N being the number of entries shown | List latest events denied because of ACLs in place |
    | [__ACL HELP__](https://redis.io/commands/acl-help) | __*O(1)*__ | Show helpful text about the different subcommands |
    | [__BGREWRITEAOF__](https://redis.io/commands/bgrewriteaof) | - | Asynchronously rewrite the append-only file |
    | [__BGSAVE__ [SCHEDULE]](https://redis.io/commands/bgsave) | - | Asynchronously save the dataset to disk |
    | [__COMMAND__](https://redis.io/commands/command) | __*O(N)*__ where N is the total number of Redis commands | Get array of Redis command details |
    | [__COMMAND COUNT__](https://redis.io/commands/command-count) | __*O(1)*__ | Get total number of Redis commands |
    | [__COMMAND GETKEYS__](https://redis.io/commands/command-getkeys) | __*O(N)*__ where N is the number of arguments to the command | Extract keys given a full Redis command |
    | [__COMMAND INFO__ command-name [command-name ...]](https://redis.io/commands/command-info) | __*O(N)*__ when N is number of commands to look up | Get array of specific Redis command details |
    | [__CONFIG GET__ parameter](https://redis.io/commands/config-get) | - | Get the value of a configuration parameter |
    | [__CONFIG REWRITE__](https://redis.io/commands/config-rewrite) | - | Rewrite the configuration file with the in memory configuration |
    | [__CONFIG SET__ parameter value](https://redis.io/commands/config-set) | - | Set a configuration parameter to the given value |
    | [__CONFIG RESETSTAT__](https://redis.io/commands/config-resetstat) | __*O(1)*__ | Reset the stats returned by INFO |
    | [__DBSIZE__](https://redis.io/commands/dbsize) | - | Return the number of keys in the selected database |
    | [__DEBUG OBJECT__ key](https://redis.io/commands/debug-object) | - | Get debugging information about a key |
    | [__DEBUG SEGFAULT__](https://redis.io/commands/debug-segfault) | - | Make the server crash |
    | [__FLUSHALL__ [ASYNC]](https://redis.io/commands/flushall) | - | Remove all keys from all databases |
    | [__FLUSHDB__ [ASYNC]](https://redis.io/commands/flushdb) | - | Remove all keys from the current database |
    | [__INFO__ [section]](https://redis.io/commands/info) | - | Get information and statistics about the server |
    | [__LOLWUT__ [VERSION version]](https://redis.io/commands/lolwut) | - | Display some computer art and the Redis version |
    | [__LASTSAVE__](https://redis.io/commands/lastsave) | - | Get the UNIX time stamp of the last successful save to disk |
    | [__MEMORY DOCTOR__](https://redis.io/commands/memory-doctor) | - | Outputs memory problems report |
    | [__MEMORY HELP__](https://redis.io/commands/memory-help) | - | Show helpful text about the different subcommands |
    | [__MEMORY MALLOC-STATS__](https://redis.io/commands/memory-malloc-stats) | - | Show allocator internal stats |
    | [__MEMORY PURGE__](https://redis.io/commands/memory-purge) | - | Ask the allocator to release memory |
    | [__MEMORY STATS__](https://redis.io/commands/memory-stats) | - | Show memory usage details |
    | [__MEMORY USAGE__ key [SAMPLES count]](https://redis.io/commands/memory-usage) | __*O(N)*__ where N is the number of samples | Estimate the memory usage of a key |
    | [__MODULE LIST__](https://redis.io/commands/module-list) | __*O(N)*__ where N is the number of loaded modules | List all modules loaded by the server |
    | [__MODULE LOAD__ path [ arg [arg ...]]](https://redis.io/commands/module-load) | __*O(1)*__ | Load a module |
    | [__MODULE UNLOAD__ name](https://redis.io/commands/module-unload) | __*O(1)*__ | Unload a module |
    | [__MONITOR__](https://redis.io/commands/monitor) | - | Listen for all requests received by the server in real time |
    | [__ROLE__](https://redis.io/commands/role) | - | Return the role of the instance in the context of replication |
    | [__SAVE__](https://redis.io/commands/save) | - | Synchronously save the dataset to disk |
    | [__SHUTDOWN__ [NOSAVE/SAVE]](https://redis.io/commands/shutdown) | - | Synchronously save the dataset to disk and then shut down the server |
    | [__SLAVEOF__ host port](https://redis.io/commands/slaveof) | - | Make the server a replica of another instance, or promote it as master. Deprecated starting with Redis 5. Use REPLICAOF instead |
    | [__REPLICAOF__ host port](https://redis.io/commands/replicaof) | - | Make the server a replica of another instance, or promote it as master |
    | [__SLOWLOG__ subcommand [argument]](https://redis.io/commands/slowlog) | - | Manages the Redis slow queries log |
    | [__SWAPDB__ index1 index2](https://redis.io/commands/swapdb) | __*O(N)*__ where N is the count of clients watching or blocking on keys from both databases | Swaps two Redis databases |
    | [__SYNC__](https://redis.io/commands/sync) | - | Internal command used for replication |
    | [__PSYNC__ replicationid offset](https://redis.io/commands/psync) | - | Internal command used for replication |
    | [__TIME__](https://redis.io/commands/time) | __*O(1)*__ | Return the current server time |
    | [__LATENCY DOCTOR__](https://redis.io/commands/latency-doctor) | - | Return a human readable latency analysis report |
    | [__LATENCY GRAPH__ event](https://redis.io/commands/latency-graph) | - | Return a latency graph for the event |
    | [__LATENCY HISTORY__ event](https://redis.io/commands/latency-history) | - | Return timestamp-latency samples for the event |
    | [__LATENCY LATEST__](https://redis.io/commands/latency-latest) | - | Return the latest latency samples for all events |
    | [__LATENCY RESET__ [event [event ...]]](https://redis.io/commands/latency-reset) | - | Reset latency data for one or more events |
    | [__LATENCY HELP__](https://redis.io/commands/latency-help) | - | Show helpful text about the different subcommands |

1. # Sets

    | [__SADD__ key member [member ...]](https://redis.io/commands/sadd) | __*O(1)*__ for each element added, so __*O(N)*__ to add N elements when the command is called with multiple arguments | Add one or more members to a set |
    | [__SCARD__ key](https://redis.io/commands/scard) | __*O(1)*__ | Get the number of members in a set |
    | [__SDIFF__ key [key ...]](https://redis.io/commands/sdiff) | __*O(N)*__ where N is the total number of elements in all given sets | Subtract multiple sets |
    | [__SDIFFSTORE__ destination key [key ...]](https://redis.io/commands/sdiffstore) | __*O(N)*__ where N is the total number of elements in all given sets | Subtract multiple sets and store the resulting set in a key |
    | [__SINTER__ key [key ...]](https://redis.io/commands/sinter) | __*O(NM)*__ worst case where N is the cardinality of the smallest set and M is the number of sets | Intersect multiple sets |
    | [__SINTERSTORE__ destination key [key ...]](https://redis.io/commands/sinterstore) | __*O(NM)*__ worst case where N is the cardinality of the smallest set and M is the number of sets | Intersect multiple sets and store the resulting set in a key |
    | [__SISMEMBER__ key member](https://redis.io/commands/sismember) | __*O(1)*__ | Determine if a given value is a member of a set |
    | [__SMISMEMBER__ key member [member ...]](https://redis.io/commands/smismember) | __*O(N)*__ where N is the number of elements being checked for membership | Returns the membership associated with the given elements for a set |
    | [__SMEMBERS__ key](https://redis.io/commands/smembers) | __*O(N)*__ where N is the set cardinality | Get all the members in a set |
    | [__SMOVE__ source destination member](https://redis.io/commands/smove) | __*O(1)*__ | Move a member from one set to another |
    | [__SPOP__ key [count]](https://redis.io/commands/spop) | __*O(1)*__ | Remove and return one or multiple random members from a set |
    | [__SRANDMEMBER__ key [count]](https://redis.io/commands/srandmember) | Without the count argument __*O(1)*__, otherwise __*O(N)*__ where N is the absolute value of the passed count | Get one or multiple random members from a set |
    | [__SREM__ key member [member ...]](https://redis.io/commands/srem) | __*O(N)*__ where N is the number of members to be removed | Remove one or more members from a set |
    | [__SUNION__ key [key ...]](https://redis.io/commands/sunion) | __*O(N)*__ where N is the total number of elements in all given sets | Add multiple sets |
    | [__SUNIONSTORE__ destination key [key ...]](https://redis.io/commands/sunionstore) | __*O(N)*__ where N is the total number of elements in all given sets | Add multiple sets and store the resulting set in a key |
    | [__SSCAN__ key cursor [MATCH pattern] [COUNT count]](https://redis.io/commands/sscan) | __*O(1)*__ for every call. __*O(N)*__ for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection | Incrementally iterate Set elements |

1. # Sorted Sets

    | [__BZPOPMIN__ key [key ...] timeout](https://redis.io/commands/bzpopmin) | __*O(log(N))*__ with N being the number of elements in the sorted set | Remove and return the member with the lowest score from one or more sorted sets, or block until one is available |
    | [__BZPOPMAX__ key [key ...] timeout](https://redis.io/commands/bzpopmax) | __*O(log(N))*__ with N being the number of elements in the sorted set | Remove and return the member with the highest score from one or more sorted sets, or block until one is available |
    | [__ZADD__ key [NX\|XX] [GT\|LT] [CH] [INCR] score member [score member ...]](https://redis.io/commands/zadd) | __*O(log(N))*__ for each item added, where N is the number of elements in the sorted set | Add one or more members to a sorted set, or update its score if it already exists |
    | [__ZCARD__ key](https://redis.io/commands/zcard) | __*O(1)*__ | Get the number of members in a sorted set |
    | [__ZCOUNT__ key min max](https://redis.io/commands/zcount) | __*O(log(N))*__ with N being the number of elements in the sorted set | Count the members in a sorted set with scores within the given values |
    | [__ZDIFF__ numkeys key [key ...] [WITHSCORES]](https://redis.io/commands/zdiff) | __*O(L + (N-K)log(N))*__ worst case where L is the total number of elements in all the sets, N is the size of the first set, and K is the size of the result set | Subtract multiple sorted sets |
    | [__ZDIFFSTORE__ destination numkeys key [key ...]](https://redis.io/commands/zdiffstore) | __*O(L + (N-K)log(N))*__ worst case where L is the total number of elements in all the sets, N is the size of the first set, and K is the size of the result set | Subtract multiple sorted sets and store the resulting sorted set in a new key |
    | [__ZINCRBY__ key increment member](https://redis.io/commands/zincrby) | __*O(log(N))*__ where N is the number of elements in the sorted set | Increment the score of a member in a sorted set |
    | [__ZINTER__ numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM\|MIN\|MAX] [WITHSCORES]](https://redis.io/commands/zinter) | __*O(NK)+O(Mlog(M))*__ worst case with N being the smallest input sorted set, K being the number of input sorted sets and M being the number of elements in the resulting sorted set | Intersect multiple sorted sets |
    | [__ZINTERSTORE__ destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM\|MIN\|MAX]](https://redis.io/commands/zinterstore) | *__O(NK)+O(Mlog(M))*__ worst case with N being the smallest input sorted set, K being the number of input sorted sets and M being the number of elements in the resulting sorted set | Intersect multiple sorted sets and store the resulting sorted set in a new key |
    | [__ZLEXCOUNT__ key min max](https://redis.io/commands/zlexcount) | __*O(log(N))*__ with N being the number of elements in the sorted set | Count the number of members in a sorted set between a given lexicographical range |
    | [__ZPOPMAX__ key [count]](https://redis.io/commands/zpopmax) | __*O(Mlog(N))*__ with N being the number of elements in the sorted set, and M being the number of elements popped | Remove and return members with the highest scores in a sorted set |
    | [__ZPOPMIN__ key [count]](https://redis.io/commands/zpopmin) | __*O(Mlog(N))*__ with N being the number of elements in the sorted set, and M being the number of elements popped | Remove and return members with the lowest scores in a sorted set |
    | [__ZRANGESTORE__ dst src min max [BYSCORE/BYLEX] [REV] [LIMIT offset count]](https://redis.io/commands/zrangestore) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements stored into the destination key | Store a range of members from sorted set into another key |
    | [__ZRANGE__ key min max [BYSCORE\|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]](https://redis.io/commands/zrange) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements returned | Return a range of members in a sorted set |
    | [__ZRANGEBYLEX__ key min max [LIMIT offset count]](https://redis.io/commands/zrangebylex) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it __*O(log(N))*__ | Return a range of members in a sorted set, by lexicographical range |
    | [__ZREVRANGEBYLEX__ key max min [LIMIT offset count]](https://redis.io/commands/zrevrangebylex) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it __*O(log(N))*__ | Return a range of members in a sorted set, by lexicographical range, ordered from higher to lower strings |
    | [__ZRANGEBYSCORE__ key min max [WITHSCORES] [LIMIT offset count]](https://redis.io/commands/zrangebyscore) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it __*O(log(N))*__ | Return a range of members in a sorted set, by score |
    | [__ZRANK__ key member](https://redis.io/commands/zrank) | __*O(log(N))*__ | Determine the index of a member in a sorted set |
    | [__ZREM__ key member [member ...]](https://redis.io/commands/zrem) | __*O(Mlog(N))*__ with N being the number of elements in the sorted set and M the number of elements to be removed | Remove one or more members from a sorted set |
    | [__ZREMRANGEBYLEX__ key min max](https://redis.io/commands/zremrangebylex) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements removed by the operation | Remove all members in a sorted set between the given lexicographical range |
    | [__ZREMRANGEBYRANK__ key start stop](https://redis.io/commands/zremrangebyrank) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements removed by the operation | Remove all members in a sorted set within the given indexes |
    | [__ZREMRANGEBYSCORE__ key min max](https://redis.io/commands/zremrangebyscore) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements removed by the operation. | Remove all members in a sorted set within the given scores |
    | [__ZREVRANGE__ key start stop [WITHSCORES]](https://redis.io/commands/zrevrange) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements returned | Return a range of members in a sorted set, by index, with scores ordered from high to low |
    | [__ZREVRANGEBYSCORE__ key max min [WITHSCORES] [LIMIT offset count]](https://redis.io/commands/zrevrangebyscore) | __*O(log(N)+M)*__ with N being the number of elements in the sorted set and M the number of elements being returned. If M is constant (e.g. always asking for the first 10 elements with LIMIT), you can consider it __*O(log(N))*__ | Return a range of members in a sorted set, by score, with scores ordered from high to low |
    | [__ZREVRANK__ key member](https://redis.io/commands/zrevrank) | __*O(log(N))*__ | Determine the index of a member in a sorted set, with scores ordered from high to low |
    | [__ZSCORE__ key member](https://redis.io/commands/zscore) | __*O(1)*__ | Get the score associated with the given member in a sorted set |
    | [__ZUNION__ numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM\|MIN\|MAX] [WITHSCORES]](https://redis.io/commands/zunion) | *__O(N)+O(Mlog(M))*__ with N being the sum of the sizes of the input sorted sets, and M being the number of elements in the resulting sorted set | Add multiple sorted sets |
    | [__ZMSCORE__ key member [member ...]](https://redis.io/commands/zmscore) | __*O(N)*__ where N is the number of members being requested | Get the score associated with the given members in a sorted set |
    | [__ZUNIONSTORE__ destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM\|MIN\|MAX]](https://redis.io/commands/zunionstore) | __*O(N)+O(M log(M))*__ with N being the sum of the sizes of the input sorted sets, and M being the number of elements in the resulting sorted set | Add multiple sorted sets and store the resulting sorted set in a new key |
    | [__ZSCAN__ key cursor [MATCH pattern] [COUNT count]](https://redis.io/commands/zscan) | __*O(1)*__ for every call. __*O(N)*__ for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection | Incrementally iterate sorted sets elements and associated scores |

1. # Streams

    | [__XINFO__ [CONSUMERS key groupname] [GROUPS key] [STREAM key] [HELP]](https://redis.io/commands/xinfo) | __*O(N)*__ with N being the number of returned items for the subcommands CONSUMERS and GROUPS. The STREAM subcommand is __*O(log N)*__ with N being the number of items in the stream | Get information on streams and consumer groups |
    | [__XADD__ key [NOMKSTREAM] [MAXLEN\|MINID [=\|~] threshold [LIMIT count]] *\|ID field value [field value ...]](https://redis.io/commands/xadd) | __*O(1)*__ when adding a new entry, __*O(N)*__ when trimming where N being the number of entires evicted | Appends a new entry to a stream |
    | [__XTRIM__ key MAXLEN\|MINID [=\|~] threshold [LIMIT count]](https://redis.io/commands/xtrim) | __*O(N)*__, with N being the number of evicted entries | Trims the stream to (approximately if '~' is passed) a certain size |
    | [__XDEL__ key ID [ID ...]](https://redis.io/commands/xdel) | __*O(1)*__ for each single item to delete in the stream, regardless of the stream size | Removes the specified entries from the stream. Returns the number of items actually deleted, that may be different from the number of IDs passed in case certain IDs do not exist |
    | [__XRANGE__ key start end [COUNT count]](https://redis.io/commands/xrange) | __*O(N)*__ with N being the number of elements being returned | Return a range of elements in a stream, with IDs matching the specified IDs interval |
    | [__XREVRANGE__ key end start [COUNT count]](https://redis.io/commands/xrevrange) | __*O(N)*__ with N being the number of elements being returned | Return a range of elements in a stream, with IDs matching the specified IDs interval, in reverse order (from greater to smaller IDs compared to XRANGE |
    | [__XLEN__ key](https://redis.io/commands/xlen) | __*O(1)*__ | Return the number of entries in a stream |
    | [__XREAD__ [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] ID [ID ...]](https://redis.io/commands/xread) | For each stream mentioned: __*O(N)*__ with N being the number of elements being returned, it means that XREAD-ing with a fixed COUNT is __*O(1)*__. Note that when the BLOCK option is used, XADD will pay __*O(M)*__ time in order to serve the M clients blocked on the stream getting new data | Return never seen elements in multiple streams, with IDs greater than the ones reported by the caller for each stream. Can block |
    | [__XGROUP__ [CREATE key groupname ID\|$ [MKSTREAM]] [SETID key groupname ID\|$] [DESTROY key groupname] [CREATECONSUMER key groupname consumername] [DELCONSUMER key groupname consumername]](https://redis.io/commands/xgroup) | __*O(1)*__ for all the subcommands, with the exception of the DESTROY subcommand which takes an additional __*O(M)*__ time in order to delete the M entries inside the consumer group pending entries list (PEL) | Create, destroy, and manage consumer groups |
    | [__XREADGROUP GROUP__ group consumer [COUNT count] [BLOCK milliseconds] [NOACK] STREAMS key [key ...] ID [ID ...]](https://redis.io/commands/xreadgroup) | For each stream mentioned: __*O(M)*__ with M being the number of elements returned. If M is constant (e.g. always asking for the first 10 elements with COUNT), you can consider it __*O(1)*__. On the other side when XREADGROUP blocks, XADD will pay the __*O(N)*__ time in order to serve the N clients blocked on the stream getting new data | Return new entries from a stream using a consumer group, or access the history of the pending entries for a given consumer. Can block |
    | [__XACK__ key group ID [ID ...]](https://redis.io/commands/xack) | __*O(1)*__ | Marks a pending message as correctly processed, effectively removing it from the pending entries list of the consumer group. Return value of the command is the number of messages successfully acknowledged, that is, the IDs we were actually able to resolve in the PEL |
    | [__XCLAIM__ key group consumer min-idle-time ID [ID ...] [IDLE ms] [TIME ms-unix-time] [RETRYCOUNT count] [FORCE] [JUSTID]](https://redis.io/commands/xclaim) | __*O(log N)*__ | Changes (or acquires) ownership of a message in a consumer group, as if the message was delivered to the specified consumer |
    | [__XAUTOCLAIM__ key group consumer min-idle-time start [COUNT count] [JUSTID]](https://redis.io/commands/xautoclaim) | __*O(1)*__ if COUNT is small | Changes (or acquires) ownership of messages in a consumer group, as if the messages were delivered to the specified consumer |
    | [__XPENDING__ key group [[IDLE min-idle-time] start end count [consumer]]](https://redis.io/commands/xpending) | __*O(N)*__ with N being the number of elements returned, so asking for a small fixed number of entries per call is __*O(1)*__. __*O(M)*__, where M is the total number of entries scanned when used with the IDLE filter. When the command returns just the summary and the list of consumers is small, it runs in __*O(1)*__ time; otherwise, an additional __*O(N)*__ time for iterating every consumer | Return information and entries from a stream consumer group pending entries list, that are messages fetched but never acknowledged |

1. # Strings

    | [__APPEND__ key value](https://redis.io/commands/append) | __*O(1)*__ | Append a value to a key
    | [__BITCOUNT__ key [start end]](https://redis.io/commands/bitcount) | __*O(N)*__ | Count set bits in a string
    | [__BITFIELD__ key [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP\|SAT\|FAIL]](https://redis.io/commands/bitfield) | __*O(1)*__ | Perform arbitrary bitfield integer operations on strings. Multiple operations, returns array |
    | [__BITOP__ operation destkey key [key ...]](https://redis.io/commands/bitop) | __*O(N)*__ | Perform bitwise operations between strings |
    | [__BITPOS__ key bit [start] [end]](https://redis.io/commands/bitpos) | __*O(N)*__ | Find first bit set or clear in a string, start and end are interpreted as range of bytes |
    | [__DECR__ key](https://redis.io/commands/decr) | __*O(1)*__ | Decrement the integer value of a key by one. Limited to 64 bit signed integers |
    | [__DECRBY__ key decrement](https://redis.io/commands/decrby) | __*O(1)*__ | Decrement the integer value of a key by the given number |
    | [__GET__ key](https://redis.io/commands/get) | __*O(1)*__ | Get the value of a key |
    | [__GETBIT__ key offset](https://redis.io/commands/getbit) | __*O(1)*__ | Returns the bit value at offset in the string value stored at key |
    | [__GETRANGE__ key start end](https://redis.io/commands/getrange) | __*O(N)*__; __*O(1)*__ for small strings | Get a substring of the string stored at a key |
    | [__GETSET__ key value](https://redis.io/commands/getset) | __*O(1)*__ | Set the string value of a key and return its old value |
    | [__INCR__ key](https://redis.io/commands/incr) | __*O(1)*__ | Increment the integer value of a key by one |
    | [__INCRBY__ key increment](https://redis.io/commands/incrby) | __*O(1)*__ | Increment the integer value of a key by the given amount |
    | [__INCRBYFLOAT__ key increment](https://redis.io/commands/incrbyfloat) | __*O(1)*__ | Increment the float value of a key by the given amount |
    | [__MGET__ key [key ...]](https://redis.io/commands/mget) | __*O(N)*__ where N is the no. of keys to retrieve | Get the values of all the given keys |
    | [__MSET__ key value [key value ...]](https://redis.io/commands/mset) | __*O(N)*__ where N is the no. of keys to set | Set multiple keys to multiple values. It is atomic |
    | [__MSETNX__ key value [key value ...]](https://redis.io/commands/msetnx) | __*O(N)*__ where N is the no. of keys to set | Set multiple keys to multiple values, only if none of the keys exist |
    | [__PSETEX__ key milliseconds value](https://redis.io/commands/psetex) | __*O(1)*__ | Set the value and expiration in milliseconds of a key |
    | [__SET__ key value [EX seconds\|PX milliseconds\|KEEPTTL] [NX\|XX] [GET]](https://redis.io/commands/set) | __*O(1)*__ | Set the string value of a key |
    | [__SETBIT__ key offset value](https://redis.io/commands/setbit) | __*O(1)*__ | Sets or clears the bit at offset in the string value stored at key |
    | [__SETEX__ key seconds value](https://redis.io/commands/setex) | __*O(1)*__ | Set the value and expiration of a key |
    | [__SETNX__ key value](https://redis.io/commands/setnx) | __*O(1)*__ | Set the value of a key, only if the key does not exist |
    | [__SETRANGE__ key offset](https://redis.io/commands/setrange) | __*O(1)*__ | Overwrite part of a string at key starting at the specified offset |
    | [__STRALGO__ LCS algo-specific-argument [algo-specific-argument ...]](https://redis.io/commands/stralgo) | __*O(strlen(s1).strlen(s2))*__ | Run algorithms (currently LCS) against strings. `STRALGO LCS [KEYS ...] [STRINGS ...] [LEN] [IDX] [MINMATCHLEN <len>] [WITHMATCHLEN]` |
    | [__STRLEN__ key](https://redis.io/commands/strlen) | __*O(1)*__ | Get the length of the value stored in a key |

1. # Transactions

    | [__DISCARD__](https://redis.io/commands/discard) | Discard all commands issued after MULTI |
    | [__EXEC__](https://redis.io/commands/exec) | Execute all commands issued after MULTI |
    | [__MULTI__](https://redis.io/commands/multi) | Mark the start of a transaction block |
    | [__UNWATCH__](https://redis.io/commands/unwatch) | Forget about all watched keys |
    | [__WATCH__ key [key ...]](https://redis.io/commands/watch) | Watch the given keys to determine execution of theMULTI/EXEC block |