import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
	vus: 100000,
	duration: "30s",
};

export default function () {
	const res = http.get("http://localhost:5000/api/v1/recipe/id/23");
	check(res, {
		"is status 200": (r) => r.status === 200,
	});
	sleep(1);
}
