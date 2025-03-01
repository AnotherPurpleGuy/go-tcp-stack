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

Current progress: Sat  1 Mar 2025 18:08:29 GMT

* Created the basic go program and docker file
* The docker build works
* Was able to deploy with podman
  - The program crashed to begin with as it did not have permission to open a socket on the host
  - Retried with out giving it access to the host network and it worked
  - Could not think of a quick way to direct traffic to the container (could try forwarding a port on the container)
  - Decided to make use of podman and put my image in a pod and run another container in that pod.
  - That got the result I wanted however I was also getting all the traffic from starting that pod and installing netcat
  - Was able to work around this by starting my conatiner second
  - I did not get the set up I was looking for though as I could see from the socket that a SYN was sent to the pods
    localhost and and ACK was being sent back however that was not in my program. The point of this is for me to build
    the whole stack including handling sending ACKs back so this is not the right set up yet
