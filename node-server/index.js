const express = require("express");
const app = express();
const port = 3000;

app.get("/hello", (req, res) => {
  res.send("hello from node server");
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
