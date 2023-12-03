import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
  vus: 1000,
  duration: "30s",
};

export default function () {
  const res = http.get("http://localhost:3001/recipe/2");
  check(res, {
    "is status 200": (r) => r.status === 200,
  });

  const res2 = http.get("http://localhost:3001/recipe/page/3");
  check(res2, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1);
}
