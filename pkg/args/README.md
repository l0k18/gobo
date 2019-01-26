# Args parser

## Types of arguments

### Trigger

Enables a step in the processing that is atomic and final.

    pod ctl init // reset to factory settings

### Variable

One or more key/value pairs. A key/value pair has no value if it is only one string. A variable defines what is effectively a parameter.

### Arrays

To save work for the user, any array of values is just a space separated list of items. If the item has a value but it is empty, the separator must be the terminating character

    pod node listeners 127.0.0.1:11011 \
        wss://gardengnomes.com rpc localhost \
        gardengnomes.com tls

This shows the list header 'listeners' and 'rpc' wwhich collect addresses until the next item does not parse as a valid address.

### Implicit boundaries

The keyword strings act as begin/end markers for groups of items. Commands are mandatory invocations, triggers are optional invocations of functions at startup.

Items inside a list can appear in any order, but no keyword can appear twice. This simple constraint means we don't have to burden the user with mandatory things that don't relate to actually doing something, since modes should be active until changed.

### Unambiguous Sequence

In all cases, if two items are members of the same group - meaning they have conflicting function in relation to each other, parsing aborts and the user is returned the error information, no forgiveness.

Yes, the file top level parsing in Go is the inspiration.

### Fast

Because this parser does not tolerate structural errors it parses very fast. It could be used to implement a binary command language also.