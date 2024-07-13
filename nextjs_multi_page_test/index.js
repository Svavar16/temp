const express = require("express");

const app = express();

// app.use(express.static())

app.get("/api/hello", (req, res) => {
  res.status(200).json({ message: "Hello World!!" });
});

app.listen(3001, (req, res) => {
  console.log("Listening on http://localhost:3001");
});
