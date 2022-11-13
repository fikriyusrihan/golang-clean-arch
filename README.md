# Clean Architecture with Golang 🏗️
Example of Clean Architecture implementation with Golang. Read more information about [Clean Architecture]. 

This project has 4 layers and each layer has a role as following:
| Directory | Layer |
| ------ | ------ |
| domain | Entities |
| infrastructure | Frameworks & Drivers |
| interface | Interface Adapters |
| usecase | Use Cases |

## Tools
* Echo (Web Framework)
* PostgreSQL (RDBMS)
* Testify (Testing)
* Viper (Configuration)
* Docker (Containerization)

## Special Thanks
I learn a lot about Clean Architecture with Golang from [Manakuro's Repository] and [Iman Tumorang's repository]. Kindly check their repository and article on Medium.

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [Clean Architecture]: <https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html> 
   [Manakuro's repository]: <https://github.com/manakuro/golang-clean-architecture>
   [Iman Tumorang's repository]: <https://github.com/bxcodec/go-clean-arch>