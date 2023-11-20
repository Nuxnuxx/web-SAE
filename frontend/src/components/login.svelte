<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />


<script>
    let email = '';
    let password = '';

    function handleLogin() {
        // Ajoutez ici la logique de connexion
        console.log('Email:', email);
        console.log('Mot de passe:', password);
    }
</script>

<h2>Quel plaisir de vous voir à nouveau !</h2>

<form on:submit|preventDefault={handleLogin}>
    <label>
        <input type="email" bind:value={email} required placeholder="Email"/>
    </label>

    <label>
        <input type="password" bind:value={password} required placeholder="Mot de passe" />
    </label>

    <a href="#">Mot de passe oublié ?</a>

    <button type="submit">Connexion 
        <span class="material-symbols-rounded">
        arrow_forward
        </span>
    </button>
</form>

<style>
    h2 {
        text-align: center;
        color: #333;
        font-weight: bold;
        font-size: 22px;
    }

    form {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 20px;
    }

    label {
        display: flex;
        flex-direction: column;
        margin-bottom: 10px;
        /*min-width: 80%;*/
        /*width selon taille de l'écran*/
        width: 50%;
        min-width: 350px;
        padding: 20px;
    }

    input {
        border: 1px solid #ccc;
        border-radius: 3%;
        border-top: none;
        border-left: none;
        border-right: none;

        /*padding du texte*/
        padding: 5px 15px;
        font-size: 16px;
        outline: none;        
    }

    /*rend le placeholder un peu plus transparent par défaut*/
    input::-webkit-input-placeholder { opacity: 0.40;}

    /* ajoute de la transparence pour le placeholder lors du focus */
    input::-webkit-input-placeholder { -webkit-transition:.5s;}
    input:-moz-input-placeholder { -moz-transition:.5s;}
    input::-moz-input-placeholder { -moz-transition:.5s;}
    input:-ms-input-placeholder { -ms-transition:.5s;}

    input:focus::-webkit-input-placeholder { color:transparent;}
    input:focus:-moz-placeholder { color:transparent; } /* FF 4-18 */
    input:focus::-moz-placeholder { color:transparent; } /* FF 19+ */
    input:focus:-ms-input-placeholder { color:transparent; } /* IE 10+ */

    a{
        text-decoration: none;
        color: #9F9F9F;
        text-decoration: underline;
        font-size: 14px;
        margin-top: -20px;
        margin-bottom: 25px;
        margin-right: -40%;
    }

    button {
        padding: 10px 20px;
        background-color: #DE403E;
        color: #fff;
        border: none;
        border-radius: 20px;
        cursor: pointer;
        font-size: 16px;
        display: flex;
        align-items: center;
    }

    .material-symbols-rounded {
        font-size: 20px;
        font-weight: bold;
        margin-left: 10px;
        margin-right: -10px;
    }

</style>