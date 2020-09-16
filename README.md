## Charly Education assignment

The assignment consists in creating two very simple applications, one that emits messages on a queue and the other one that reads from the queue and if the message is of the expected type processes it.

#### The publisher

It published messages on a topic in the following JSON format:

```
{
    Type: "A", // that is a string
    Data: 5,   // that is an int
}
```

When you run it has to publish 10 messages on the topic. From 1 to 5 of type A and from 1 to 5 of type B.

Make sure that the messages are correctly delivered on the topic.

You will need a topic ID and a project ID as parameters. You can use environment variables or whatever you prefer.

#### The listener

It gets the messages published by the publisher and processes all the messages of a specific type (for example type "A"). That type can be set by some parameter or environment variable.

When a message has to be processed means you have to write its content on a file with the following function.

```
func printToFile(f *os.File, s string) error {
	x := rand.Int31n(10)
	if x < 6 {
		return errors.New("sorry couldn't make it")
	}

	fs := fmt.Sprintf("processed msg: %s\n", s)
	if _, err := f.WriteString(fs); err != nil {
		panic(fmt.Sprintf("unexpected error while writing on file, err: %v", err))
	}

	return nil
}
```

That function is likely to fail and the idea is that it represents a third party service. We assume that the error is always temporary, so if it returns an error it may succeed at a later stage. We want you to use ack/nack accordingly to reprocess the message at a later stage.

You will need a project ID, a topic ID, a subscription ID, a message type and a file path as parameters.

#### General notes

We set up the project as two different projects, a publisher and a listener but feel free to change that if you prefer, just explain your choices.

### What is important

* The logs should be something you want in production to understand what's going on. This is a simple application though, so there's no need to identify the same message from when it is published to when it is consumed.

* If you don't have enough time or it requires you too much time, pseudocode is enough and we can discuss about the solution later. It would be nice if you could read the bare minimum documentation for pubsub, in order to publish and listen for a message in order to handle or have and idea of how to handle these two operations.

* The idea is not to give you tricky questions but to give an idea of some of the technologies we use and the sense of some of the things we do.

* For any problem don't hesitate to ask. Even if something doesn't work and you feel dumb asking that is always better than assuming we did everything perfect. :) Just first try yourself and look on the internet. :)

### Technical information about pubsub

That is one of the Google Cloud services we use and it's messaging service and [here](https://cloud.google.com/pubsub/docs/overview) you can find more documentation about it.

One thing important is that you have to create one or more topic and one or more subscription. That can be done by an external main and you can assume that they exist in the publisher and subscriber(s). Or each application that needs a topic or a subscription makes sure it exists before doing anything else. That's up to you which way you prefer as long as you tell us what to do to run the application properly. :)

### Technical information about the emulator

In order to run pubsub emulator you just need to run `docker-compose up`.

That is our docker image for pubsub emulator. Official Google documentation can be found [here](https://cloud.google.com/pubsub/docs/emulator) but that shouldn't be needed.

In order to work properly with the emulator the following variable MUST be exported in any terminal will have to deal with pubsub emulator.

```
export PUBSUB_EMULATOR_HOST=127.0.0.1:8085
```

### To run solution

Just tell us where to download the solution and how to run it. It can be a `docker-compose` a script or `go run main.go`.