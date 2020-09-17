## Charly Education assignment

The assignment consists in creating two very simple applications. One application (publisher) emits messages on a queue and the other (listener) reads from it. The listener processes the message only when it is of the expected type.

#### The publisher

It publishes messages on a topic in the following JSON format:

```
{
    Type: "A", // string
    Data: 5,   // int
}
```

When you run the program, it publishes 10 messages on this topic. From data 1 to 5 of type A and from data 1 to 5 of type B.

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

It is likely that this function fails, and the concept is to represent an interaction with a third-party service. We assume that the error is always temporary and so, if it returns an error, it will succeed at a later stage. We want you to use ack/nack accordingly, in order to reprocess the message later.

You will need a project ID, a topic ID, a subscription ID, a message type and a file path as parameters.

#### General notes

We set up the project as two different projects, the publisher and the listener, but feel free to change that if you prefer, just explain your choices.

### What is important

* The logs should be something you want in production to understand what's going on. This is a simple application though, so there's no need to identify the same message from when it is published to when it is consumed.

* If you don't have enough time or it requires you too much time, pseudocode is enough and we can discuss the solution together. It would be nice if you could read the documentation for pubsub and try to handle publishing/listening a pubsub message or at least, to get an idea about it.

* The purpose of this assignment is to give you an idea of some of the technologies that we use and the sense of some of the things we do.

* In case you face any problems, don't hesitate to ask. Even if something doesn't work and you feel dumb asking that, it is always better asking, rather than assuming that we did everything perfect. :) Just first try yourself and look on the internet. :)

### Technical information about pubsub

[Pubsub](https://cloud.google.com/pubsub/docs/overview) is the messaging service that we use, provided by Google Cloud services.

The important thing here, is that you have to create one or more topics and one or more subscriptions. This can be done by an external main and you can assume that the topics/subscriptions exist in the publisher and subscriber(s). Otherwise, each application that needs a topic or a subscription, it makes sure that the topic/subscription exists before doing anything else. It's up to you which way you prefer as long as you tell us how to run the application properly. :)

### Technical information about pubsub emulator

In order to run pubsub emulator you just need to run `docker-compose up`.

That is our docker image for pubsub emulator. Official Google documentation can be found [here](https://cloud.google.com/pubsub/docs/emulator) but that shouldn't be needed.

In order to work properly with the emulator the following variable MUST be exported in any terminal that deals with pubsub emulator.

```
export PUBSUB_EMULATOR_HOST=127.0.0.1:8085
```

### To run solution

Just tell us where to download the solution and how to run it. It can be a `docker-compose` , a script or `go run main.go`.