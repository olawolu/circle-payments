# Payment Processor

This is a simple payment processing backend using Circle's APIs.

Billing a card is done in USD and funds are recieved in a circle wallet. More info about circle's api can be found [here](https://developers.circle.com/docs/circle-payments-api-quickstart)

This implementation uses mongodb as a backing store to keep track of the status of payments made per user.

It's not perfect, feedback on how the code could be better is appreciated.
