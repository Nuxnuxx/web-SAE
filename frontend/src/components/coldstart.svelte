<script lang="ts">
  let currentQuestion = 0;
  let answers = [];

  const questions = [
    "Nous allons maintenant vous poser quelque<br><span class='rouge'>questions</span> pour apprendre à vous connaître.",
    "Quelle tranche de <span class='rouge'>prix</span> préférez vous ?",
    "Avec quel niveau de <span class='rouge'>difficulté</span> êtes vous familier ?",
    "Et pour les <span class='rouge'>épices</span> ?",
    "Pour finir, fait nous savoir ce que tu aimes en<br>choisissant <span class='rouge'>au moins 3</span> recettes."
  ];

  const options = [
    [], // pas de choix
    [
      { label: "€", selected: false, colorful: false, id: 1 },
      { label: "€", selected: false, colorful: false, id: 2 },
      { label: "€", selected: false, colorful: false, id: 3 },
      { label: "€", selected: false, colorful: false, id: 4 },
      { label: "€", selected: false, colorful: false, id: 5 },
    ], // Options pour la Question 2
    [
      { label: "<span class='material-symbols-rounded'>sentiment_satisfied</span>", selected: false, colorful: false, id: 1},
      { label: "<span class='material-symbols-rounded'>sentiment_calm</span", selected: false, colorful: false, id: 2},
      { label: "<span class='material-symbols-rounded'>sentiment_content</span", selected: false, colorful: false, id: 3},
      { label: "<span class='material-symbols-rounded'>sentiment_neutral</span", selected: false, colorful: false, id: 4},
      { label: "<span class='material-symbols-rounded'>mood_bad</span>", selected: false, colorful: false, id: 5},
    ],
    [
      { label: "<span class='material-symbols-rounded'>block</span>", selected: false, colorful: false},
      { label: "<span class='material-symbols-rounded'>local_fire_department</span", selected: false, colorful: false},
      { label: "<span class='material-symbols-rounded'>fire_truck</span", selected: false, colorful: false},
    ],
    //prendre des recipeCard de recipeCard.svelte
    [
        { label: "Poulet au curry", selected: false, colorful: false},
        { label: "Boeuf", selected: false, colorful: false},
        { label: "Poisson", selected: false, colorful: false},
    ]
  ]

  function handleNext() {
    currentQuestion++;
    console.log(currentQuestion)
  }

  function handlePass() {
    window.location.href = "/";
  }

  function handleSubmit() {
    // Send answers to the server
  }

  function handleOptionClick(index: number) {
    options[currentQuestion][index].selected = !options[currentQuestion][index].selected;
    // enlève les autres selected
    options[currentQuestion].forEach((option, i) => {
        if (i !== index) {
            option.selected = false;
        }
    });

    updateColorfulIcons(index);
  }

  function updateColorfulIcons(index: number) {
    const id = options[currentQuestion][index].id;

    if(id !== undefined){
        options[currentQuestion].forEach((option, i) => {
            if (i <= index && i > index - id) {
                option.colorful = true;
            } else {
                option.colorful = false;
            }
        });
    }
    else{
        options[currentQuestion].forEach((option, i) => {
            if (i === index) {
                option.colorful = true;
            } else {
                option.colorful = false;
            }
        });
    }
  }

</script>

<main>
    <div class="container">
        {#if currentQuestion < questions.length - 1 && currentQuestion !== 0}
            <h1>{@html questions[currentQuestion]}</h1>

            <div class="container__options">
                {#if options[currentQuestion].length > 0}
                    {#each options[currentQuestion] as { label, selected, colorful }, index}
                        <button on:click={() => handleOptionClick(index)} class={colorful ? 'colorful' : 'noColorful'}>{@html label}</button>
                    {/each}
                {/if}
            </div>

            <button class="container__buttonNext" on:click={handleNext}
                ><span class="buttonNext__text">Suivant</span>
                <span class="material-symbols-rounded"> arrow_forward </span>
            </button>
        {:else if currentQuestion === 0}
            <h1>{@html questions[currentQuestion]}</h1>
            <button class="container__buttonNext" on:click={handleNext}
                ><span class="buttonNext__text">Suivant</span>
                <span class="material-symbols-rounded"> arrow_forward </span>
            </button>
            <button class="container__buttonPass" on:click={handlePass}>Passez cette étape</button>
        {:else}
            <h1>{@html questions[currentQuestion]}</h1>

            <div class="container__options">
                {#if options[currentQuestion].length > 0}
                    {#each options[currentQuestion] as { label, selected, colorful }, index}
                        <button on:click={() => handleOptionClick(index)} class={colorful ? 'colorful' : 'noColorful'}>{@html label}</button>
                    {/each}
                {/if}
            </div>

            <button on:click={handleSubmit}>Submit</button>
        {/if}
    </div>
</main>


<style lang="scss">

    .container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        margin-bottom: 10%;
        margin-left: 20%;
        margin-right: 20%;

        .container__options {
            display: flex;
            justify-content: space-between;
            width: 100%;
            margin-top: 20px;
            margin-bottom: 20px;

            button {
                /*all unset puis charger les styles*/
                all: unset;
                font-size: 40px;
                cursor: pointer;

                /*si pas colorful*/
                &.noColorful {
                    opacity: 0.2;
                }

                /*si colorful*/
                &.colorful {
                    color: #82D3E6; /* Couleur pour les icônes en couleurs */
                    opacity: 1;
                }
            }
        }

        .container__buttonNext {
            padding: 10px 50px;
            background-color: #de403e;
            color: #fff;
            border: none;
            border-radius: 20px;
            cursor: pointer;
            font-size: 16px;
            display: flex;
            align-items: center;

            .buttonNext__text {
                margin-right: 10px;
            }
        }

        .container__buttonPass {
            margin-top: 10px;
            border: none;
            background-color: transparent;
            opacity: 0.4;
            text-decoration: underline;
        }
    }

    .rouge {
        color: red;
    }
</style>
