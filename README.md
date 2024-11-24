# go-dao
DAO interface

## Intention
Bring an abstraction between the service layer and the target storage. This interface allows to access different storage implementations without major changes in the app and service layer overall.
The interface is generic and providing a high flexibility for implementation.

Besides that, there are interfaces for configuration, factory, filter, and transaction.
