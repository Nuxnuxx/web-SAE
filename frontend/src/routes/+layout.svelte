<script>
	import Header from "../components/header.svelte";
</script>

<Header />
<main>
	<slot />
</main>

<nav>
	<h1>This is the footer</h1>
	<a href="/">Link1</a>
	<a href="/">Link2</a>
	<a href="/">Link3</a>
</nav>
