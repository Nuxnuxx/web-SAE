<script lang="ts">
    import Selector from "./selector.svelte";

    let answers: number[] = [];
</script>

<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0" />

<div class="filter">
    <div class="filter__title">
        <span class="material-symbols-outlined">tune</span>
        <span class="filter__title__text">Filter</span>
    </div>

    <div class="filter__card">
        <div class="filter__card__price">
            <span class="title__text">Prix</span>
            <Selector
				data={[
					"euro_symbol",
					"euro_symbol",
					"euro_symbol",
					"euro_symbol",
					"euro_symbol",
				]}
				rating_like={true}
				bind:selected={answers[1]}
			/>
        </div>

        <div class="filter__card__difficulty">
            <span class="title__text">Difficulté</span>
            <Selector
                data={[
                    "mood_bad",
					"sentiment_neutral",
					"sentiment_content",
					"sentiment_calm",
					"sentiment_satisfied",
                ]}
                rating_like={true}
                bind:selected={answers[2]}
            />
        </div>

        <div class="filter__card__tag">
            <span class="title__text">Tags</span>
            <Selector
                data={[
                ]}
                rating_like={true}
                bind:selected={answers[3]}
            />
        </div>
    </div>
</div>


<style lang="scss">
    .filter {
        
        .filter__title {
            display: flex;
            align-items: center;
            margin-bottom: 1rem;

            .material-symbols-outlined {
                font-size: 2rem;
                margin-right: 1rem;
            }

            .filter__title__text {
                font-size: 1.5rem;
                font-weight: 500;
            }
        }

        .filter__card {
            display: flex;
            flex-direction: column;
            border: 1px solid #e0e0e0;

            width: 70%;

            .filter__card__price,
            .filter__card__difficulty,

            .title__text {
                font-size: 1.2rem;
                font-weight: 500;
            }
        }
    }
</style>