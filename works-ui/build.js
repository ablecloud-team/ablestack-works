const PropertiesReader = require("properties-reader", {
  writer: { saveSections: true },
});
const fs = require("fs");

const configProd = PropertiesReader(".env.production");
// const countProd = configProd.get("VUE_APP_VERSION_CODE") + 1;
// const countDev = configDev.get("VUE_APP_VERSION_CODE") + 1;
// console.log(config, count);

// var m = new Date();
// var dateString =
//   m.getFullYear() +
//   ("0" + (m.getMonth() + 1)).slice(-2) +
//   ("0" + m.getDate()).slice(-2); +
// ("0" + m.getHours()).slice(-2) +
// ("0" + m.getMinutes()).slice(-2) +
// ("0" + m.getSeconds()).slice(-2);

try {
  const data = fs.readFileSync("/mnt/jenkins-work/versionInfo.txt", "utf8");
  console.log(data);
  configProd.set("VUE_APP_VERSION", data);

  configProd.save(".env.production");
  // configDev.save(".env.development");
} catch (err) {
  console.log(err);
}
