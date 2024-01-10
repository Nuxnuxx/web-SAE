export function makeUrl(name: string, price: string, difficulty: string) {
	const url = `/search${name ? `?name=${name}&` : "?"}page=0${
		price ? `&price=${price}` : ""
	}${difficulty ? `&difficulty=${difficulty}` : ""}`;
	return url;
}
