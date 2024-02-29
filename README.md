# switcheo-code-challenge
## Problem 4: 3 ways to sum to n
I created the 3 functions in Rust. I calculated the sum to n using a while loop, recursion  and also mathematical formula. The while loop and recursion solutions are both O(n) whereas the sum of Arithmetic Progression would likely take O(1) time

## Problem 5: A crude chain
THe task here was to implement a CRUD interface that allows the user to interact with articles, that contain a title and body. The blockchain in question can be found in `src/problem5/blockchain`. I followed the guide on the Ignite CLI documentation, and I was able to implement all of the CRUD features.

However, some queries are still not working as well. For example, when we perform `list-post`, I am currently only returning `QueryListPostResponse`, which was initialized with just a single post at the start of the list. As such `list-post` is not accurate yet.

## Consensus Breaking
### What is it?
Blockchain technology relies on participation of different nodes in order to verify the validity of the state of the blockchain, and thus the transaction. Consensus is important as it is the sole mechanism that ensures that the transaction is valid. When this consensus is broken, there is no longer agreement about the state of the blockchain, we can lead to data inconsistencies, security risks DUE to these inconsistencies. There will also be loss of trust within the participants of the decentralized network.

### How does my change break consensus?
Blockchains creted by Cosmos SDK will have a "Proof of Stake" consensus protocol that requires validators to secure the chain. That being said, in order to perform initial coin allocations and setting validators, these parameters sould be stored in a JSON formatted genesis file. This genesis file contains the initial state. My idea was to find where the initial condition was stated. (I assumed it was config.yml). Then, I removed the account balance from Alice's account. I believe that upon generation of the Genesis, it could be the case that the other participants are unable to know Alice's information. This could potentially lead to consensus breaking.

## Problem 6: Transaction Broadcaster Service

1. First, we want to create a HTTP API endpoint, which makes use of the broadcast service internal API. We should only expect to receive incoming POST requests that contain data as well as message type. Based on whether we succeed or fail, we return HTTP Response 200 OK (success) or 4XX-5XX (fail)

2. The broadcaster service should receive the request from the endpoint above. It will sign the data using their private key, before broadcasting it to an EVM-compatible blockchain network. We can do this by maintaining a message queue, and enqueue the broadcast requests as they are created.
Then, we catch the case where the broadcasting fails, and retry automatically. We could use the exponential backoff algorithm here (Increase the time to retry exponentially for every failed retry), enqeueing the failed process back to the message queue.  By doing this, we can ensure our system doesn't get overloaded.

3. In order to track the status of a transactions, we can create a DB, where the records contain (Transaction ID, Data, Status, Timestamp, Retry Count, Additional Metadata)
There should be an interface to query this database by Transaction ID.

4. For administrators to manually retry transactions, they should have a separate interface, which allows them to see and query all failed transactions, select 1 or more failed transaction. The broadcaster service in 2 should then retry the transaction by resigning the data etc. 
