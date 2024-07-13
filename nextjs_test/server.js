const express = require("express");
const path = require("path");
const fs = require("fs");

const app = express();

const staticPath = path.join(
  __dirname,
  "testing",
  "out",
  "_next",
  "static",
  "css"
);

const indexPath = path.join(__dirname, "testing", "out", "index.html");

try {
  if (fs.existsSync(indexPath)) {
    app.use(express.static(path.join(staticPath)));
    console.log("The file exists");
    console.log(typeof indexPath);
    app.get("/", (req, res) => {
      res.sendFile(indexPath);
    });
  }
} catch (error) {
  console.log(error);
}

app.get("/api/test", (req, res) => {
  res.status(200).send("Works");
});

app.listen(3001, () => {
  console.log("Test, http://localhost:3001");
});
