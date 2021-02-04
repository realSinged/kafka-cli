# Kafka-cli

A command line tool for apache kafka, include topic,consumer,producer, admin's operations

## Features
- **General**
    - Works with modern Kafka cluster (1.0+)
    - Connection on standard or ssl, sasl cluster
    - Multi cluster

- **Topics**
    - list
    - describe topic partitions,replicas
    - create
    - delete
    - add partitions
    - 
* Browse Topic datas

View data, offset, key, timestamp & headers

Automatic deserialization of avro message encoded with schema registry

Configurations view

Logs view

Delete a record

Empty a Topic (Delete all the record from one topic)

Sort view

Filter per partitions

Filter with a starting time

Filter data with a search string

* Consumer Groups (only with kafka internal storage, not with the old Zookeeper one)

List with lag, topics assignments

Partitions view & lag

ACLS view

Node leader & assignments view

Display active and pending consumers groups

Delete a consumer group

Update consumer group offsets to start / end / timestamp

* Schema Registry

List schema

Create / Update / Delete a schema

View and delete individual schema version

* Connect

List connect definition

Create / Update / Delete a definition

Pause / Resume / Restart a definition or a task

* Nodes

List
Configurations view

Logs view

Configure a node

* ACLS

List principals

List principals topic & group acls

* Authentification and Roles

Read only mode

BasicHttp with roles per user

User groups configuration

Filter topics with regexp for current groups

Ldap configuration to match AKHQ groups/roles

## Installation

    git clone https://github.com/realSinged/kafka-cli.git
    cd kafka-cli
    go build cmd/kafka-cli.go

## Usage
**Overview**:

    ./kafka-cli
    
    A command line tools for apache kafka, include topic,consumer,producer, admin's operations
    
    Usage:
      kafka-cli [flags]
      kafka-cli [command]
    
    Available Commands:
      admin       
      consume     Consume kafka message with given topic and group_id
      help        Help about any command
      topic       Kafka topic operations
    
    Flags:
      -h, --help   help for kafka-cli
    
    Use "kafka-cli [command] --help" for more information about a command.
    
**Topic**:
    
    ./kafka-cli.go topic -h   
                                       
    Topic operations, include topic create、list、delete、detail, topic partition create
    
    Usage:
      kafka-cli topic [flags]
    
    Examples:
    
    # Create a topic
        ./kafka-cli topic -c=singed  --partition-num=10 --replica-num=1
        result: 
            {"level":"info","ts":1612423941.894614,"caller":"log/log.go:16","msg":"Create topic success","topic":"singed","partition num":10,"replica num":1}
    
    # List all available topics.
        ./kafka-cli topic -l
    
    # List details for the given topics.more than one should be separated by commas
        ./kafka-cli topic --describe=singed
        result:
            ****************************************
                    TOPIC:singed
                    DETAIL:{
                    "Err": 0,
                    "Name": "singed",
                    "IsInternal": false,
                    "Partitions": [
                    {
                            "Err": 0,
                            "ID": 0,
                            "Leader": 0,
                            "Replicas": [
                                     0
                            ],
                            "Isr": [
                            0
                            ],
                            "OfflineReplicas": null
                    },
                    {
                            "Err": 0,
                            "ID": 1,
                            "Leader": 0,
                            "Replicas": [
                            0
                            ],
                            "Isr": [
                            0
                            ],
                            "OfflineReplicas": null
                    }
                    ]
                    }
    
    # Delete a topic.
            ./kafka-cli topic -d=singed
            result:
                    {"level":"info","ts":1612424432.454704,"caller":"log/log.go:16","msg":"Delete Topic success","topic":"singed"}
    
    # Add partition number of topic
       ./kafka-cli topic --add-partition=singed --partition-num=3
            result:
                    {"level":"info","ts":1612424575.056782,"caller":"log/log.go:16","msg":"Add partition success","topic":"singed","partition num":3}
    
    
    Flags:
          --add-partition string      The Topic which need to create partition, partition num must higher than which already exists
      -b, --bootstrap-server string   The Kafka server to connect to.more than one should be separated by commas (default "localhost:9092")
      -c, --create string             Create a new topic.
      -d, --delete string             Delete a topic.
          --describe string           List details for the given topics.more than one should be separated by commas
      -h, --help                      help for topic
      -l, --list                      List all available topics.
          --partition-num int32       The specified partition when create topic or add partition (default 1)
          --replica-num int16         The specified replica when create topic (default 1)

**Producer**:

    A kafka synchronous producer, with pretty much config options. but it's not asynchronous, which means it will wait for result before return
    
    Usage:
      kafka-cli producer [flags]
    
    Examples:
    
    # Produce a message
        ./kafka-cli producer --bootstrap-servers=localhost:9092 --key=13 --partitioner=random --topic=singed --value='test value'
        result:
            {"level":"info","ts":1612429377.79058,"caller":"log/log.go:16","msg":"Send message success","partition":2,"offset":0}
    
    
    Flags:
      -b, --bootstrap-servers string   The Kafka server to connect to.more than one should be separated by commas (default "localhost:9092")
          --headers string             The headers of the message. Example: -headers=foo:bar,bar:foo
      -h, --help                       help for producer
          --key string                 the key of message
          --partition int32            The partition which message produce to, if provided, it will use manual partitioner (default -1)
          --partitioner string         The partitioning scheme to use. Can be hash, manual, or random (default "hash")
          --topic string               REQUIRED: The topic id to produce messages to.
          --value string               REQUIRED: The message content which is going to be produced
    
Please use `./kafka-cli -h` or `./kafka-cli [command] -h` for more detail.

## Compatibility
- **Tested on**
    - apache kafka 2.13.0
    - golang 1.15.7