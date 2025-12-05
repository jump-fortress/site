<script lang="ts">
  import stock from '$lib/assets/static/rocketlaunchers/stock.png';
  import original from '$lib/assets/static/rocketlaunchers/original.png';
  import mangler from '$lib/assets/static/rocketlaunchers/mangler.png';

  let { launcher }: { launcher: string } = $props();
</script>

{#if launcher === 'Stock'}
  <img src={stock} alt="" draggable="false" class="size-fit object-none select-none" />
{:else if launcher === 'Original'}
  <img src={original} alt="" draggable="false" class="size-fit object-none select-none" />
{:else if launcher === 'Mangler'}
  <img src={mangler} alt="" draggable="false" class="size-fit object-none select-none" />
{/if}

<style>
  img {
    filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg)
      brightness(100%) contrast(108%);
  }
</style>
