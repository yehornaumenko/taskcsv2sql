<h1>Run with --config flag</h1>
<p>E.g. <code>go run main.go --config=config.yaml </code>></p>

<h1>Building a docker image</h1>
<p>To build a docker image run <code>docker build -t taskcsv2sql:latest .</code> </p>

<h1>Running a docker container</h1>
<p>To run a docker container move to the root folder of this repository in a terminal and exec <code>docker run -v $PWD/vehicles.csv:/vehicles.csv</code> </p>