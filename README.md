### Main pupose of an application
It's not about to build a backend application that make long url into short.

It's about domain-driven design implementation in golang. You can use it.


## What is domain-driven design.
Well it's a part of clean and onion architecture. It's main core of it.
There is 4 layers.

2 of them are internal. 2 others are replaceable. What does is mean?

Do you remember open-closed principle. Yep, it is from solid pattern!
Domain and application layers are business layers. So it doesn't matter do you want to make it look like website, desktop apllication, cli or whatever. 
Internal layers are immutable. You can replace redis to postgres, website to desktop application, but business logic would be the same. 
It means that it is opened to add different variant to store data or present it, but you cann't change way of using it.

### 1. Domain
It's internal layer of DDD and is used for defining entities and structure of repository(used to store data).
It is top layer. It depends on nothing. Only business logic can be there.

### 2. Application
It's also internal layer. But it is used to coordinate other 2 layers(adapters) using domain layer.
There is interface. What things should Interface layer present.
Service. What service should application provide.
Common. It is about result in common varian.
Command. It is about result but for specific commands using common.
Queries. It is about presenting in a real way data. I don't use it in this small project. But in next TaskFlow project I will.

### 3. Infrastracture
It is about working with data. It is also called "second adapter layers".
Do you remeber repo from domain layer. That was interface. Here we need to implement it. In my case there're just create and get, but it can be whatever you need. 

### 4. Interface
Presenting data should be made by using interface of application layer. Don't you forget it.
It validate data and call interface method(not implemented yet), it will be implemented in other layer.


### main layer
Personally I thing it is a separated layer because you need to mix infrastracture and interface. But Robert Martin don't think so!


![This is how it look like](https://i0.wp.com/herbertograca.com/wp-content/uploads/2018/11/100-explicit-architecture-svg.png?fit=1200%2C821&ssl=1)
