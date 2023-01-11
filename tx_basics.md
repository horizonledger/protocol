
# basic principles of horizon protocol

vertex - any endpoints in the system (validators, nodes, peers, clients, ...)

messages are the data or programs which flow between vertex

transactions are the state transitions which are recorded permanently

## names

there are different kinds of names, and they are built into the system from the start

if a user creates a publickey immediately a name is required, so that humans don't deal with pubkeys

similarly transactionsid can be named. block1/txid123

domain names can be cross owned by validating dns records and attaching the public key

therefore any URI can be connected to a public key


// uname-register
// uname-transfer
// dname-register. dns record contanis public key. verify and make onchain tx
// transfer-value. money transaction
// contract-create
// contract-call