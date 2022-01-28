console.log("Connecting Redis Timeseries...", process.env);
const RedisTimeSeries = require("redistimeseries-js");
const rts = require("redistimeseries-js");
const rts_client = new rts({
  url: process.env.REDIS_HOST,
});

function GetRange(key, sub, from, to, aggr, tb) {
  return new Promise((r, e) => {
    rts_client
      .revrange(`${key}:${sub}`, from, to)
      .aggregation(aggr, tb)
      .send()
      .then((res) => r(res[0][1]))
      .catch(() => r(null));
  });
}

async function MakeRow(key, from, to, aggr, tb) {
  let row = [new Promise((r) => r(to))];

  row.push(GetRange(key, "gas", from, to, aggr, tb));
  row.push(GetRange(key, "noise", from, to, aggr, tb));

  let meta = await new Promise((resolve, reject) => {
    rts_client.client.HGETALL(`${key}:meta`, (error, result) => {
      if (error) {
        console.error(error);
        return reject(error);
      }

      return resolve(result);
    });
  });

  row.push(meta.lat);
  row.push(meta.lng);
  row.push(key);
  row.push(meta.title);

  return Promise.all(row);
}

async function GatherData(from, to, aggr, tb) {
  console.log("Gathering Data", from, to, aggr, tb);
  // let keys = await rts_client.client.KEYS("*:meta");
  let keys = await new Promise((resolve, reject) => {
    rts_client.client.KEYS("*:meta", (error, result) => {
      if (error) {
        console.error(error);
        return reject(error);
      }

      return resolve(result);
    });
  });
  console.log("Got keys", keys);
  let promises = [];
  for (let key of keys) {
    key = key.split(":")[0];
    promises.push(MakeRow(key, from, to, aggr, tb));
  }

  return await Promise.all(promises);
}

module.exports = async () => {
  await rts_client.connect();
  console.log("Redis Timeseries connected");

  return { GatherData };
};
