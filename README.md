PUULP Metrics is a dashboard for monitoring data metrics and creating per client/brand dashboards.

back-end :
to validate use of pocketbase, test with prod-like data. it's not possible to scale horizontally (maybe this could be done with litestream?). it likely doesn't (perfectly) fit the usecase of loading in 3rd party ad data.

alternative : use PB as a separate service for all non-data centric functionality. e.g., user auth, file uploads, client management. create a separate service to act as the 'data engine', loading external data, serving metrics/ad data.

pocketbase - https://github.com/pocketbase/pocketbase https://pocketbase.io/
litestream - https://github.com/benbjohnson/litestream https://litestream.io/

## included in main.go is a go hook for opening a connection to a postgresql db and running a query & printing the results.

peel is the UI (Vuejs, if using PB consider switching to svelte)

it is currently a shadcn-ui-vue premade "block" dashboard.
