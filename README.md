PUULP Metrics is a tool for creating per client/brand dashboards for monitoring metrics.

# riind :

Pocketbase service for user auth, org/client/brand management, file uploads, dashboard configs, etc.

It's not possible to scale horizontally (maybe this could be done with litestream?). it likely doesn't (perfectly) fit the usecase of loading in 3rd party data. Since sqlite locks the entire db during a write opperation, pit will be split off into a separate service for processing of data.

relevant links:

- pocketbase - https://github.com/pocketbase/pocketbase
- litestream - https://github.com/benbjohnson/litestream

Included in main.go is a go hook for opening a connection to a postgresql db and running a query & printing the results.

# piit :

Data engine/service. number crunching, ETL, data pipelines.

# peel :

The UI (Vuejs, if using PB consider switching to svelte). it is currently a shadcn-ui-vue premade "block" dashboard. shadcn-svelte currently has some css positioning impacting legends on charts.
