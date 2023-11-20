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

    <button type="submit">Se connecter</button>
</form>

<style>
    h2 {
        text-align: center;
        color: #333;
        font-weight: normal;
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
        width: 40%;
        padding: 20px;
    }

    input {
        border: 1px solid #ccc;
        border-radius: 4px;
        border-top: none;
        border-left: none;
        border-right: none;

        /*padding du texte*/
        padding: 5px 20px;
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
        color: #333;
        font-size: 14px;
        margin-top: -20px;
        margin-bottom: 20px;
        margin-right: -30%;
    }

    button {
        padding: 10px 20px;
        background-color: #333;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }
</style>