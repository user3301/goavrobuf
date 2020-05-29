# goavrobuf

Goavrobuf is a library that converts kafka avro schema to protocol buffer. This repo is a POC that attempts to deterministically convert fafka schema into a tree structure, and iterate the tree in breadth first search (BFS) manner to generate protocol buffer definition.
