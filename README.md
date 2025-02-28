# Go TCP Stack

This is a personal project to imporve my networking and golang knowlege

The goal of this project is to build my work TCP stack. But what does that mean?

I should be able to spin up a container (I am currently thinking using of using popman) this container will run a go
program which will be listenting for traffic on a given port number.

I will then send tcp traffic to the container via localhost and get responce

So what do I think would be an appropriate amount of set up?

Well I think it should be the absolue basic "loop" possible i.e. :

* Go program that opens a raw socket
* Program is built and put into a container
* container is deployed with the correct settings
* some kind of traffic is send to the container which can be logged out and inspected

That sounds like a reasonable set up to me

