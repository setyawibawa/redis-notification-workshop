# Redis Messaging App

This app is used for a quick, send-then-destroy, secret chat, so the messages are only stored on Redis temporarily.

There are two components to be used:
1. Sender (`sender/main.go`)
   This component is the main client for a user to be online and send messages to the system
2. Display (`display/main.go`)
   This component is used to display which users are online in the system, and the messages sent

You do not need to edit anything on the Sender. However, the Display implementation is incomplete.

The Display should:
1. Print which users are online, *once*, using `printSomeoneOnline()` method. The Display can show users that went online before the Display component is run
2. Print which users went offline, using `printSomeoneOffline()` method.
3. Print messages sent by users, using `printMessage()` method.

Currently, the Display code is only prefilled with boilerplate Redis subscription code and predefined print methods. Complete the Display implementation, and configure your Redis instance so that it can utilize Redis Keyspace Notification properly