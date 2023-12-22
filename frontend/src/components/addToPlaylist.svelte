<script lang="ts">
	import { enhance } from '$app/forms';
    import { onMount } from 'svelte';
    export let idRecipe: number;
    
    let isOpen = false;
  
    onMount(() => {
      document.addEventListener('click', handleOutsideClick);
      return () => {
        document.removeEventListener('click', handleOutsideClick);
      };
    });
  
    function toggleDropdown() {
      isOpen = !isOpen;
      updateDropdownVisibility();
    }
  
    function handleOutsideClick(event: Event) {
      const target = event.target as HTMLElement;
      if (!target.closest('.dropdown')) {
        isOpen = false;
      }
    }

    function updateDropdownVisibility() {
        const dropdownContent = document.querySelector('.dropdown-content') as HTMLElement;
        if (dropdownContent) {
        	dropdownContent.style.display = isOpen ? 'block' : 'none';
        }
    }
  </script>
  
  <style lang="scss">
    .dropdown {
    }
  
    .dropdown-content {
      display: none;
      position: absolute;
      background-color: #f9f9f9;
      box-shadow: 0px rgba(0, 0, 0, 0.2);
      z-index: 1;
      
    }
  
    .dropdown-item {
      cursor: pointer;
      /* reset all button properties */
      background: none;
	  border: none;
	  padding: 1rem;
	  cursor: pointer;
    }
  
    .dropdown-item:hover {
      background-color: #ddd;
    }

    button.material-symbols-rounded {
		/* reset all button properties */
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}
  </style>
  
  <div class="dropdown">
    <!-- Button to change isOpen to show or not the dropdown -->
    <button on:click={toggleDropdown} class="material-symbols-rounded">playlist_add</button>

    <div class="dropdown-content">
      <form use:enhance method="post" action="?/addPlaylistRecipe">
        <input hidden name="idRecipe" value={idRecipe} type="text" />
          <div>
            <button type="submit" name="idPlaylist" value={1} class="dropdown-item">
				Playlist 1
			</button>
          </div>
          <div>
            <button type="submit" name="idPlaylist" value={2} class="dropdown-item">
				Playlist 2
			</button>
          </div>
          <div>
            <button type="submit" name="idPlaylist" value={3} class="dropdown-item">
				Playlist 3
			</button>
          </div>
      </form>
    </div>
  </div>
  