# Watchdog SDK

## Overview
Watchdog is a network monitoring service built on top of EdgeLQ and, just like EdgeLQ platform, uses same Goten
framework. Therefore, it follows same conventions, is gRPC and REST compatible.

Watchdog SDK provides:
* Set of proto files. They can be used to learn about API and can help in generating client library in
  practically any programming language.
* Dedicated Golang client library for this service. It was generated using proto files by dedicated protoc
  compilers from Goten.
* Docs in HTML format - located [here](./docs/apis), with full API reference.
* [Examples](./examples/cmd) - they use our Golang library.

In simple words, Watchdog SDK provides code for client-based applications talking to Watchdog service.

It depends on [Goten SDK](github.com/cloudwan/goten-sdk) and [EdgeLQ SDK](github.com/cloudwan/edgelq-sdk).

## Repository structure
Goten enforces common standards regarding code structure and its practically same in every SDK.

Description of basic directories:

### proto/v1alpha2
Provides description of APIs, resources and other objects in protobuf format. They can be used to learn about service
or be used in order to generate client libraries in other programming languages.

Separation of APIs and resources into subpackages allows building more light-weight applications.

### resources/v1alpha2/\<resource-name\>
Golang module for single resource. Contains definition of it and all helper objects - dedicated Filter, FieldMask,
set of FieldPaths etc.

### client/v1alpha2/\<api-name\>
Contains basic (almost like raw, but type safe) client for communication with RPC API and definitions for most
request/response objects.

### client/v1alpha2/watchdog
Contains sum of all clients for Watchdog. May not be recommended if we want to minimize size of application runtime.

### access/v1alpha2/\<resource-name\>
Contains high-level client-based components like Watcher - dedicated for each resource.

## How to use
Watchdog SDK utilizes [Goten SDK](github.com/cloudwan/goten-sdk) and [EdgeLQ SDK](github.com/cloudwan/edgelq-sdk),
which contain basic instructions. It does not require much extra - in order to develop Golang application with this SDK,
you need to install Go and include SDK in your list of dependencies. You can use [examples](./examples) as a help.

In order to generate SDK in other language, you need to follow steps regarding protoc installation, as described in Goten
and EdgeLQ SDKs. The only extra thing to do, is to include directory with this repository (on your machine) as additional
include proto path.

## Notes about examples
Since Goten and EdgeLQ enforce same standards everywhere (Goten enforces clients/API, EdgeLQ authentication and authorization),
we have put actually more detailed examples in EdgeLQ SDK repository. EdgeLQ SDK notes on examples contains also all
necessary information about authentication, so it may be good starting point if you have problems. This note will be
placed in every SDK repo that depends on EdgeLQ SDK.
