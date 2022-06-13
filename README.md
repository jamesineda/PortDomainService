# Port Domain Service

Run command: `./PortDomainService --FILEPATH=path/to/ports.json`

The service will process the specified JSON file and Create or Update a record in the database (in memory). Once finished the service will exit, unless an INT or TERM signal interrupts the process which will initiate a graceful shutdown. 