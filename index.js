const express = require("express");
const app = express();

const bodyParser = require("body-parser");
const jsonParser = bodyParser.json();

(async function main() {
  const { GatherData } = await require("./lib/rts.js")();

  app.get("/", (req, res) => {
    console.log("Test invoked");
    res.status(200).send("");
  });

  app.post("/search", jsonParser, (req, res) => {
    console.log("Search invoked", req.body);
    res.status(200).send([
      {
        text: "all",
        value: 1,
      },
    ]);
  });

  app.post("/query", jsonParser, (req, res) => {
    console.log("Query received");
    const from = new Date(req.body.range.from).getTime();
    const to = new Date(req.body.range.to).getTime();

    const payload = req.body.targets[0].payload;
    let aggr = "none";
    if (typeof payload == "object" && payload.aggr) {
      aggr = payload.aggr;
    }

    GatherData(from, to, aggr);
    res.status(200).send(require("./lib/sample-data"));
  });

  const port = 8000;
  app.listen(port, () => {
    console.log(`App listening on port ${port}`);
  });
})();
