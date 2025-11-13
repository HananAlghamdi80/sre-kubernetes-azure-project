const express = require("express");
const client = require("prom-client");
const app = express();

client.collectDefaultMetrics();

app.get("/healthz", (req, res) => res.send("ok"));     // Liveness
app.get("/ready", (req, res) => res.send("ready"));    // Readiness

const requests = new client.Counter({
    name: "api_requests_total",
    help: "Total number of API requests"
});

app.get("/api", (req, res) => {
    requests.inc();
    res.json({ message: "API service working ✅" });
});

app.get("/metrics", async (req, res) => {
    res.set("Content-Type", client.register.contentType);
    res.send(await client.register.metrics());
});

app.listen(3000, () => console.log("API running on port 3000"));
