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

  app.post("/query", jsonParser, async (req, res) => {
    console.log("Query received");
    const from = new Date(req.body.range.from).getTime();
    const to = new Date(req.body.range.to).getTime();

    const payload = req.body.targets[0].payload;
    let aggr = "avg";
    let tb = to - from;
    if (typeof payload == "object") {
      if (payload.aggr) aggr = payload.aggr;
      if (payload.tb) tb = payload.tb;
    }

    try {
      rows = await GatherData(from, to, aggr, tb);
      console.dir(rows);
      res.status(200).send([
        {
          columns: [
            { text: "Time", type: "time" },
            { text: "Gas", type: "number" },
            { text: "Noise", type: "number" },
            { text: "lat", type: "number" },
            { text: "lng", type: "number" },
            { text: "id", type: "string" },
            { text: "title", type: "string" },
          ],
          rows,
          type: "table",
        },
      ]);
    } catch (e) {
      console.error(e);
      res.status(400).send(e);
    }
  });

  const port = 8000;
  app.listen(port, () => {
    console.log(`App listening on port ${port}`);
  });
})();
