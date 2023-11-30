<script lang="ts">
  
    import { register } from 'swiper/element/bundle';

    import type { Recipe } from "../app";
    import RecipeCard from '../components/recipeCard.svelte';
	import { onDestroy } from 'svelte';
    export let listRecipe: Recipe[];

    register();
</script>

<!--
    TODO : custom navigation arrow next
-->

<div class="carousel">
    <swiper-container slides-per-view="2" speed="500" loop="true" css-mode="true"
    navigation="true"
    >
    <!-- Breakpoints me met une erreur dans le fichier swiper-core.mjs
        breakpoints={{
        600: {
          slidesPerView: 3,
        },
        768: {
          slidesPerView: 4,
        },
      }}-->
        {#each listRecipe as recipe}
            <swiper-slide>
                <div class="carousel__slide">
                    <RecipeCard data={recipe} />
                </div>
            </swiper-slide>
        {/each}
    </swiper-container>
</div>

<style lang="scss">
    .carousel {
        width: 90%;
        margin: 0 auto;

        .carousel__slide {
            width: 90%;
            height: 90%;
        }
    }
</style>