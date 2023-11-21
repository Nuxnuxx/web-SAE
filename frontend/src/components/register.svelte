<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />


<script>
    let email = '';
    let firstName = '';
    let lastName = '';
    let password = '';
    let confirmPassword = '';
    let gender = '';

    function handleSubmit() {
        // Handle form submission here
        console.log('Form submitted');
        console.log('Email:', email);
        console.log('First Name:', firstName);
        console.log('Last Name:', lastName);
        console.log('Password:', password);
        console.log('Confirm Password:', confirmPassword);
        console.log('Gender:', gender);
    }
</script>

<main>
    <h2>Rejoignez nous, et trouver l'inspiration<br>culinaire qui vous ressemble</h2>

    <form on:submit|preventDefault={handleSubmit}>
        <label for="email">
            <input 
                type="email"
                id="email" 
                bind:value={email} 
                required 
                placeholder="Email*" 
                autocomplete="username"
                />
        </label>
       
        <div class="name">
            <label class="lastname">
                <input 
                    type="text" 
                    id="lastName" 
                    bind:value={lastName} 
                    required 
                    placeholder="Nom*"
                    autocomplete="family-name"
                    />
            </label>

            <label class="firstname">
                <input 
                    type="text" 
                    id="firstName"
                    bind:value={firstName} 
                    required 
                    placeholder="Prénom*"
                    autocomplete="given-name"
                    />
            </label>
        </div>

        <label>
            <input 
                type="password" 
                id="password" 
                bind:value={password} 
                required 
                placeholder="Mot de passe*"
                autocomplete="new-password"
                />
        </label>

        <label>
            <input
                type="password" 
                id="confirmPassword" 
                bind:value={confirmPassword} 
                required 
                placeholder="Confirmer mot de passe*"
                autocomplete="new-password"
                />
        </label>

        <div class="option">
        <label class="checkbox"><input type="checkbox" bind:group={gender} value={"Homme"} /> Un pirate</label>
        <label class="checkbox"><input type="checkbox" bind:group={gender} value={"Femme"} /> Une pirate</label>
        <label class="checkbox"><input type="checkbox" bind:group={gender} value={"Autre"} /> Autre</label>
        </div>

        <button type="submit">
            Connexion
		    <span class="material-symbols-rounded"> arrow_forward </span>
	    </button>
    </form>
</main>



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

    label:not(.checkbox) {
        display: flex;
        flex-direction: column;
        margin-bottom: 10px;
        width: 50%;
        min-width: 350px;
        padding: 20px;
    }

    label.checkbox {
        margin-bottom: 5%;
    }

    @media screen and (min-width: 1000px) {
        div.name {
            display: flex;
            justify-content: space-between;
        }

        div.name label {
            width: 25%;
            min-width: 260px;
        }
    }

    div.option {
        display: flex;
        justify-content: space-between;
        width: 50%;
        min-width: 350px;
    }


    input:not([type=checkbox]) {
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