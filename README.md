This project is a simulated live market analysis RISK TOOL


The aim of this project is to gain proficiency: GoLang SQL (learnt in uni but didnt use in a actual project, learnt basic commands and queries to DB rather than implimenting) Backend interactions (how it interacts with the front end) Possible API approaches (possibly might have to impliment websockets to allow constant communications between back and front in realtime)


DISCLAIMER: NO AI WAS USED TO CREATE ANY PIECE OF CODE I AM DOING THIS FOR SELF LEARNING!




Key information:
we are taking an iterative programming approach.

First we will attempt to create a basic system in 1 process and no go routines (Price generator -> tick channel -> risk workers -> risk engine -> in-memory store)
Add Goroutines
API accepts portfolio, trade, limit, and query requests.
market-data ingestor connects to a feed and publishes normalized price events.
risk worker consumes price events, calculates risk, persists snapshots, and emits alerts.
The alert worker receives breach events and delivers notifications.
SQL Server stores durable business state.
SQS separates producers from consumers and absorbs temporary traffic spikes.
WebSockets push live risk changes to connected clients.
AWS runs the services and provides logs, metrics, alarms, networking, and secrets.


CYBERSECURITY ASPECT:
while its all fun to code it is important to add security and keep key features of programming:

•	Validate data when it enters the system.
•	Keep risk calculations independent of storage and transport technology.
•	Bound goroutines, channel buffers, request bodies, queues, and client buffers.
•	Let the goroutine that creates a channel decide when it is safe to close it.
•	Never hold a mutex while performing network or database work.   (DEADDDDDDDDLOCKKKKKKKKKKKKKK very scary)




For future self: This project inspiration came about when i went into the final stage interview and the questions mainly revolved around the deep understanding of networking. I answered the questions however was unable to answer when they dug deep as networking wasnt taught in university, so this project is to understand the full pipeline of different interactions. Context aside: final question at the end to answer: Do i fully understand how different applications interact with each other? (practical response not uni theory response)








