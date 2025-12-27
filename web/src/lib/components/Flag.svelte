<script lang="ts">
  let { code } = $props();
</script>

<!-- show nothing for no flag -->
{#if code}
  <img
    src="https://flagcdn.com/32x24/{code}.png"
    srcset="https://flagcdn.com/64x48/{code}.png 2x,
    https://flagcdn.com/96x72/{code}.png 3x"
    width="32"
    height="24"
    alt=""
    class="object-none" />
{/if}
