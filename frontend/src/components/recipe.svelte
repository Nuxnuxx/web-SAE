<script lang="ts">
  export let recipe = {
    title: 'Recipe Title',
    description: 'Recipe Description',
    temps: "40min",
    difficulty: "Facile",
    prix: "Abordable",
    imageRecette: 'https://assets.afcdn.com/recipe/20170607/67644_origin.jpg',
    ingredients: ['Ingredient 1', 'Ingredient 2'],
    imageIngredients: ['https://assets.afcdn.com/recipe/20170607/67644_origin.jpg', 'https://assets.afcdn.com/recipe/20170607/67365_origin.jpg'],    
    instructions: ['Recipe Instructions', 'Recipe Instructions 2']
  };
</script>

<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@24,400,0,0" />

<div>
  <div class="recipe__header">
    <h2>{recipe.title}</h2>
    <div class="recipe__header__infos">
      <div class="recipe__header__infos__item">
        <span class="material-symbols-rounded">alarm</span>
        <p>{recipe.temps}</p>
      </div>
      
      <div class="recipe__header__infos__item">
        <span class="material-symbols-rounded">sentiment_satisfied</span>
        <p>{recipe.difficulty}</p>
      </div>
      <div class="recipe__header__infos__item">
        <span class="material-symbols-rounded">euro_symbol</span>
        <p>{recipe.prix}</p>
      </div>
    </div>
  </div>
  
  <div class="recipe__line">
  </div>

  <div class="recipe__incredient">
    <ul class="recipe__ingredient__list">
      {#each recipe.ingredients as ingredient, i}
        <li>
          <img src={recipe.imageIngredients[i]} alt={ingredient} />
          <span class="recipe__ingredient__text">{ingredient}</span>
        </li>
      {/each}
    </ul>

    <div>
      <img class="recipe__img" src={recipe.imageRecette} alt="Image de la recette" aria-hidden="true" />
    </div>
  </div>

  <div class="recipe__line">
    <span class="recipe__line__text">
      Préparation
    </span>
  </div>

  <div class="recipe__instruction">
    {#each recipe.instructions as instruction, i}
      <p>Etape {i + 1}</p>
      <p class="recipe__instruction__text">{instruction}</p>
    {/each}
  </div>

  <div class="recipe__line">
  </div>
</div>



<style lang="scss">
  .recipe__header {
    h2 {
      color: #DE403E;
      text-align: center;
      font-family: Leckerli One;
      font-size: 96px;
      font-style: normal;
      font-weight: 400;
      line-height: normal;
      margin-bottom: 2%;
    }

    .recipe__header__infos {
      display: flex;
      justify-content: space-around;
      margin: 0 auto;
      width: 50%;
      margin-bottom: 50px;
      .recipe__header__infos__item {
        display: flex;
        flex-direction: row;
        align-items: center;
        
        .material-symbols-rounded {
          font-size: 35px;
          color:#DE403E;
          margin-right: 5px;
        }
      }
    }
  }


  .recipe__incredient {
    @media screen and (min-width: 600px) {
      display: flex;
      flex-direction: row;
      justify-content: space-around;
    }

    .recipe__ingredient__list {
      
      @media screen and (min-width: 500px) {
        columns: 2;
        -webkit-columns: 2;
        -moz-columns: 2;
      }      
      list-style-type: none;

      /*mettre le texte au centre de la colonne*/
      text-align: center;
      
      
      li {
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-bottom: 5px;
      }

      img {
        width: 100px;
        height: 100px;
        margin-right: 5px; 
        
        border: 1px solid black;
        border-radius: 10%;
      }
    }
  }


  .recipe__instruction {
    padding: 3% 10%;
    font-size: 20px;
    font-family: Leckerli One;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
    text-align: justify;

    p {
      color: #000;
      font-family: Leckerli One;
      font-size: 36px;
      font-style: normal;
      font-weight: 400;
      line-height: normal;
    }

    .recipe__instruction__text {
      color: #000;
      font-family: Inter;
      font-size: 24px;
      font-style: normal;
      font-weight: 400;
      line-height: normal;
    }
  }


  .recipe__img {
    width: 250px;
    height: 250px;
  }


  .recipe__line {
    width: 60%; 
    height: 0px;
    margin: 0 auto; 
    border-bottom: 1px solid black; 
    text-align: center;
    opacity: 0.5;
    padding-bottom: 2%;

    @media screen and (max-width: 800px) {
      margin-bottom: 5%;
    }

    .recipe__line__text {
      color: #000;
      font-family: Leckerli One;
      font-size: 40px;
      padding: 0 30px;
      background-color: #FFF;
    }
  }

</style>