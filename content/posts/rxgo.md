---
title: "Rxgo"
date: 2020-12-09T17:26:53+08:00
draft: false
---

ReactiveX
ReactiveX, or Rx for short, is an API for programming with Observable streams. This is the official ReactiveX API for the Go language.

ReactiveX is a new, alternative way of asynchronous programming to callbacks, promises and deferred. It is about processing streams of events or items, with events being any occurrences or changes within the system. A stream of events is called an Observable.

An operator is a function that defines an Observable, how and when it should emit data. The list of operators covered is available here.

The RxGo implementation is based on the concept of pipelines. A pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function.



Let's see at a concrete example with each box being an operator:

We create a static Observable based on a fixed list of items using Just operator.
We define a transformation function (convert a circle into a square) using Map operator.
We filter each yellow square using Filter operator.
In this example, the final items are sent in a channel, available to a consumer. There are many ways to consume or to produce data using RxGo. Publishing the results in a channel is only one of them.

Each operator is a transformation stage. By default, everything is sequential. Yet, we can leverage modern CPU architectures by defining multiple instances of the same operator. Each operator instance being a goroutine connected to a common channel.

The philosophy of RxGo is to implement the ReactiveX concepts and leverage the main Go primitives (channels, goroutines, etc.) so that the integration between the two worlds is as smooth as possible.