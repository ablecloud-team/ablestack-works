const PropertiesReader = require("properties-reader", {
  writer: { saveSections: true },
});

// const configPath = ".env.development";
// if (process.env.NODE_ENV === "production") {
//   configPath = ".env.production";
// }

const configProd = PropertiesReader(".env.production");
const configDev = PropertiesReader(".env.development");
const countProd = configProd.get("VUE_APP_VERSION_CODE") + 1;
const countDev = configDev.get("VUE_APP_VERSION_CODE") + 1;
// console.log(config, count);

var m = new Date();
var dateString =
  m.getFullYear() +
  ("0" + (m.getMonth() + 1)).slice(-2) +
  ("0" + m.getDate()).slice(-2) +
  ("0" + m.getHours()).slice(-2) +
  ("0" + m.getMinutes()).slice(-2) +
  ("0" + m.getSeconds()).slice(-2);

configProd.set("VUE_APP_RELEASE_DATE", dateString);
configDev.set("VUE_APP_RELEASE_DATE", dateString);
configProd.set("VUE_APP_VERSION_CODE", countProd);
configDev.set("VUE_APP_VERSION_CODE", countDev);

configProd.save(".env.production");
configDev.save(".env.development");
