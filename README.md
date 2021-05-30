<p align="center">
  <img src="https://i.imgur.com/s9Gvwsg.png">
</p>

[![Tacklebox](https://img.shields.io/badge/tacklebox-case%20study-blue)](https://tacklebox-webhooks.github.io)

## Overview

These are the instructions to install and use the Go client library for Tacklebox.
Tacklebox is an open-source serverless framework that offers webhooks as a service.

It includes:
- a [CLI tool](https://github.com/tacklebox-webhooks/cli) to deploy and manage AWS infrastructure
- 4 client libraries ([JavaScript](https://github.com/tacklebox-webhooks/javascript),
    [Ruby](https://github.com/tacklebox-webhooks/ruby),
    [Python](https://github.com/tacklebox-webhooks/python),
    and [Go](https://github.com/tacklebox-webhooks/golang))
- a RESTful API with docs
- a management UI


You can read more about our case study [here](https://tacklebox-webhooks.github.io"),
and you can also watch our presentation [here](https://www.youtube.com/watch?v=QEFFlWNNwk8&t=1s).

## The Team

**Juan Palma** *Software Engineer* Phoenix, AZ

**Kevin Counihan** *Software Engineer* Seattle, WA

**Armando Mota** *Software Engineer* Los Angeles, CA

**Kayl Thomas** *Software Engineer* Atlanta, GA

## Getting Started

### Install the Go library

To install the Go library for Tacklebox, run:

```bash
go get github.com/tacklebox-webhooks/golang
```
---

### Use the Go library

Once you install the Go library for Tacklebox, you can start using it like so:

```go
import "github.com/tacklebox-webhooks/golang"

// Initialize a Tacklebox object using the API Key and API Host
// obtained after running 'tacklebox deploy'
tb := tacklebox.Webhooks{ApiKey: API_KEY, BaseUrl: API_HOST}
```

Once you include the package and initialize a Tacklebox object, you can do
many things. For example, you can create a service like so:

```go
serviceData := map[string]interface{}{
    "name": "service1"}

tb.Service().Create(serviceData)
```

Once you create services, event types, users and subscriptions,
you can create a new event like so:

```go
serviceId := "d90af763-5839-4a90-834c-5512980984f5"
userId := "cabea1b5-b485-41b7-8146-72ece22dc458"

eventData := map[string]interface{}{
	"event_type": eventType,
	"payload": map[string]interface{}{
		"message": "Hello from the Go wrapper!"},
	"idempotency_key": "8624862"}

tb.Event().Create(serviceId, userId, eventData)
```

If you want to see the message log for a specific user and service:

```Go
serviceId := "d90af763-5839-4a90-834c-5512980984f5"
userId := "cabea1b5-b485-41b7-8146-72ece22dc458"

tb.Message().List(serviceId, userId)
```
