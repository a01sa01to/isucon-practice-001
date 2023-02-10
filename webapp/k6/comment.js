import http from "k6/http";
import { BaseUrl, Username, Password } from "./config.js";
import { check } from "k6";
import { parseHTML } from "k6/html";

export default function () {
  const loginRes = http.post(`${BaseUrl}/login`, {
    account_name: Username,
    password: Password,
  });
  console.log(`loginRes: ${JSON.stringify(loginRes)}`);
  check(loginRes, {
    "is status 200": (r) => r.status === 200,
  });
}