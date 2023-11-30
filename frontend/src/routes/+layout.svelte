<script>
	import Footer from "../components/footer.svelte";
	import Header from "../components/header.svelte";
</script>

<Header />
<main>
	<slot />
</main>

<Footer />

<style lang="scss">
	main {
		min-height: 100vh;
	}
</style>
