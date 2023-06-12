# assignment_demo_2023

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This project was built following the guide by Wei Xing so credit goes to him! \
URL to his guide: https://o386706e92.larksuite.com/docx/QE9qdhCmsoiieAx6gWEuRxvWsRc \
PS: This technolgy is extremely foreign to me so this was a good exposure! Hopefully one day I will be able to do something similar
w/o referencing a guide

This project only adds in code into the primarily the rpc-server and docker-compose yaml \
k8 deployment.yaml is also added in 

## main.go


Addition to this file includes the use of the initialisation of a redis server connected to a preset port number\
The RedisClient for go is also introduced to allow it to be called in the main method\

## redis.go

Most of the logic of for the redis client is provided in this file, such as initialising a RedisClient and its functions

#### Functions


InitClient(ctx context.Context, address, password string)
- Initialises a redisclient using the Redis package provided, and no password is set now

SaveMessage(ctx context.Context, roomID string, message *Message)
- Converts a message from its predefined struct to JSON for ease of serialisation into storage\
- This JSON message is then saved according to its roomID given


RoomIDMessages(ctx context.Context, roomID string, start, end int64, reverse bool)
- Obtains a roomID and returns the messages stored in the Room based off its recency\
- start and end defines the number of message to be sent back to the request\
- reverse indicates the order in which you wish to view the messages\


## message.go

This defines the fields for struct Message so that the JSON knows how to deserialise and serialise it

