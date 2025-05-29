Where the backend code will live. Currently comparing python & golang web frameworks (unordered).

All options are suitble for a web application. Gather data on usage and maintenance before picking e.g., https://pypistats.org/packages/flask or https://pepy.tech/projects/fastapi?timeRange=threeMonths&category=version&includeCIDownloads=true&granularity=daily&viewType=line&versions=0.115.12%2C0.115.11%2C0.115.10

Top consideration = ^

Python:

- Flask
- FastAPI ^
- Starlette (probably with other deps for core func)

Go:

- chi ^
- stdlib ^
- gin
