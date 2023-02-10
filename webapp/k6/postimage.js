import http from "k6/http";
import { BaseUrl, getRandomAccount } from "./config.js";
import { check } from "k6";
import { parseHTML } from "k6/html";

const testimage = open("./testimage.png", "b");

export default function () {
  const [Username, Password] = getRandomAccount();
  const loginRes = http.post(`${BaseUrl}/login`, {
    account_name: Username,
    password: Password,
  });
  check(loginRes, {
    "successfully logged in": (r) => r.status === 200,
  });

  const doc = parseHTML(loginRes.body);

  const csrf_token = doc.find('input[name=csrf_token]').first().attr('value')

  const post_res = http.post(BaseUrl, {
    csrf_token,
    file: http.file(testimage, "testimage.png", "image/png"),
    body: "Rakutan Mobile"
  });

  check(post_res, {
    "successfully posted": (r) => r.status === 200,
  });
}