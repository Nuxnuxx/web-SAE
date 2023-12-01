import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
  vus: 1000000,
  duration: "30s",
};

export default function () {
  const res = http.get("http://localhost:1323/2");
  check(res, {
    "status was 200": (r) => r.status == 200,
  });
  sleep(1);
}
