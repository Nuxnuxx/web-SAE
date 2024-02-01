export function makeUrl(name: string, price: string, difficulty: string) {
	let page = 0;

	const url = `/search${name ? `?name=${name}&` : "?"}page=${page}${
		price ? `&price=${price}` : ""
	}${difficulty ? `&difficulty=${difficulty}` : ""}`;
	return url;
}
