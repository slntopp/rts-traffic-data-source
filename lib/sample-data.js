const sample = [
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
    rows: [
      [1643094782000, 90, 100, 52.153641, 11.637535, "0x9d", "st29"],
      [1643094782000, 30, 10, 52.153688, 11.644451, "0x9f", "st30"],
      [1643094782000, 60, 80, 52.147398, 11.6397, "0x9e", "st32"],
      [1643094782000, 5, 5, 52.153783, 11.645348, "0x9a", "st31"],
    ],
    type: "table",
  },
];

module.exports = sample;
