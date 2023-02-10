import http from "k6/http";

const BaseUrl = "http://localhost:80";

export default function () {
  http.get(BaseUrl + "/");
}
