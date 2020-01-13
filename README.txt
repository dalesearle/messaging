Consumable: An immutable structure consumed by an interested party.
Container: A structure used to transport data through the infrastructure.
Package: A package moves in containers
Vertical: An attribute of a package used to determine domain, dental, auto, vet etc. Considered topics at the routing
    level.  Needs to be a setting or argument.  Do we want to send to multiple verticals?

Routing priority:
    1. Vertical (Topic)
    2. Package ID
Vertical(s)
A package can be sent to multiple verticals
A package ID is specific to its contents and encodings.  Packages with the same data but different encodings must have different ID's

Handlersf
Container handling
Package handling
Consumable handling

Options
Packages are tied to verticals as well as protocols, services listen to topics based on the package ID.
Vertical servers listen to vertical topics and manage messaging into and out of that domain
