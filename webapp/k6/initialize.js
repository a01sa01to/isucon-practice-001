import http from "k6/http";
import { BaseUrl } from "./config.js";

export default function () {
  http.get(`${BaseUrl}/initialize`, {
    timeout: "10s",
  })
}
