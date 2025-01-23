const t = require("./build/Release/speedy.node");

console.log(t.connect("sa", "", "localhost", "sigma", 1433));
