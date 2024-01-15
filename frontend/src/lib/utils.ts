export function makeUrl(name: string, price: string, difficulty: string) {
	let page = 0;

	// FIXME: if bug page 0
	if (!name && !price && !difficulty) {
		page = 1;
	}
	const url = `/search${name ? `?name=${name}&` : "?"}page=${page}${
		price ? `&price=${price}` : ""
	}${difficulty ? `&difficulty=${difficulty}` : ""}`;
	return url;
}
